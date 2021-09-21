package goextron

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	fmt.Println(string(a))
	return nil
}

func (g Goextron) SetMute(muted bool) error {
	muteReq := "[{\"uri\":\"/video/in/channel/2/active_input\",\"value\":5}, {\"uri\":\"/audio/dsp/out/mute\",\"value\":0}]"
	if muted {
		muteReq = "[{\"uri\":\"/video/in/channel/2/active_input\",\"value\":4}, {\"uri\":\"/audio/dsp/out/mute\",\"value\":1}]"
	}
	body := strings.NewReader(muteReq)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%sapi/swis/resources", g.addr), body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic: %s", g.auth))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	var res []defaultResp
	respStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respStr ))
	err = json.Unmarshal(respStr, &res)
	if err != nil {
		return fmt.Errorf("unexpected response from smp: %v", err)
	}
	return nil
}

