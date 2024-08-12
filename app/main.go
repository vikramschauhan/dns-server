package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		//receivedData := string(buf[:size])
		//fmt.Println("Received %d bytes from %s %s", size, source, receivedData)

		header := dns.ParseHeader(buf[:size])
		fmt.Println("Header:", header)
		questions := dns.ParseQuestion(buf[:size])
		fmt.Println("Questions:", questions)
		answers := dns.ParseAnswer(buf[:size])
		fmt.Println("Answers:", answers)

		response := dns.Message{
			Header:    header,
			Questions: questions,
			Answers:   answers,
		}
		_, err = udpConn.WriteToUDP(response.Bytes(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
