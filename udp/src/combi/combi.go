package main

import (
	"log"
	"net"
	"time"
)

func broadcast_ip() {
	c, err := net.ListenPacket("udp", ":0")

	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	dst, err := net.ResolveUDPAddr("udp", "255.255.255.255:8032")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if _, err := c.WriteTo([]byte("hello"), dst); err != nil {
			log.Fatal(err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func listenUDP() {
	c, err := net.ListenPacket("udp", ":8032")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	for {
		b := make([]byte, 512)
		n, peer, err := c.ReadFrom(b)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, "bytes read from", peer, "saying", string(b[0:n]))
	}
}

func main() {
	go broadcast_ip()
	go listenUDP()

	time.Sleep(100000 * time.Millisecond)

}
