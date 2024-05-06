package dns

type Message struct {
	Header    Header
	Questions []Question
}

func (message Message) Bytes() []byte {
	headerBytes := message.Header.Bytes()
	questionBytes := make([]byte, 0)
	for _, question := range message.Questions {
		questionBytes = append(questionBytes, question.Bytes()...)
	}
	return append(headerBytes, questionBytes...)
}
