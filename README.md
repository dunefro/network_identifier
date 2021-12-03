# network_identifier
Project helps to identify the network, broadcast address and no of possible hosts for Ipv4 address

1. To use it directly as a go file
```
$ go run main.go IP/NETMASK
```
It will behave something like this
```
$ go run main.go 172.36.14.23/13
Network Prefix:  172.32.0.0
Host:  0.4.14.23
Broadcast address:  172.39.255.255
Total subnets:  524286
```
2. To build the package
```
$ go build
```
Then use the binary accroding to your `GOBIN` variable
```
$ ./network_identifier 172.36.14.22/255.255.0.0
Network Prefix:  172.36.0.0
Host:  0.0.14.22
Broadcast address:  172.36.255.255
Total subnets:  65534
```