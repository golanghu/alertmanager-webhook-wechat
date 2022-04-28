package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/golanghu/alerthook/pkg/config"
	"github.com/golanghu/alerthook/pkg/notify/webhook"
	"github.com/sirupsen/logrus"

	"net/http"
	"time"
)

var (
	weixinWebHook *webhook.WeixinWebHook
)

func Run(flags *config.CMDFlags) {

	var err error
	if flags.WebhookAddr == "" {
		logrus.Fatal("webhook start failed, ", "webhookaddr empty.")
	}
	weixinWebHook, err = webhook.New(flags.WebhookAddr)
	if err != nil {
		logrus.Fatal("init webhook err:%s", err)
	}
	logrus.Info("webhook listen on ", flags.ListenAddr)
	gin.SetMode(gin.ReleaseMode)
	handler := gin.Default()
	v1 := handler.Group("/v1")

	{
		v1.POST("/alert/receive", AlertReceive)
	}
	server := &http.Server{
		Addr:           flags.ListenAddr,
		Handler:        handler,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		logrus.Fatal("Monitor start failed, ", err)
	}
}
