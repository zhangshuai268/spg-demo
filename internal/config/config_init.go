package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Conf *Config

func InitConfig(paths ...string) (*Config, error) {
	var path string
	if len(paths) == 0 {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(dir, "config.json")
		if os.Getenv("STAGE") != "" {
			path = filepath.Join(dir, "config_"+os.Getenv("STAGE")+".json")
		}
	} else {
		path = paths[0]
	}

	config := new(Config)
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("打开配置文件错误" + path + err.Error())
	}

	confByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("读取配置文件错误" + err.Error())
	}

	err = json.Unmarshal(confByte, config)
	if err != nil {
		return nil, errors.New("读取配置文件错误" + err.Error())
	}

	Conf = config

	return Conf, nil
}
