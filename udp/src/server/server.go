package main

import (
	"fmt"
	"net"
	"os"
)

func getIP() net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4
		}
	}
	return nil
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	ip := getIP()
	if ip != nil {
		msg := "Hi client, I am: 1-" + ip.String()
		_, err := conn.WriteToUDP([]byte(msg), addr)
		if err != nil {
			fmt.Printf("Couldn't send response %v", err)
		}
	}
	fmt.Printf("Couldn't find IP to broadcast")
}

func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: 30001,
		IP:   net.ParseIP("192.168.1.35"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr)
	}
}
