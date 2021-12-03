package format

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UintToAddr(ip [4]uint8) string {
	var ipstr [4]string
	for i, v := range ip {
		ipstr[i] = fmt.Sprint(v)
	}
	return strings.Join(ipstr[:], ".")
}

func AddrToUint8(s string) [4]uint8 {
	var addr [4]uint8
	addrStr := strings.Split(s, ".")
	if len(addrStr) != 4 {
		fmt.Println("Invalid Address")
		os.Exit(1)
	}
	for i, octet := range addrStr {
		val, err := strconv.ParseInt(octet, 10, 16)
		if err != nil {
			fmt.Println("IP address invalid")
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
