package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type ErrMessage struct {
	Error string
}

type NonceMessage struct {
	Nonce int64
}

type HashMessage struct {
	Hash string
}

type FortuneInfoMessage struct {
	FortuneServer string
	FortuneNonce  int64
}

type FortuneReqMessage struct {
	FortuneNonce uint64
}

type FortuneMessage struct {
	Fortune string
}

// Hash function
func newSHA256(data []byte, nonce uint64) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func main() {

	nonceReply := make([]byte, 1024)

	conn, err := net.Dial("udp", ":1234")
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}

	fmt.Fprintf(conn, "Arb Message")

	_, err = bufio.NewReader(conn).Read(nonceReply)
	fortNonce := &FortuneReqMessage{
		FortuneNonce: binary.BigEndian.Uint64(nonceReply),
	}

	fmt.Printf("Nonce: %d\n", fortNonce.FortuneNonce)
}
