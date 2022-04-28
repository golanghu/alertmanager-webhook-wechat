package config

import (
	"fmt"
	"testing"
)

func TestCMDFlags_Init(t *testing.T) {
	f := &CMDFlags{}
	f.Init()
	f.LogLevel = "debug"

	fmt.Println(f.LogLevel)
}
