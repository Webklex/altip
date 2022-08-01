package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func IsValidIp(addr string) bool {
	if addr == "" {
		return false
	}
	segs := 0
	chcnt := 0
	accum := 0

	for _, ch := range addr {
		if ch == '.' {
			if chcnt == 0 {
				return false
			}

			if segs++; segs == 4 {
				return false
			}
			chcnt = 0
			accum = 0
			continue
		}
		cn, _ := strconv.Atoi(string(ch))
		if (cn < 0) || (cn > 9) {
			return false
		}
		if accum = accum*10 + cn - 0; accum > 255 {
			return false
		}
		chcnt++
	}

	if segs != 3 || chcnt == 0 {
		return false
	}
	return true
}

func Tokenize(addr string) []uint {
	segments := strings.Split(addr, ".")
	tokens := make([]uint, len(segments))
	for index, t := range segments {
		i, _ := strconv.Atoi(t)
		tokens[index] = uint(i)
	}
	return tokens
}

func SimpleTransform(format string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i == 3 {
			result += fmt.Sprintf(format, tokens[i])
		} else {
			result += fmt.Sprintf(format+".", tokens[i])
		}
	}
	return result
}

func ConditionalTransform(cond int, format, fallbackFormat string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i >= cond {
			if i == 3 {
				result += fmt.Sprintf(format, tokens[i])
			} else {
				result += fmt.Sprintf(format+".", tokens[i])
			}
		} else {
			result += fmt.Sprintf(fallbackFormat+".", tokens[i])
		}
	}
	return result
}

func TransformLeftShift(shift int, format, fallbackFormat string, tokens []uint) string {
	result := ""
	for i := 0; i < 4; i++ {
		if i < shift {
			result += fmt.Sprintf(format+".", tokens[i])
		} else {
			result += fmt.Sprintf(fallbackFormat, (tokens[2]<<8)|tokens[3])
			break
		}
	}
	return result
}

func Resolve(host string) (string, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}
	for _, ip := range ips {
		if addr := ip.String(); strings.Contains(addr, ":") == false {
			return addr, nil
		}
	}
	return "", nil
}

func stringListContains(check string, items []string) bool {
	for _, item := range items {
		if check == item {
			return true
		}
	}
	return false
}

func filterStringList(items []string) (result []string) {
	for _, item := range items {
		if stringListContains(item, result) == false {
			result = append(result, item)
		}
	}
	return
}

func Obfuscate(prefix, addr string) (ips []string) {
	tokens := Tokenize(addr)

	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, addr))
	ips = append(ips, fmt.Sprintf("%s%d\n", prefix, (tokens[0]<<24)|(tokens[1]<<16)|(tokens[2]<<8)|tokens[3]))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, SimpleTransform("0x%02X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, SimpleTransform("%04o", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, SimpleTransform("0x%010X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, SimpleTransform("%010o", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(3, "%d", "0x%02X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(2, "%d", "0x%02X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(1, "%d", "0x%02X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(3, "%d", "0x%0X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(2, "%d", "0x%0X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(1, "%d", "0x%0X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(3, "%d", "%04o", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(2, "%d", "%04o", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(1, "%d", "%04o", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, TransformLeftShift(2, "0x%02X", "%d", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, TransformLeftShift(2, "0x%0X", "%d", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, TransformLeftShift(2, "%04o", "%d", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, fmt.Sprintf("0x%02X.%d", tokens[0], (tokens[1]<<16)|(tokens[2]<<8)|tokens[3])))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, fmt.Sprintf("0x%0X.%d", tokens[0], (tokens[1]<<16)|(tokens[2]<<8)|tokens[3])))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, fmt.Sprintf("%04o.%d", tokens[0], (tokens[1]<<16)|(tokens[2]<<8)|tokens[3])))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(2, "%04o", "0x%02X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(2, "%04o", "0x%0X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(1, "%04o", "0x%02X", tokens)))
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, ConditionalTransform(1, "%04o", "0x%0X", tokens)))

	ips = append(ips, fmt.Sprintf("%s0x%X\n", prefix, (tokens[0]<<24)|(tokens[1]<<16)|(tokens[2]<<8)|tokens[3]))
	ips = append(ips, fmt.Sprintf("%s0%o\n", prefix, (tokens[0]<<24)|(tokens[1]<<16)|(tokens[2]<<8)|tokens[3]))

	result := ""
	for i := 0; i < 2; i++ {
		if i >= 1 {
			result += fmt.Sprintf("%04o.", tokens[i])
			result += fmt.Sprintf("%d", (tokens[2]<<8)|tokens[3])
		} else {
			result += fmt.Sprintf("0x%02X.", tokens[i])
		}
	}
	ips = append(ips, fmt.Sprintf("%s%s\n", prefix, result))

	parts := make([]string, 0)
	for _, token := range tokens {
		if token != 0 {
			parts = append(parts, fmt.Sprintf("%d", token))
		}
	}
	if tokens[1] == 0 {
		ips = append(ips, fmt.Sprintf("%s%d.%d.%d\n", prefix, tokens[0], tokens[2], tokens[3]))
		if tokens[2] == 0 {
			ips = append(ips, fmt.Sprintf("%s%d.%d\n", prefix, tokens[0], tokens[3]))
			if tokens[3] == 0 {
				ips = append(ips, fmt.Sprintf("%s%d\n", prefix, tokens[0]))
			}
		}
	}
	if tokens[2] == 0 {
		ips = append(ips, fmt.Sprintf("%s%d.%d.%d\n", prefix, tokens[0], tokens[1], tokens[3]))
		if tokens[3] == 0 {
			ips = append(ips, fmt.Sprintf("%s%d.%d\n", prefix, tokens[0], tokens[1]))
		}
	}
	if tokens[3] == 0 {
		ips = append(ips, fmt.Sprintf("%s%d.%d.%d\n", prefix, tokens[0], tokens[1], tokens[2]))
	}

	return filterStringList(ips)
}
