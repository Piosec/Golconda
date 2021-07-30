package client

import (
	"errors"
	"golconda/src"
	"golconda/src/log"
	"net"
	"os"
	"strings"
	"syscall"
)

//PortRunner get a target and a portList to check if the connection works.
func PortRunner(target string, portsList string, excludeports string) {
	log.Log.Debug("The target to reach is: " + target)
	log.Log.Debug("The port(s) to try are (Before check): " + portsList)
	log.Log.Debug("The port(s) to exclude are (Before check): " + excludeports)
	var ports []string
	// Get an array of ports
	ports = src.CheckPorts(portsList)
	ports = src.PortsToExclude(ports,excludeports)
    log.Log.Debug("The port(s) to try are (After check): " , ports)
	// Must be debbuging
	log.Log.Debug("[+] Checking available ports: ", ports)
	// foreach ports try to get a TCP connection
	for i := 0; i < len(ports); i++ {
		log.Log.Debug("Try connection on port: ", ports[i])
		_, err := net.Dial("tcp", target+":"+ports[i])
		if errors.Is(err, syscall.ECONNREFUSED) {
			log.Log.Info("[-] Port " + ports[i] + " closed.")
		} else if err != nil {
			log.Log.Error("[-] Unqualified error during TCP connection: " + err.Error())
			os.Exit(1)
		}
	}
	log.Log.Info("({}) All ports on the remote target are tested.")
}

// GetClientCommand get an IP, port(s) and language to return oneliner
func GetClientCommand(ip string, ports string, language string) string {
	log.Log.Debug("Generating client command from these values [ " + ip + ", " + ports + "," + language + " ]")
	listPort := strings.Join(src.CheckPorts(ports), ",")
	switch language {
	case "python":
		return "python -c \"import socket;[True for port in (" + listPort + ") if 0 == socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect_ex(('" + ip + "', port  )) ]\""
	case "bash":
		return "bash -c 'for port in {" + listPort + "}; do echo >/dev/tcp/" + ip + "/$port;done'"
	case "powershell":
		return "(" + listPort + ") | % {echo ((new-object Net.Sockets.TcpClient).Connect(" + ip + ",$_))} 2>$null"
	case "perl":
		return "perl -MIO::Socket -e 'for $port (" + ports + "){$socket=IO::Socket::INET->new(Proto=>tcp,PeerAddr=>\"" + ip + "\",PeerPort=>$port);}'"
	case "ruby":
		return "ruby -rsocket -e '\"" + ports + "\".split(\",\").each do |port| sock = Socket.new(:INET, :STREAM);raw = Socket.sockaddr_in(port,\"" + ip + "\");sock.connect(raw);sock.close end'"
	case ".net":
		var rev = "using System;using System.Net;using System.Net.Sockets;namespace portscanner {class portscan {static void Main(){string s = \"" + ports + "\";string[] ports = s.Split(',');foreach (var tmpPort in ports){var port = Convert.ToInt32(tmpPort);IPEndPoint ip = new IPEndPoint(IPAddress.Parse(\"" + ip + "\"), port);Socket server = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);server.Connect(ip);}}}}"
		return "Here is the command to compile that oneliner, don't forget to replace the .net version: \r\n\r\necho " + rev + " > portscanner.net && c:\\Windows\\Microsoft.NET\\Framework\\v{.netVersion}\\csc.exe portscanner.net && portscanner.exe"
	default:
		return "The language requested is not provided."
	}

}
