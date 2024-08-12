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

func ParseQuestion(data []byte) []Question {
	questions := make([]Question, 0)
	for i := 0; i < int(binary.BigEndian.Uint16(data[4:6])); i++ {
		question := Question{
			Name:  decodeDomainName(data, 12),
			Type:  uint16(1),
			Class: uint16(1),
		}
		questions = append(questions, question)
	}
	return questions
}
