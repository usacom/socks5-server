# socks5-server
Simple socks5 server. You may user it as a proxy for Telegram

![size](https://img.shields.io/badge/image%20size-4.47MB-brightgreen.svg)

# Usage

First, you need to add user-pass into `users.example` and rename it to `users`

## bin
```
go build s5
PORT=1111 FILE="/path/to/file" ./s5
```