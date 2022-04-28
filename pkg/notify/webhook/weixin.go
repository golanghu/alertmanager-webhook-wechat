package webhook

import (
	"fmt"
	"os"
	"text/template"
)

const (
	TEXT = "text"
)

var MsgTpl *template.Template

type WexinData struct {
	MsgType string     `json:"msgtype"`
	Text    WeixinText `json:"text"`
}

type WeixinText struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

func init() {
	t, err := template.ParseGlob("./config/weixin.tpl")
	if err != nil {
		fmt.Println(" init message.tpl error:", err)
		os.Exit(1)
	}
	MsgTpl = t
}
