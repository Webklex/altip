# AltIP - Convert an IP into Alternative / Obfuscated versions
[![Demo][ico-website-status]][link-demo]
[![License][ico-license]][link-license]

[![Ko-Fi][ico-kofi]][link-kofi]

Inspired by [OsandaMalith/IPObfuscator](https://github.com/OsandaMalith/IPObfuscator). An 
online demo is available at: https://altip.gogeoip.com

## Usage
### CLI
```
Usage of altip:
  -a, --address string    IP or Domain to obfuscate
  -p, --prefix string     Prefix to be added in front of the obfuscated ip
  -H, --host string       API host address to bind to (default "127.0.0.1")
  -P, --port integer      API port to listen on (default 8066)
  -s, --serve       	  Serve a public api endpoint
  -h, --help              Prints help information 
```

#### Example without prefix
```bash
altip -a 127.0.0.1
```
```text
127.0.0.1
2130706433
0x7F.0x00.0x00.0x01
...
```

#### Example with prefix
```bash
altip -a 127.0.0.1 -p http://
```
```text
http://127.0.0.1
http://2130706433
http://0x7F.0x00.0x00.0x01
...
```

#### Serve the API
```bash
altip --serve
```
```text
Listening on: http://127.0.0.1:8066/
```


### API
API url: `/{ip or hostname}/{optional prefix}`

> The given hostname will be resolved (if possible) to its corresponding ip address, the results may vary.

#### Example without prefix
```bash
curl https://altip.gogeoip.com/127.0.0.1
```
```text
127.0.0.1
2130706433
0x7F.0x00.0x00.0x01
0177.0000.0000.0001
0x000000007F.0x0000000000.0x0000000000.0x0000000001
0000000177.0000000000.0000000000.0000000001
0x7F.0x00.0x00.1
0x7F.0x00.0.1
0x7F.0.0.1
0x7F.0x0.0x0.1
0x7F.0x0.0.1
0177.0000.0000.1
0177.0000.0.1
0177.0.0.1
0x7F.0x00.1
0x7F.0x0.1
0177.0000.1
0x7F.1
0177.1
0x7F.0x00.0000.0001
0x7F.0x0.0000.0001
0x7F.0000.0000.0001
0x7F000001
017700000001
0x7F.0000.1
127.0.1
127.1
```

#### Example with prefix
```bash
curl https://altip.gogeoip.com/127.0.0.1/http
```
```text
http://127.0.0.1
http://2130706433
http://0x7F.0x00.0x00.0x01
http://0177.0000.0000.0001
http://0x000000007F.0x0000000000.0x0000000000.0x0000000001
http://0000000177.0000000000.0000000000.0000000001
http://0x7F.0x00.0x00.1
http://0x7F.0x00.0.1
http://0x7F.0.0.1
http://0x7F.0x0.0x0.1
http://0x7F.0x0.0.1
http://0177.0000.0000.1
http://0177.0000.0.1
http://0177.0.0.1
http://0x7F.0x00.1
http://0x7F.0x0.1
http://0177.0000.1
http://0x7F.1
http://0177.1
http://0x7F.0x00.0000.0001
http://0x7F.0x0.0000.0001
http://0x7F.0000.0000.0001
http://0x7F000001
http://017700000001
http://0x7F.0000.1
http://127.0.1
http://127.1
```


## License
The MIT License (MIT). Please see [License File][link-license] for more information.

[ico-kofi]: https://ko-fi.com/img/githubbutton_sm.svg
[ico-license]: https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square
[ico-website-status]: https://img.shields.io/website?down_message=Offline&label=Demo&style=flat-square&up_message=Online&url=https%3A%2F%2Faltip.gogeoip.com%2F

[link-kofi]: https://ko-fi.com/webklex
[link-demo]: https://altip.gogeoip.com/
[link-license]: https://github.com/Webklex/altip/blob/master/LICENSE