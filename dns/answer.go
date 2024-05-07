package dns

import (
	"bytes"
	"encoding/binary"
)

type Answer struct {
	Name     string //Domain Name
	Type     uint16 //Resource Record Type
	Class    uint16 //Resource Record Class
	TTL      uint32 //Time to Live
	RDLength uint16 //Resource Data Length
	RData    string //Resource Data
}

func (answer Answer) Bytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, encodeData(answer.Name))
	binary.Write(&buf, binary.BigEndian, answer.Type)
	binary.Write(&buf, binary.BigEndian, answer.Class)
	binary.Write(&buf, binary.BigEndian, answer.TTL)
	binary.Write(&buf, binary.BigEndian, answer.RDLength)
	binary.Write(&buf, binary.BigEndian, encodeData(answer.RData))
	return buf.Bytes()
}

func ParseAnswer(data []byte) Answer {
	answer := Answer{}
	answer.Name = decodeDomainName(data, 12)
	answer.Type = uint16(1)
	answer.Class = uint16(1)
	answer.TTL = binary.BigEndian.Uint32(data[len(answer.Name)+5 : len(answer.Name)+9])
	answer.RDLength = uint16(4)
	answer.RData = "8.8.8.8"

	return answer
}
