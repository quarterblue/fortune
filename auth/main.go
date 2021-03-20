package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args

	PORT := arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	defer conn.Close()

	var nonce int = 0

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("Error occured while reading buffer:", err)
		}
		fmt.Println("Recieved Arb Message:", string(buf[0:n-1]))

		nonceToSend := []byte(strconv.Itoa(nonce))
		nonce++
		_, err = conn.WriteToUDP(nonceToSend, addr)
		if err != nil {
			log.Fatal("Error occured while sending message back:", err)
			return
		}

	}
}
