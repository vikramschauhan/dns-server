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
	RData    []byte //Resource Data
}

func (answer Answer) Bytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, encodeData(answer.Name))
	binary.Write(&buf, binary.BigEndian, answer.Type)
	binary.Write(&buf, binary.BigEndian, answer.Class)
	binary.Write(&buf, binary.BigEndian, answer.TTL)
	binary.Write(&buf, binary.BigEndian, answer.RDLength)
	binary.Write(&buf, binary.BigEndian, answer.RData)
	return buf.Bytes()
}

func ParseAnswer(data []byte) []Answer {
	answers := make([]Answer, 0)
	name := decodeDomainName(data, 12)
	for i := 0; i < int(binary.BigEndian.Uint16(data[4:6])); i++ {
		answer := Answer{
			Name:     name,
			Type:     uint16(1),
			Class:    uint16(1),
			TTL:      60,
			RDLength: uint16(4),
			RData:    []byte("\x08\x08\x08\x08"),
		}
		answers = append(answers, answer)
	}
	return answers
}
