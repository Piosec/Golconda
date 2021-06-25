package src

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
    "sort"
)

// Read file
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func GetTopPorts(number int) string{

    listingPorts := make(map[string]float64)
    file, err := os.Open("src/nmap-services")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line,"#") == false {
            splited := strings.Split(line,"	")
            //Output:  unknown 65514/tcp 0.000076
            //check if port is tcp or udp
            protocol := strings.Split(splited[1],"/")
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

    sortedPorts := make([]string, 0, len(listingPorts))
    for port := range listingPorts {
        sortedPorts = append(sortedPorts, port)
    }

    sort.Slice(sortedPorts, func(i, j int) bool {
        return listingPorts[sortedPorts[i]] > listingPorts[sortedPorts[j]]
    })
        count := 0
        ports := ""
       for _, name := range sortedPorts {
           if count == number {
                break
           }else if ports != ""{
                ports += ","
           }
           ports +=  name
           count += 1
    }
    return ports

}