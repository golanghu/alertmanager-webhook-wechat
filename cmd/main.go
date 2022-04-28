package main

import (
	"fmt"
	//log "github.com/sirupsen/logrus"
	"github.com/golanghu/alerthook/pkg/apiserver"
	"github.com/golanghu/alerthook/pkg/config"
	log "github.com/golanghu/alerthook/pkg/logger"
)

var (
	gitRev    = ""
	buildDate = ""
	goVersion = ""
)
var flags *config.CMDFlags

func main() {
	flags = &config.CMDFlags{}
	flags.Init()

	if flags.Version {
		fmt.Printf("Build Platform: %s\n", goVersion)
		fmt.Printf("Git Rev: %s\n", gitRev)
		fmt.Printf("Build Date: %s\n", buildDate)
		return
	}

	log.SetLevel(flags.LogLevel)
	log.SetLogFile(flags.LogDir)

	// Run webhook server.
	apiserver.Run(flags)
}
