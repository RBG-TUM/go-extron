package goextron

type defaultResp struct {
	Meta struct {
		Uri string `json:"uri"`
	} `json:"meta"`
	Result int `json:"result"`
}
