package src

import (
	"golconda/src/log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// checkPortsRange check whether the port is between 1 and 65535
func checkPortsRange(portsRange []string) {
	for i := 0; i < len(portsRange); i++ {
		converted, err := strconv.Atoi(portsRange[i])
		if err != nil {
			log.Log.Error("[-] Conversion error: " + err.Error())
			os.Exit(1)
		}
		if converted < 1 || converted > 65535 {
			log.Log.Error("[-] The value " + strconv.Itoa(converted) + " is not in the ports range (1-65535).")
			os.Exit(1)
		}
	}
}

// CheckPorts get a string of ports and convert it to an array of ports
func CheckPorts(ports string) []string {
	// The given string can be one port, X ports separated by comma or port a to b separated by dash
	var portsList []string
	//Check if the string contains -
	if strings.Contains(ports, "-") {
		// Split the string to get the values
		splited := strings.Split(ports, "-")
		converted0, err := strconv.Atoi(splited[0])
		if err != nil {
			log.Log.Error("[-] Conversion error: " + err.Error())
			os.Exit(1)
		}
		converted1, err := strconv.Atoi(splited[1])
		if err != nil {
			log.Log.Error("[-] Conversion error: " + err.Error())
			os.Exit(1)
		}
		for i := converted0; i < converted1+1; i++ {
			portsList = append(portsList, strconv.Itoa(i))
		}
		// Check if the string contains ,
	} else if strings.Contains(ports, ",") {
		portsList = strings.Split(ports, ",")
	} else {
		portsList = append(portsList, ports)
	}
	// Check ports in range 1 --> 65535
	checkPortsRange(portsList)
	return portsList
}

// ValidateIP get a string IP and check if that string match the following regex
func ValidateIP(ip string) bool {
	var check = false
	if len(ip) > 6 && len(ip) < 16 {
		check, _ := regexp.MatchString(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}`, ip)
		log.Log.Debug("IP address is " + ip)
		log.Log.Debug("check value is ", check)
		return check
	}
	return check
}
