package format

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CIDRToUint8(s string) [4]uint8 {
	var nm [4]uint8
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
	return nm
}

func Uint8ToCIDR(addr [4]uint8) int {
	count := 0
	for i := 0; i < 4; i++ {
		octetarr := strings.Split(fmt.Sprintf("%b", addr[i]), "")
		count += calculateOne(octetarr)
	}
	return count
}
func calculateOne(s []string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "1" {
			count += 1
		}
	}
	return count
}
