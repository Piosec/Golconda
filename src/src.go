package src

import (
    "fmt"
    "strconv"
    "os"
    "strings"
    "github.com/TwinProduction/go-color"
    "regexp"
    )
func checkPortsRange(portsRange []string){
    for i:=0;i<len(portsRange);i++ {
        converted, err := strconv.Atoi(portsRange[i])
        if err != nil {
            fmt.Println("Conversion error: ", err)
            os.Exit(1)
        }
        if converted < 1 ||  converted > 65535 {
            fmt.Println(color.Ize(color.Red,"[-] ") + "The value " + color.Ize(color.Red,strconv.Itoa(converted)) + " is not in the ports range (1-65535).")
            os.Exit(1)
        }
    }
}

func CheckPorts(ports string) []string {

    var ports_list []string
    if strings.Contains(ports,"-") {
        splited := strings.Split(ports,"-")
        converted0, err := strconv.Atoi(splited[0])
        if err != nil {
            fmt.Println("Conversion error: ", err)
            os.Exit(1)
        }
        converted1, err := strconv.Atoi(splited[1])
        if err != nil {
            fmt.Println("Conversion error: ", err)
            os.Exit(1)
        }
        for i := converted0;i<converted1+1;i++ {
            ports_list = append(ports_list, strconv.Itoa(i))
        }
    } else if strings.Contains(ports,",") {
        ports_list = strings.Split(ports,",")
    } else {
        ports_list = append(ports_list,ports)
    }
    // Check ports in range 1 --> 65535
    checkPortsRange(ports_list)
    return ports_list
}

func ValidateIP(ip string) bool {
    var check = false
    if len(ip) > 6 && len(ip) < 16{
        check,_ := regexp.MatchString(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}`,ip)
        return check
    }
    return check
}