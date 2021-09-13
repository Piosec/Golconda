package client

import (
	"errors"
	"golconda/src"
	"golconda/src/log"
	"net"
	"os"
	"strings"
	"syscall"
	"strconv"
	"time"
	)

//PortRunner get a target and a portList to check if the connection works.
func PortRunner(target string, portsList string, excludeports string, timeout int) {
	log.Log.Debug("The target to reach is: " + target)
	log.Log.Debug("The port(s) to try are (Before check): " + portsList)
	log.Log.Debug("The port(s) to exclude are (Before check): " + excludeports)

	//var ports []string
	// Get an array of ports
	//ports := src.CheckPorts(portsList)
	ports := src.PortsToExclude( src.CheckPorts(portsList),src.CheckPorts(excludeports))
    log.Log.Debug("The port(s) to try are (After check): " , ports)
	// Must be debbuging
	log.Log.Debug("[+] Checking available ports: ", ports)
	// foreach ports try to get a TCP connection
	for i := 0; i < len(ports); i++ {
		log.Log.Debug("Try connection on port: ", ports[i])
		_, err := net.DialTimeout("tcp", target+":"+ports[i],  time.Duration(timeout) * time.Second)
		if errors.Is(err, syscall.ECONNREFUSED) {
			log.Log.Info("[-] Port " + ports[i] + " closed.")
		} else if os.IsTimeout(err)  {
		    log.Log.Info("Timeout reached on port ", ports[i])
		    continue
		}else if err != nil {
			log.Log.Error("[-] Unqualified error during TCP connection: " + err.Error())
			os.Exit(1)
		}
	}
	log.Log.Info("({}) All ports on the remote target are tested.")
}

// GetClientCommand get an IP, port(s) and language to return oneliner
func GetClientCommand(ip string, ports string, language string, excludeports string, timeout int) string {
	log.Log.Debug("Generating client command from these values [ " + ip + ", " + ports + "," + language + " ]")
	//listPort := strings.Join(src.CheckPorts(ports), ",")
	//var listPort string
	portsArray := src.PortsToExclude(src.CheckPorts(ports),src.CheckPorts(excludeports))
	listPort := strings.Join(portsArray, ",")

    languageList := make(map[string]string)
    languageList["python"] =  "python -c \"import socket;socket.setdefaulttimeout(" + strconv.Itoa(timeout) +");[True if 0 == socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect_ex(('" + ip + "', port)) else next for port in (" + listPort + ")]\""
    languageList["bash"] =  "bash -c 'for port in {" + listPort + "}; do timeout " + strconv.Itoa(timeout) +" echo >/dev/tcp/" + ip + "/$port;done' 2>/dev/null"
    languageList["powershell"] = "(" + listPort + ") | % {echo ((new-object Net.Sockets.TcpClient).BeginConnect('" + ip + "',$_,$null,$null).AsyncWaitHandle.WaitOne(" + strconv.Itoa(timeout * 1000) + "))} 2>$null"
    languageList["perl"] = "perl -MIO::Socket -e 'for $port (" + listPort + "){$socket=IO::Socket::INET->new(Proto=>tcp,PeerAddr=>\"" + ip + "\",PeerPort=>$port, Timeout => " + strconv.Itoa(timeout) +") ;}'"
    languageList["ruby"] = "ruby -rsocket -rtimeout -e '\"" + listPort + "\".split(\",\").each do |port| sock = Socket.new(:INET, :STREAM);raw = Socket.sockaddr_in(port,\"" + ip + "\");begin Timeout.timeout(" + strconv.Itoa(timeout) +") do sock.connect(raw) end rescue nil end;sock.close end'"
    languageList[".net"] = "Here is the command to compile that oneliner, don't forget to replace the .net version: \r\n\r\necho using System;using System.Net;using System.Net.Sockets;namespace portscanner {class portscan {static void Main(){string s = \"" + listPort + "\";string[] ports = s.Split(',');foreach (var tmpPort in ports){var port = Convert.ToInt32(tmpPort);var timestamp = Convert.ToInt32(" + strconv.Itoa(timeout) +");TcpClient client = new TcpClient();client.ConnectAsync(\"" + ip + "\", port).Wait(timestamp * 1000);client.Close();}}}} > portscanner.net && c:\\Windows\\Microsoft.NET\\Framework\\v{.netVersion}\\csc.exe portscanner.net && portscanner.exe"

    return languageList[language]

/*
	switch language {
	case "python":
		return "python -c \"import socket;socket.setdefaulttimeout(" + strconv.Itoa(timeout) +");[True if 0 == socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect_ex(('" + ip + "', port)) else next for port in (" + listPort + ")]\""
	case "bash":
		return "bash -c 'for port in {" + listPort + "}; do timeout " + strconv.Itoa(timeout) +" echo >/dev/tcp/" + ip + "/$port;done' 2>/dev/null"
	case "powershell":
		return "(" + listPort + ") | % {echo ((new-object Net.Sockets.TcpClient).BeginConnect('" + ip + "',$_,$null,$null).AsyncWaitHandle.WaitOne(" + strconv.Itoa(timeout * 1000) + "))} 2>$null"
	case "perl":
		return "perl -MIO::Socket -e 'for $port (" + listPort + "){$socket=IO::Socket::INET->new(Proto=>tcp,PeerAddr=>\"" + ip + "\",PeerPort=>$port, Timeout => " + strconv.Itoa(timeout) +") ;}'"
	case "ruby":
		return "ruby -rsocket -rtimeout -e '\"" + listPort + "\".split(\",\").each do |port| sock = Socket.new(:INET, :STREAM);raw = Socket.sockaddr_in(port,\"" + ip + "\");begin Timeout.timeout(" + strconv.Itoa(timeout) +") do sock.connect(raw) end rescue nil end;sock.close end'"
	case ".net":
		var rev = "using System;using System.Net;using System.Net.Sockets;namespace portscanner {class portscan {static void Main(){string s = \"" + listPort + "\";string[] ports = s.Split(',');foreach (var tmpPort in ports){var port = Convert.ToInt32(tmpPort);var timestamp = Convert.ToInt32(" + strconv.Itoa(timeout) +");TcpClient client = new TcpClient();client.ConnectAsync(\"" + ip + "\", port).Wait(timestamp * 1000);client.Close();}}}}"
		return "Here is the command to compile that oneliner, don't forget to replace the .net version: \r\n\r\necho " + rev + " > portscanner.net && c:\\Windows\\Microsoft.NET\\Framework\\v{.netVersion}\\csc.exe portscanner.net && portscanner.exe"
	default:
		return "The language requested is not provided."
	}
*/
}
