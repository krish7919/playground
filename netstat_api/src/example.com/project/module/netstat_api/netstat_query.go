package main

import (
	"fmt"
	"github.com/drael/GOnetstat"
	"net"
	"strings"
)

func main() {
	// get openvpn listen port
	processes := GOnetstat.Tcp()
	for _, process := range processes {
		if strings.ToLower(process.Name) == "openvpn" &&
			process.State == "LISTEN" {
			fmt.Println(process.State)
			fmt.Println(process.Name)
			fmt.Println("Port: ", process.Port)
			fmt.Println("Foreign: ", process.ForeignPort)
		}
	}

	// get openvpn network addr
	var ovpnIntfs []net.Interface
	var addresses []net.Addr
	var nwAddrs []net.IP
	var err error

	ovpnIntfs, err = net.Interfaces()
	if err != nil {
		fmt.Println("Error while getting intf list: ", err)
	}
	nwAddrs = []net.IP{}
	for _, intf := range ovpnIntfs {
		//if strings.HasPrefix(intf.Name, "box-server-" ) == true {
		if strings.HasPrefix(intf.Name, "box-server") == true {
			// get address/addresses assigned to this interface
			addresses, err = intf.Addrs()
			if err != nil {
				fmt.Println("Error while getting addr list: ", err)
			}
			for _, addr := range addresses {
				fmt.Println("net: ", addr.Network())
				fmt.Println("str: ", addr.String())
				x, y, z := net.ParseCIDR(addr.String())
				if z != nil {
					fmt.Println("Error while parsing addr: ", err)
				}
				fmt.Println("ip: ", x)
				fmt.Println("ipnetip: ", y.IP)
				fmt.Println("ipnetmask: ", y.Mask)
				// cache the nw addr
				nwAddrs = append(nwAddrs, y.IP)
			}
		}
	}
	fmt.Printf("addrs: %+v\n", nwAddrs)
	// since we know our server net is 10.x.0.0/16, find the next largest x
	// from the cached array and increment by 1 to get the next n/w addr
	var largestNw net.IP
	largestNw = nil
	for _, ip := range nwAddrs {
		if largestNw == nil {
			largestNw = ip
		} else {
			// is current ip larger then cached one?
			if int(largestNw[1]) < int(ip[1]) {
				largestNw = ip
			}
		}
	}
	fmt.Println("largest: ", largestNw)
	largestNw[1] = byte(int(largestNw[1]) + 1)
	fmt.Println("new ip: ", largestNw)
}
