package weapp

import "github.com/medivhzhan/weapp/v2"

type Client struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

func NewClient(appid, sercret string) *Client {
	return &Client{
		Appid:  appid,
		Secret: sercret,
	}
}
func (c *Client) Code2Session(code string) (r *weapp.LoginResponse, err error) {
	r, err = weapp.Login(c.Appid, c.Secret, code)
	return
}
