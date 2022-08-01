package main

import (
	"altip/utils"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
)

const usage = `Usage of altip:
  -a, --address string    IP or Domain to obfuscate
  -p, --prefix string     Prefix to be added in front of the obfuscated ip
  -H, --host string       API host address to bind to (default "127.0.0.1")
  -P, --port integer      API port to listen on (default 8066)
  -s, --serve       	  Serve a public api endpoint
  -h, --help              Prints help information 
`

func getAlternativeIps(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	addr, _ := utils.Resolve(vars["ip"])
	prefix := ""

	if vars["prefix"] != "" {
		if reg, err := regexp.Compile("[^a-zA-Z\\d]+"); err == nil {
			prefix = reg.ReplaceAllString(vars["prefix"], "") + "://"
		}
	}

	if utils.IsValidIp(addr) == false {
		_, _ = fmt.Fprintf(w, "error: invalid ip")
		return
	}

	for _, ip := range utils.Obfuscate(prefix, addr) {
		_, _ = fmt.Fprintf(w, "%s", ip)
	}

}

func getHome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "AltIP usage examples:\n")
	_, _ = fmt.Fprintf(w, "curl https://altip.gogeoip.com/222.165.163.91\n")
	_, _ = fmt.Fprintf(w, "..with prefix:\n")
	_, _ = fmt.Fprintf(w, "curl https://altip.gogeoip.com/222.165.163.91/http\n")
	_, _ = fmt.Fprintf(w, "\n")
	_, _ = fmt.Fprintf(w, "Host it yourself, source code available at: https://github.com/Webklex/altip\n")
}

func main() {
	var (
		prefix, address, host string
		port                  uint
		serve                 bool
	)
	flag.StringVar(&address, "address", "", "IP or Domain to obfuscate")
	flag.StringVar(&address, "a", "", "IP or Domain to obfuscate")
	flag.StringVar(&prefix, "prefix", "", "Prefix to be added in front of the obfuscated ip")
	flag.StringVar(&prefix, "p", "", "Prefix to be added in front of the obfuscated ip")
	flag.StringVar(&host, "host", "127.0.0.1", "API host address to bind to")
	flag.StringVar(&host, "H", "127.0.0.1", "API host address to bind to")
	flag.UintVar(&port, "port", 8066, "API port to listen on")
	flag.UintVar(&port, "P", 8066, "API port to listen on")
	flag.BoolVar(&serve, "serve", false, "Serve a public api endpoint")
	flag.BoolVar(&serve, "s", false, "Serve a public api endpoint")
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	if serve {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", getHome)
		router.HandleFunc("/{ip}", getAlternativeIps)
		router.HandleFunc("/{ip}/{prefix}", getAlternativeIps)

		addr := fmt.Sprintf("%s:%d", host, port)
		fmt.Printf("Listening on: http://%s/\n", addr)

		log.Fatal(http.ListenAndServe(addr, router))
	} else if address != "" {
		if utils.IsValidIp(address) == false {
			fmt.Println("error: invalid ip")
			return
		}

		for _, ip := range utils.Obfuscate(prefix, address) {
			fmt.Printf("%s", ip)
		}
	}
}
