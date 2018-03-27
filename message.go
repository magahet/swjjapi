package main

type Message struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Level    string `json:"level"`
	Sessions []int  `json:"sessions"`
}

func (msg *Message) send() error {
	return nil
}
