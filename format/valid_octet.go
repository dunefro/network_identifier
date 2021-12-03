package format

import "fmt"

func validOctet(n int64) (uint8, error) {
	if (n < 0) || (n > 255) {
		fmt.Println(n, "not allowed in the octet")
		return 0, fmt.Errorf("values allowed in the octect are 0-255, %d doesn't fall in range", n)
	}
	return uint8(n), nil
}
