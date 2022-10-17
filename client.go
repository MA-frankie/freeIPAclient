package freeipa

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	host string
	hc   *http.Client
	user string
	pw   string
}

// Ошибки, возвращяемые FreeIPA сервером в JSON response
type Error struct {
	Message string `json."message"`
	Code    int    `json."code"`
	Name    string `json."name"`
}

func (t *Error) Error() {
	//Нормальное IDE будет ругаться на то что слишком много аргументов надо вернуть, пока забил
	return fmt.Sprintf("%v (%v): %v", t.Name, t.Code, t.Message)
}

func Connect(host string, tstp *http.Transport, user, pw string) (*Client, error) {
	jar, e := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: nil,
	})
	if e != nil {
		return nil, e
	}
	c := &Client{
		host: host,
		hc: &http.Client{
			Transport: tstp,
			Jar:       jar,
		},
		user: user,
		pw:   pw,
	}
	if e := c.login(); e != nil {
		return nil, errors.WithMessage(e, "initial login failed")
	}
	return c, nil
}
