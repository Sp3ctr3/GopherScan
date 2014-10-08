package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	var ports_list = map[int]string{
		21:  "ftp",
		22:  "SSH",
		23:  "telnet",
		443: "https",
		80:  "http",
	}
	var multiple bool = false
	if len(os.Args) > 0 {
		fmt.Println("Scanning ", os.Args[1])
	} else {
		fmt.Println("Enter a hostname")
		os.Exit(1)
	}
	a := os.Args[1]
	ports := [...]int{21, 80, 443}
	addrs, _ := net.LookupHost(a)
	if len(addrs) > 1 {
		fmt.Println("Multiple IPs found for the host")
		multiple = true
	}
	for _, addr := range addrs {
		if multiple {
			fmt.Println("For IP", addr)
		}
		for _, element := range ports {
			tcpaddr, _ := net.ResolveTCPAddr("tcp4", addr+":"+strconv.Itoa(element))
			_, err := net.DialTCP("tcp", nil, tcpaddr)
			if err == nil {
				fmt.Println(element, "Port is open: ", ports_list[element])
			} else {
				fmt.Println(element, "Closed")
			}
		}
	}
}
