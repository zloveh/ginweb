package util

import (
	"github.com/Unknwon/goconfig"
)

func NewConfig(configFile string) (*goconfig.ConfigFile, error) {
	return goconfig.LoadConfigFile(configFile)
}
