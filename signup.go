package main

type Session struct {
	Name     string `json:"name" binding:"required,max=50"`
	Birthday string `json:"birthday binding:"max=50"`
	Level    string `json:"level binding:"max=500"`
	Sessions []int  `json:"sessions" binding:"required,max=10"`
}

type SignUpForm struct {
	ID       string     `json:"id"`
	Name     string     `json:"name" binding:"required,max=50"`
	Phone    string     `json:"phone" binding:"required,max=50"`
	Email    string     `json:"email" binding:"required,max=50"`
	Token    string     `json:"token" binding:"required,max=50"`
	Cost     int        `json:"cost" binding:"required,max=9999"`
	Sessions []*Session `json:"sessions" binding:"required,max=10"`
	Request  string     `json:"request" binding:"max=500"`
}

func getAllSignups() ([]*SignUpForm, error) {
	return []*SignUpForm{}, nil
}

func (sf *SignUpForm) save() error {
	return nil
}

func (sf *SignUpForm) update() error {
	return nil
}
