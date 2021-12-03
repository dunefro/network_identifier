package net

import (
	"math"
)

func GetNetwork(addr [4]uint8, mask [4]uint8) [4]uint8 {
	network := [4]uint8{}
	for i := 0; i < 4; i++ {
		network[i] = addr[i] & mask[i]
	}
	return network
}

func GetHost(addr [4]uint8, mask [4]uint8) [4]uint8 {
	host := [4]uint8{}
	for i := 0; i < 4; i++ {
		host[i] = addr[i] & (^mask[i])
	}
	return host
}

func GetBroadcastAddr(addr [4]uint8, mask [4]uint8) [4]uint8 {
	brd := [4]uint8{}
	for i := 0; i < 4; i++ {
		brd[i] = addr[i] | (^mask[i])
	}
	return brd
}
func GetTotalSubnet(cidr int) int {

	return int(math.Pow(float64(2), float64(32-cidr))) - 2
}
