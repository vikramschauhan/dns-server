package dns

import (
	"bytes"
	"encoding/binary"
	"strings"
)

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func (question Question) Bytes() []byte {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, encodeDomainName(question.Name))
	binary.Write(&buf, binary.BigEndian, question.Type)
	binary.Write(&buf, binary.BigEndian, question.Class)
	return buf.Bytes()
}

func encodeDomainName(domain string) []byte {
	domainParts := strings.Split(domain, ".")
	buf := make([]byte, 0)
	for _, part := range domainParts {
		buf = append(buf, byte(len(part)))
		buf = append(buf, []byte(part)...)
	}
	buf = append(buf, 0)
	return buf
}
