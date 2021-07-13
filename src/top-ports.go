package src

import (
	"os"
	//    "bufio"
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed nmap-services
var f embed.FS

// Read file
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// GetTopPorts get an int and return a string of the top ports generated from nmap-services file.
func GetTopPorts(number int) string {

	data, _ := f.ReadFile("nmap-services")
	listingPorts := make(map[string]float64)
	var sData = string(data)
	splitedFile := strings.Split(sData, "\r\n")
	for i := 0; i < len(splitedFile); i++ {
		if strings.HasPrefix(splitedFile[i], "#") == false && splitedFile[i] != "" {
			splited := strings.Split(splitedFile[i], "	")
			//Output:  unknown 65514/tcp 0.000076
			//check if port is tcp or udp
			protocol := strings.Split(splited[1], "/")
			if protocol[1] == "tcp" {
				frequency, err := strconv.ParseFloat(splited[2], 64)
				if err != nil {
					fmt.Printf("%d of type %T", frequency, frequency)
					os.Exit(1)
				}
				listingPorts[protocol[0]] = frequency
			}
			//else if  protocol[1] == "udp" {
			//not now
			//}
			//fmt.Println(splited[0], splited[1], splited[2])
		}
	}

	// Add all ports to a map
	sortedPorts := make([]string, 0, len(listingPorts))
	for port := range listingPorts {
		sortedPorts = append(sortedPorts, port)
	}
	// Sort ports
	sort.Slice(sortedPorts, func(i, j int) bool {
		return listingPorts[sortedPorts[i]] > listingPorts[sortedPorts[j]]
	})

	// Get the number ports required
	count := 0
	ports := ""
	for _, name := range sortedPorts {
		if count == number {
			break
		} else if ports != "" {
			ports += ","
		}
		ports += name
		count += 1
	}
	return ports

}
