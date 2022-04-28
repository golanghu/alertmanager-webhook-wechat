package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Notifier implements a Notifier for generic webhooks.
type WeixinWebHook struct {
	address string
	//tmpl    *template.Template

	client *http.Client
}

// New returns a new Webhook.
func New(addr string) (*WeixinWebHook, error) {
	return &WeixinWebHook{
		address: addr,
		//tmpl:    t,
		client: http.DefaultClient,
	}, nil
}

// Notify implements the Notifier interface.
func (n *WeixinWebHook) Notify(data interface{}) error {

	var buf []byte
	var err error
	if buf, err = json.Marshal(data); err != nil {
		//logrus.Errorf("sendToWeixin.err:%s",err)
		return err
	}
	var req *http.Request

	req, err = http.NewRequest("POST",
		n.address,
		bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
