package messages

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Type    string
	Payload string
}

func Unmarshal(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Printf("Error while unmarsharing %v: %v", data, err)
		return nil, err
	}

	return &msg, nil
}

func (m *Message) Marshal() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error while marsharing %v: %v", m, err)
		return nil, err
	}

	return data, nil
}
