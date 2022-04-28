package config

import (
	"flag"
	"github.com/spf13/viper"
	"strings"
)

// CMDFlags are the flags used by the cmd
type CMDFlags struct {
	Version    bool
	ListenAddr string
	LogLevel   string
	LogDir     string

	//PrometheusURL string
	//PrometheusEtc string
	WebhookAddr string

	DBName     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
}

var Cfg *CMDFlags

// Init initializes and parse the flags
func (c *CMDFlags) Init() {

	// version flags
	flag.BoolVar(&c.Version, "v", false, "show version (short for -version) ")
	flag.StringVar(&c.ListenAddr, "listen", ":8080", "Addr to listen on for api server. ")
	flag.StringVar(&c.LogLevel, "log-level", "info", "Set logger level. ")
	flag.StringVar(&c.LogDir, "log-dir", "", "Given a dir to flush logs. ")
	flag.StringVar(&c.WebhookAddr, "web-hook", "", "Addr to send notify to weixin.")

	flag.Parse()

	// init viper env.
	viper.AutomaticEnv()
	// overwrite by env.
	if s := viper.GetString("listen"); s != "" {
		c.ListenAddr = s
	}
	if s := viper.GetString("log_level"); s != "" {
		c.LogLevel = strings.ToLower(s)
	}

	if s := viper.GetString("log_dir"); s != "" {
		c.LogDir = s
	}

	if s := viper.GetString("web_hook"); s != "" {
		c.WebhookAddr = s
	}

	Cfg = c

}
