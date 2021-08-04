package goextron

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Goextron struct {
	addr string
	auth string
}

func New(addr string, user string, password string) Goextron {
	return Goextron{addr: addr, auth: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, password)))}
}

func (g Goextron) GetAuthStatus() error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%sapi/swis/session/info?uri=", g.addr), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic: %s", g.auth))
	resp, _ := client.Do(req)
	a, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	println(string(a))
	return nil
}
