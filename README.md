[![Go](https://github.com/Piosec/Golconda/actions/workflows/go-release.yml/badge.svg)](https://github.com/Piosec/Golconda/actions/workflows/go-release.yml)

# Golconda

Golconda is a client/server that aim to find reverse port open. 

```
$ ./golconda 
A Fast client / server reverse port discovery made with love.
         Please use it with love too <3.

Usage:
  golconda [command]

Available Commands:
  client      Use client to discover reverse ports.
  help        Help about any command
  server      Use server to listen reverse ports.

Flags:
  -h, --help   help for golconda

Use "golconda [command] --help" for more information about a command.
```

# Examples

## Basic

```bash
golconda server -p 8080,9000,9001
```

```bash
golconda client -t 127.0.0.1 -p 8080,9000,9001
```

## Port range

```bash
golconda server -p 8080-9001
```

```bash
golconda client -t 127.0.0.1 -p 8080-9001
```

## Print client oneliner

```bash
golconda client -t 127.0.0.1 -p 8080,9000,9001 -c powershell
```

## Top ports 

This option is based on the nmap-services file from: https://github.com/nmap/nmap

```bash
golconda client -t 10.10.10.9 --top-ports 100
```

```bash
golconda server --top-ports 100
```

## Monitor interface 

This feature monitor upcoming connections from a target. 

### Linux

Based on tcpdump, it's not working on Windows now. 

```bash
golconda server -d -i eth0 -t 10.10.10.10 
```

### Windows

First, get the mac address of the adaptator you want to monitor 

```bash
ipconfig /all
```

Then use **getmac** to list the Windows device name. 

```bash
Physical Address    Transport Name
=================== ==========================================================
00-FF-78-16-5B-54   Media disconnected
00-FF-1D-D9-30-08   Media disconnected
C0-B8-83-EB-39-E9   \Device\Tcpip_{BE995E05-272D-4004-BD00-48767E71E5FF}
00-50-56-C0-00-01   \Device\Tcpip_{FF430151-6520-4A67-AAF9-82B6A9A3F75A}
00-50-56-C0-00-08   \Device\Tcpip_{DB6EA806-DB19-4E40-BDDA-5ECEADAC1413}
C0-B8-83-EB-39-ED   Media disconnected
00-FF-64-8C-46-01   Media disconnected
00-FF-B2-BC-FF-42   Media disconnected
```

Change Tcpip to NPF.

```bash
golconda.exe server -d -i \Device\NPF_{DB6EA806-DB19-4E40-BDDA-5ECEADAC1413} -t 10.10.10.10 
```

# Errors 

## Too many open files

```bash
[-] Error starting listener. listen tcp :XXXXX: socket: too many open files
```

### Check your limit size

```bash
$ ulimit -n
```

### Change your limit size

```bash
$ ulimit -n 65600
```

# Todo 

- [x] New server feature, listening to the network interface to monitor upcoming ports
- [x] Change the top-ports configuration to be more accurate
- [ ] Exclusion ports
- [ ] Add logging
- [ ] Change GetClientCommand to key value strings