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

func MonitorInterface(interface_name string, target string){

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
      for sig := range c {
        fmt.Printf("\n" + color.Ize(color.Blue, "Program %v by the user. See you soon! :D"), sig)

        os.Exit(1)
      }
    }()
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
    var eth layers.Ethernet
    var ip4 layers.IPv4
    //var ip6 layers.IPv6
    var tcp layers.TCP
    //var udp layers.UDP

    parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet,&eth, &ip4, &tcp)
    decodedLayers := make([]gopacket.LayerType, 0, 10)

    for pkt := range packets {

        err = parser.DecodeLayers(pkt.Data(),&decodedLayers)
        if err != nil {
            fmt.Println(color.Ize(color.Red,"[-] " ) + "Error: Decoding a layer ", err)
        }

        if net.ParseIP(target).Equal(ip4.SrcIP){
            fmt.Println(color.Ize(color.Green,"[+] " ) + "Reverse port open: ", tcp.DstPort)
        }
    }
}
