package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

type FortuneNonce struct {
	mu    sync.Mutex
	nonce int
}

// Gorountine function to read packets from UDP in separate threads
func listen(conn *net.UDPConn, quit chan struct{}) {
	buffer := make([]byte, 1024)
	n, remoteAddr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
	}
}

func main() {

	arguments := os.Args
	PORT := arguments[1]

	s, err := net.ResolveUDPAddr("udp4", PORT)

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}

	defer conn.Close()

	fortuneNonce := &FortuneNonce{nonce: 0}

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("Error occured while reading buffer:", err)
		}
		fmt.Println("Recieved Arb Message:", string(buf[0:n-1]))

		// Critical section start
		fortuneNonce.mu.Lock()
		nonceToSend := []byte(strconv.Itoa(fortuneNonce.nonce))
		fortuneNonce.nonce++
		fortuneNonce.mu.Unlock()
		// Critical section end

		_, err = conn.WriteToUDP(nonceToSend, addr)
		if err != nil {
			log.Fatal("Error occured while sending message back:", err)
			return
		}

	}

}
