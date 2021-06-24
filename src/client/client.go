package client

import (
    "fmt"
    "net"
    "errors"
    "syscall"
    "os"
    "strings"
    "github.com/TwinProduction/go-color"
    "golconda/src"
)

func PortRunner(target string, ports_list string){
    var ports []string
    ports = src.CheckPorts(ports_list)
    // Must be debbuging
    fmt.Println(color.Ize(color.Green,"[+]") + " Checking available ports: ",ports)
    for i:=0;i<len(ports);i++ {
       _, err := net.Dial("tcp",target +":" + ports[i])
        if errors.Is(err,syscall.ECONNREFUSED) {
            fmt.Println(color.Ize(color.Yellow,"[-] ") + "Port " + ports[i] + " closed." )
        }else if err != nil {
            fmt.Println(color.Ize(color.Red,"[-] ") + "Error starting listener.", err.Error())
            os.Exit(1)
        }
    }
    fmt.Println(color.Ize(color.Purple,"({}) ") + "All ports on the remote targets are tested.")
}

func GetClientCommand(ip string, ports string, language string) string{
    listPort := strings.Join(src.CheckPorts(ports),",")
    switch(language){
        case "python":
            return "python -c \"import socket;[True for port in (" + listPort + ") if 0 == socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect_ex(('" + ip + "', port  )) ]\""
        case "bash":
            return  "bash -c 'for port in {" + listPort + "}; do echo >/dev/tcp/" + ip + "/$port;done'"
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
