package configloader

import (
	"errors"
	"fmt"
	"github.com/TedForV/goutil/filesystem"
	"github.com/jinzhu/configor"
)

// define the config file type
const (
	_ = iota
	CONFIG_JSON
	CONFIG_YAML
	CONFIG_XML
)

// Load is a func that load config file,and convert to a certain model
func Load(configPath string, configFileType int, config interface{}) error {
	existed, err := filesystem.IsPathExisted(configPath)
	if err != nil {
		return err
	}
	if !existed {
		return errors.New("configPath is wrong.")
	}
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
