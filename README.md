# Alternative / Obfuscated IPs
Inspired by https://github.com/OsandaMalith/IPObfuscator

Demo available online at: https://altip.gogeoip.com

### Usage
API url: /{ip}/{optional prefix}

curl https://altip.gogeoip.com/222.165.163.91
```text
3735397211
0xDE.0xA5.0xA3.0x5B
0336.0245.0243.0133
0x00000000DE.0x00000000A5.0x00000000A3.0x000000005B
0000000336.0000000245.0000000243.0000000133
0xDE.0xA5.0xA3.91
0xDE.0xA5.163.91
0xDE.165.163.91
0336.0245.0243.91
0336.0245.163.91
0336.165.163.91
0xDE.0xA5.41819
0336.0245.41819
0xDE.10855259
0336.10855259
0xDE.0xA5.0243.0133
0xDE.0245.0243.0133
0xDE.0245.41819
```

curl https://altip.gogeoip.com/222.165.163.91/http
```text
http://3735397211
http://0xDE.0xA5.0xA3.0x5B
http://0336.0245.0243.0133
http://0x00000000DE.0x00000000A5.0x00000000A3.0x000000005B
http://0000000336.0000000245.0000000243.0000000133
http://0xDE.0xA5.0xA3.91
http://0xDE.0xA5.163.91
http://0xDE.165.163.91
http://0336.0245.0243.91
http://0336.0245.163.91
http://0336.165.163.91
http://0xDE.0xA5.41819
http://0336.0245.41819
http://0xDE.10855259
http://0336.10855259
http://0xDE.0xA5.0243.0133
http://0xDE.0245.0243.0133
http://0xDE.0245.41819
```

## License
The MIT License (MIT). Please see [License File][link-license] for more information.

[link-license]: https://github.com/Webklex/altip/blob/master/LICENSE