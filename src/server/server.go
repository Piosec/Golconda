package server

import (
    "fmt"
    "net"
    "os"
    "os/signal"
    "strings"
    "github.com/TwinProduction/go-color"
    "golconda/src"
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