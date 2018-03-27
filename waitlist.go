package main

type WaitListEntry struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Level    string `json:"level"`
	Sessions []int  `json:"sessions"`
}

func (wle *WaitListEntry) save() error {
	return nil
}
