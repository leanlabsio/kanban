package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg         *ini.File
	CustomPath  string
	App_Version string
)

// NewContext created new context for settings
func NewContext() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatal(4, "Fail to parse 'conf/app.ini': %v", err)
	}

	if CustomPath != "" {
		Cfg.Append(CustomPath)
	}
}
