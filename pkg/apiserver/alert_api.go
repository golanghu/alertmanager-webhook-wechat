// Package apiserver
/*
接收alertmanager发送的告警信息，根据规则进行过滤，然后发送给企业微信群
*/
package apiserver

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golanghu/alerthook/pkg/notify/webhook"
	"net/http"
	"time"
)

type AlertsReceiveReq struct {
	Receiver string
	Status   string
	Alerts   []Alert
}

type Alert struct {
	Status string            `json:"status"`
	Labels map[string]string `json:"labels"`

	// Extra key/value information which does not define alert identity.
	Annotations map[string]string `json:"annotations"`

	// The known time range for this alert. Both ends are optional.
	StartsAt    time.Time `json:"startsAt,omitempty"`
	EndsAt      time.Time `json:"endsAt,omitempty"`
	FingerPrint string    `json:"fingerprint"`
}

func AlertReceive(c *gin.Context) {
	var req AlertsReceiveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("AlertReceive.err:%s", err)
		c.JSON(http.StatusOK, nil)
		return
	}

	logrus.Debugf("receive alert:%+v", req)

	//解析appname、rulename
	for _, value := range req.Alerts {
		sendToWeixin(value)
	}
	c.JSON(http.StatusOK, nil)
	return
}

func sendToWeixin(alert Alert) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("sendToWeixin.err:%s", err)
		}
	}()

	data, err := GetMsgToWeixin(alert)
	if err != nil {
		logrus.Errorf("sendToWeixin.err:%s", err)
		return
	}

	if err := weixinWebHook.Notify(data); err != nil {
		logrus.Errorf("notify to weixin err:%s", err)
	}
	return
}

func GetMsgToWeixin(alert Alert) (*webhook.WexinData, error) {

	weiXinData := &webhook.WexinData{
		MsgType: webhook.TEXT,
	}
	data := webhook.WeixinText{
		MentionedList: []string{"@all"},
	}

	msgData := WeiXinMsg{
		Status:     alert.Status,
		AlertLevel: alert.Labels["severity"],
		AlertName:  alert.Labels["alertname"],
		Instance:   alert.Labels["instance"],
		StartAt:    alert.StartsAt.Format("2006-01-02 15:04:05"),
		EndAt:      alert.EndsAt.Format("2006-01-02 15:04:05"),
		AlertMsg:   alert.Annotations["summary"],
	}

	var b []byte
	a := bytes.NewBuffer(b)
	if err := webhook.MsgTpl.Execute(a, msgData); err != nil {
		logrus.Error("weixin msg tpl execute error:", err)
		return nil, err
	}

	data.Content = a.String()
	weiXinData.Text = data
	return weiXinData, nil

}

type WeiXinMsg struct {
	Status     string
	AlertLevel string
	AlertName  string
	Instance   string
	StartAt    string
	EndAt      string
	AlertMsg   string
}
