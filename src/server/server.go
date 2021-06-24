package server

import (
    "fmt"
    "net"
    "os"
    "os/signal"
    "strings"
    "github.com/TwinProduction/go-color"
    "golconda/src"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
    "time"

)

const (
    defaultSnapLen = 262144
)


func PortHandlers(ports_list string){
    var ports []string
    ports = src.CheckPorts(ports_list)
    // Must be debbuging
    fmt.Println(color.Ize(color.Green,"[+]") + " Listening ports are: ",ports)
    var listeners []net.Listener
    for i:=0;i<len(ports);i++ {
        ln, err := net.Listen("tcp",":" + ports[i])
        if err != nil {
            fmt.Println(color.Ize(color.Red,"[-] ") + "Error starting listener.", err.Error())
            os.Exit(1)
        }
        listeners = append(listeners,ln)
    }

    connChan := make(chan net.Conn)
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    for i :=0; i<len(listeners);i++ {
        go func(i int){
            conn, err := listeners[i].Accept()
            if err != nil {
                fmt.Println("Error to accept the connection", err.Error())
                os.Exit(1)
            }
            connChan <- conn
        }(i)
    }

    var limit int
    for{
        select{
            case conn := <- connChan:
                fmt.Println(color.Ize(color.Green,"[+] ") + "Reverse port open: ",strings.Split(conn.LocalAddr().String(),":")[1])
                conn.Close()
                limit = limit + 1 
                break
            case <- c:
            fmt.Println("\n" + color.Ize(color.Blue,"Program ended by the user. See you soon! :D "))
            os.Exit(1)
        }
        if limit == len(listeners){
            fmt.Println(color.Ize(color.Purple,"({}) " ) + "All listener(s) are open to reverse shell.")
            os.Exit(1)
        }
    }
}

func MonitorInterface(interface_name string, target string, ports_list string){

    fmt.Println(interface_name, target,ports_list)
    var interface_bool bool = false
    if os.Getuid() != 0 {
        fmt.Println(color.Ize(color.Red,"[-] ") + "Dump mode must be run as root")
        os.Exit(1)
    }
    // Get interfaces names
    devices, _ := pcap.FindAllDevs()
    for i :=0;i<len(devices);i++ {
        if devices[i].Name == interface_name {
            fmt.Println(color.Ize(color.Green,"[+] ") + "Start monitoring interface " + interface_name)
            interface_bool = true
            break
        }
    }
    if interface_bool == false {
        fmt.Println(color.Ize(color.Red,"[-] ") + "Incorrect interface name")
        os.Exit(1)
    }

    handle, err := pcap.OpenLive(interface_name,defaultSnapLen, true, 5*time.Minute)
    if err != nil {
        panic(err)
    }
    defer handle.Close()

    packets := gopacket.NewPacketSource(handle,handle.LinkType()).Packets()
    for pkt := range packets {
    //    fmt.Println(pkt)
        var ip layers.IPv4
        ipLayer := pkt.Layer(layers.LayerTypeIPv4)
        if ipLayer != nil {
            ip, _ := ipLayer.(*layers.IPv4)
            fmt.Println(ip.SrcIP,ip.DstIP)
            if ip.SrcIP == net.ParseIP(target) {
                fmt.Println("IP target is " + target)
            }
        }

        tcpLayer := pkt.Layer(layers.LayerTypeTCP)
        if tcpLayer != nil {
            tcp, _ := tcpLayer.(*layers.TCP)
            fmt.Println(tcp.SrcPort,tcp.DstPort)
            if tcp.SrcPort == 443 {
                fmt.Println("port 443 ")
            }
        }
        fmt.Println(net.ParseIP(target))
        fmt.Println(ip.SrcIP)
       // if ip.SrcIP == net.ParseIP(target) {
       //     fmt.Println("IP target is " + target)
       // }
    }
}
