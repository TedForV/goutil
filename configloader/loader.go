package configloader

import (
	"fmt"
	"github.com/jinzhu/configor"
)

const (
	_ = iota
	CONFIG_JSON
	CONFIG_YAML
	CONFIG_XML
)

func Load(configPath string, configFileType int, config interface{}) error {
	switch configFileType {
	case CONFIG_YAML:
		return loadYaml(configPath, config)
	default:
		config = nil
		return fmt.Errorf("path:%s,type:%d, config can't be found.", configPath, configFileType)
	}
}

func loadYaml(configPath string, config interface{}) error {
	return configor.Load(config, configPath)
}
