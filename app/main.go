package main

import (
	"fmt"
	"github.com/codecrafters-io/dns-server-starter-go/dns"
	"net"
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

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		questions := []dns.Question{
			{
				Name:  "codecrafters.io",
				Type:  1,
				Class: 1,
			},
		}

		answers := []dns.Answer{
			{
				Name:     "codecrafters.io",
				Type:     1,
				Class:    1,
				TTL:      60,
				RDLength: 4,
				RData:    "8.8.8.8",
			},
		}

		header := dns.Header{
			ID:      1234,
			QR:      1,
			OPCODE:  0,
			AA:      0,
			TC:      0,
			RD:      0,
			RA:      0,
			Z:       0,
			RCODE:   0,
			QDCOUNT: uint16(len(questions)),
			ANCOUNT: uint16(len(answers)),
			NSCOUNT: 0,
			ARCOUNT: 0,
		}

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
