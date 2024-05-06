package dns

import "strings"

type Message struct {
	Header    Header
	Questions []Question
	Answers   []Answer
}

func (message Message) Bytes() []byte {
	headerBytes := message.Header.Bytes()
	questionBytes := make([]byte, 0)
	for _, question := range message.Questions {
		questionBytes = append(questionBytes, question.Bytes()...)
	}
	answerBytes := make([]byte, 0)
	for _, answer := range message.Answers {
		answerBytes = append(answerBytes, answer.Bytes()...)
	}
	return append(append(headerBytes, questionBytes...), answerBytes...)
}

func encodeData(data string) []byte {
	parts := strings.Split(data, ".")
	buf := make([]byte, 0)
	for _, part := range parts {
		buf = append(buf, byte(len(part)))
		buf = append(buf, []byte(part)...)
	}
	buf = append(buf, 0)
	return buf
}
