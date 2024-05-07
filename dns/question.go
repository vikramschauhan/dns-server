package dns

import (
	"bytes"
	"encoding/binary"
)

type Question struct {
	Name  string //Domain Name
	Type  uint16 //Resource Record Type
	Class uint16 //Resource Record Class
}

func (question Question) Bytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, encodeData(question.Name))
	binary.Write(&buf, binary.BigEndian, question.Type)
	binary.Write(&buf, binary.BigEndian, question.Class)
	return buf.Bytes()
}

func ParseQuestion(data []byte) Question {
	question := Question{}
	question.Name = decodeDomainName(data, 12)
	question.Type = uint16(1)
	question.Class = uint16(1)
	return question
}
