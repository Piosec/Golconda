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

This option is based on the nmap-service file from: https://github.com/nmap/nmap

```bash
golconda client -t 10.10.10.9 --top-ports 100
```

```bash
golconda server --top-ports 100
```

## Dump 

Based on tcpdump, it's not working on Windows now. 

```bash
golconda server -d -i eth0 -t 10.10.10.10 
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

- New server feature, listening to the network interface to monitor upcoming ports
- Change the top-ports configuration to be more accurate
