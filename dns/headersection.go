package dns

import "encoding/binary"

type Header struct {
	ID      uint16 //Packet Identifier
	QR      uint8  //Query/Response Indicator
	OPCODE  uint8  //Operation Code
	AA      uint8  //Authoritative Answer
	TC      uint8  //Truncation
	RD      uint8  //Recursion Desired
	RA      uint8  //Recursion Available
	Z       uint8  //Reserved
	RCODE   uint8  //Response Code
	QDCOUNT uint16 //Question Count
	ANCOUNT uint16 //Answer Record Count
	NSCOUNT uint16 //Name Server Count
	ARCOUNT uint16 //Additional Record Count
}

func (header Header) Bytes() []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:2], header.ID)
	flag := uint16(header.QR)<<15 |
		uint16(header.OPCODE)<<11 |
		uint16(header.AA)<<10 |
		uint16(header.TC)<<9 |
		uint16(header.RD)<<8 |
		uint16(header.RA)<<7 |
		uint16(header.Z)<<4 |
		uint16(header.RCODE)
	binary.BigEndian.PutUint16(buf[2:4], flag)
	binary.BigEndian.PutUint16(buf[4:6], header.QDCOUNT)
	binary.BigEndian.PutUint16(buf[6:8], header.ANCOUNT)
	binary.BigEndian.PutUint16(buf[8:10], header.NSCOUNT)
	binary.BigEndian.PutUint16(buf[10:12], header.ARCOUNT)
	return buf
}
