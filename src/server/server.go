package server

import (
	"fmt"
	"github.com/TwinProduction/go-color"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golconda/src"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"
	"golconda/src/log"
)

const (
	defaultSnapLen = 262144
)

// PortHandlers get a list of ports, start listening on them and check if a connection arrives or not.
func PortHandlers(portsList string) {

	var ports []string
	// Get ports array
	ports = src.CheckPorts(portsList)
	// Must be debbuging
	log.Log.Debug("[+] Listening ports are: ", ports)
	var listeners []net.Listener
	// Start listening on all given ports
	for i := 0; i < len(ports); i++ {
		ln, err := net.Listen("tcp", ":"+ports[i])
		if err != nil {
		    log.Log.Error("[-] Error starting listener." + err.Error())
			os.Exit(1)
		}
		listeners = append(listeners, ln)
	}

	//Accept connection to listening ports and handle errors
	connChan := make(chan net.Conn)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for i := 0; i < len(listeners); i++ {
		go func(i int) {
			conn, err := listeners[i].Accept()
			if err != nil {
				log.Log.Error("[-] Error to accept the connection" + err.Error())
				os.Exit(1)
			}
			connChan <- conn
		}(i)
	}

	var limit int
	for {
		select {
		case conn := <-connChan:
			fmt.Println(color.Ize(color.Green, "[+] ")+"Reverse port open: ", strings.Split(conn.LocalAddr().String(), ":")[1])
			conn.Close()
			limit = limit + 1
			break
		case <-c:
			fmt.Println("\n" + color.Ize(color.Blue, "Program ended by the user. See you soon! :D "))
			os.Exit(1)
		}
		if limit == len(listeners) {
			fmt.Println(color.Ize(color.Purple, "({}) ") + "All listener(s) are open to reverse shell.")
			os.Exit(1)
		}
	}
}

// dumpLinuxSystem will listeing on an interface and display incoming ports from a target
func dumpLinuxSystem(interfaceName string, target string) {

	var interfaceBool bool = false
	if os.Getuid() != 0 {
		log.Log.Error("[-] Dump mode must be run as root")
		os.Exit(1)
	}
	// Get interfaces names
	devices, _ := pcap.FindAllDevs()
	for i := 0; i < len(devices); i++ {
		if devices[i].Name == interfaceName {
			fmt.Println(color.Ize(color.Green, "[+] ") + "Start monitoring interface " + interfaceName)
			interfaceBool = true
			break
		}
	}
	if interfaceBool == false {
		log.Log.Error("[-] Incorrect interface name")
		os.Exit(1)
	}

	handle, err := pcap.OpenLive(interfaceName, defaultSnapLen, true, 5*time.Minute)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	var eth layers.Ethernet
	var ip4 layers.IPv4
	//var ip6 layers.IPv6
	var tcp layers.TCP
	//var udp layers.UDP

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &tcp)
	decodedLayers := make([]gopacket.LayerType, 0, 10)

	for pkt := range packets {

		err = parser.DecodeLayers(pkt.Data(), &decodedLayers)
		if err != nil {
			log.Log.Error("[-] Error: Decoding a layer " + err.Error())
		}

		if net.ParseIP(target).Equal(ip4.SrcIP) {
			fmt.Println(color.Ize(color.Green, "[+] ")+"Reverse port open: ", tcp.DstPort)
		}
	}

}

// dumpWindowsSystem will listeing on an interface and display incoming ports from a target
func dumpWindowsSystem(interfaceName string, target string) {
	//fmt.Println("This feature is not working yet")
	var interfaceBool bool = false
	devices, _ := pcap.FindAllDevs()

	for i := 0; i < len(devices); i++ {
		if devices[i].Name == interfaceName {
			fmt.Println(color.Ize(color.Green, "[+] ") + "Start monitoring interface " + interfaceName)
			interfaceBool = true
			break
		}
	}
	if interfaceBool == false {
		log.Log.Error("[-] Incorrect interface name")
		os.Exit(1)
	}

	handle, err := pcap.OpenLive(interfaceName, defaultSnapLen, true, 5*time.Minute)
	if err != nil {
	    log.Log.Error("[-] Cannot open pcap live.")
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(handle, handle.LinkType()).Packets()
	var eth layers.Ethernet
	var ip4 layers.IPv4
	var ip6 layers.IPv6
	var tcp layers.TCP
	var udp layers.UDP
	var payload gopacket.Payload

	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &ip6, &tcp, &udp, &payload)
	decodedLayers := make([]gopacket.LayerType, 0, 10)
	var decodingError []string
	var listingPorts []layers.TCPPort
	for pkt := range packets {

		err = parser.DecodeLayers(pkt.Data(), &decodedLayers)
		if err != nil {
			if strings.Contains(err.Error(), "No decoder") {
				// Filter errors protocol
				var protoCheck = false
				proto := strings.Split(err.Error(), " ")[5]
				for _, d := range decodingError {
					if d == proto {
						protoCheck = true
					}
				}
				if protoCheck == false {
					decodingError = append(decodingError, strings.Split(err.Error(), " ")[5])
					log.Log.Error("[-] Error Decoding a layer " , decodingError)
				}
			} else {
				log.Log.Error("[-] Error Decoding a layer " + err.Error())
			}
		}
		if net.ParseIP(target).Equal(ip4.SrcIP) {
			// List ports open
			var openPorts = false
			for _, d := range listingPorts {
				if d == tcp.DstPort {
					openPorts = true
				}
			}
			if openPorts == false {
				listingPorts = append(listingPorts, tcp.DstPort)
				fmt.Println(color.Ize(color.Green, "[+] ")+"Reverse port open: ", tcp.DstPort)
			}

		}
	}
}

//MonitorInterface get interface name and the target to check incoming ports
func MonitorInterface(interfaceName string, target string) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\n"+color.Ize(color.Blue, "Program %v by the user. See you soon! :D"), sig)

			os.Exit(1)
		}
	}()
	if runtime.GOOS == "linux" {
		dumpLinuxSystem(interfaceName, target)
	} else if runtime.GOOS == "windows" {
		dumpWindowsSystem(interfaceName, target)
	} else {
    	log.Log.Error("[-] The system is not recognized, report it." + string(runtime.GOOS))
	}
}
