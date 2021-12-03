package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dunefro/network_identifier/format"
	"github.com/dunefro/network_identifier/net"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No IP was passed")
		fmt.Println("Usage: network_identifier IP/NETMASK")
		os.Exit(1)
	}
	fullip := strings.Split(os.Args[1], "/")
	addr := getAddress(fullip[0])
	netmask := getNetmask(fullip[1])
	fmt.Println("Network Prefix: ", format.UintToAddr(net.GetNetwork(addr, netmask)))
	fmt.Println("Host: ", format.UintToAddr(net.GetHost(addr, netmask)))
	fmt.Println("Broadcast address: ", format.UintToAddr(net.GetBroadcastAddr(addr, netmask)))
	fmt.Println("Total subnets: ", net.GetTotalSubnet(format.Uint8ToCIDR(netmask)))

}
func getNetmask(s string) [4]uint8 {
	if strings.Contains(s, ".") {
		return format.AddrToUint8(s)
	} else {
		return format.CIDRToUint8(s)
	}
}

func getAddress(s string) [4]uint8 {
	return format.AddrToUint8(s)
}
