package notifier

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Payload struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconUrl   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

type Notifier struct {
	webhook_url string
	Payload
}

func New(webhook_url string) *Notifier {
	return &Notifier{webhook_url, Payload{}}
}

func (n Notifier) Ping(contents string) string {
	n.Text = contents
	params, _ := json.Marshal(n.Payload)

	resp, _ := http.PostForm(
		n.webhook_url,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(body)
}
