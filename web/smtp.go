package web

import (
	"net/smtp"
)

type Poster struct {
	client *smtp.Client
}

func NewPoster(addr string) *Poster {
	p := &Poster{}
	client, err := smtp.Dial(addr)
	if err != nil {
		panic(err)
	}
	p.client = client
	return p
}

func (p *Poster) SendEmail() {

}
