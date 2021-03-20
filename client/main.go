package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
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
	FortuneNonce int
}

type FortuneMessage struct {
	Fortune string
}

// Hash function to convert Nonce + Secret
func newSHA256(nonce int, secret string) []byte {
	nonceByte := []byte(strconv.Itoa(nonce))
	bytestring := append(nonceByte, []byte(secret)...)
	hash := sha256.Sum256(bytestring)
	return hash[:]
}

func main() {

	arguments := os.Args[1:]

	if len(arguments) != 3 {
		fmt.Println("Usage: go run main.go [local UDP ip:port] [aserver ip:port] [secret]")
		os.Exit(1)
	}

	localAddr := arguments[1]
	aserverAddr := arguments[2]
	secretMsg := arguments[3]

	nonceReply := make([]byte, 1024)
	conn, err := net.Dial("udp", aserverAddr)
	if err != nil {
		log.Fatal("Something went wrong:", err)
		return
	}

	fmt.Fprintf(conn, "Arb Message")

	_, err = bufio.NewReader(conn).Read(nonceReply)
	byteToInt, err := strconv.Atoi(string(nonceReply))
	if err != nil {
		log.Fatal("Could not convert string ASCII to int: ", err)
	}
	fortNonce := &FortuneReqMessage{
		FortuneNonce: byteToInt,
	}

	newHash := newSHA256(fortNonce.FortuneNonce, secretMsg)

	fmt.Printf("Nonce: %d\n", fortNonce.FortuneNonce)
}
