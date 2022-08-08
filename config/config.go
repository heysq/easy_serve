package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	ServeType_HTTP = "http"
	ServeType_GRPC = "grpc"

	ServeEnv_Pro = "pro"
)


type Config struct {
	Service Service `json:"service" yaml:"service"`
}

type Service struct {
	Env       string `json:"env" yaml:"env"`
	ServeType string `json:"serve_type" yaml:"serve_type"`
	ServePort int `json:"serve_port" yaml:"serve_port"`
}

var C = &Config{}

func InitConf(configFile string) error {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(conf, C)
	if err != nil {
		return err
	}
	return nil
}
