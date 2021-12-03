package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func validOctet(n int64) (uint8, error) {
	if (n < 0) || (n > 255) {
		fmt.Println(n, "not allowed in the octet")
		return 0, fmt.Errorf("Values allowed in the octect are 0-255, %d doesn't fall in range", n)
	}
	return uint8(n), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No IP was passed")
		fmt.Println("Usage: network_identifier IP/NETMASK")
		os.Exit(1)
	}
	fullip := strings.Split(os.Args[1], "/")
	var addr [4]uint8
	addr = getAddress(fullip[0])

	netmask := getNetmask(fullip[1])
	fmt.Println(addr, netmask)

}
func getNetmask(s string) [4]uint8 {
	var nm [4]uint8
	if strings.Contains(s, ".") {
		netmaskStr := strings.Split(s, ".")
		if len(netmaskStr) != 4 {
			fmt.Println("Invalid netmask")
			os.Exit(2)
		}
		for i, octet := range netmaskStr {
			val, err := strconv.ParseInt(octet, 10, 16)
			if err != nil {
				fmt.Println("Ip address invalid")
				fmt.Println()
				os.Exit(2)
			}
			validOctet, err := validOctet(val)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			nm[i] = validOctet
		}
	} else {
		cidr, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if cidr > 32 || cidr < 0 {
			fmt.Println("Impossible netmask. Range 0-32")
			os.Exit(2)
		}

		for i := 0; cidr > 0; i++ {
			if cidr < 8 {
				nm[i] = uint8(256 - int(math.Pow(float64(2), float64(8-cidr))))
			} else {
				nm[i] = uint8(255)
			}
			cidr = cidr - 8
		}
	}
	return nm
}

func getAddress(s string) [4]uint8 {
	var addr [4]uint8
	addrStr := strings.Split(s, ".")
	if len(addrStr) != 4 {
		fmt.Println("Invalid Address")
		os.Exit(1)
	}
	for i, octet := range addrStr {
		val, err := strconv.ParseInt(octet, 10, 16)
		if err != nil {
			fmt.Println("Ip address invalid")
			fmt.Println()
			os.Exit(2)
		}
		validOctet, err := validOctet(val)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		addr[i] = validOctet
	}
	return addr
}
