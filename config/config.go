package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	ServeType_HTTP = "http"
	ServeType_GRPC = "grpc"
	ServeEnv_Pro   = "pro"
	ServeEnv_Test  = "test"
)

type Config struct {
	Service Service `json:"service" yaml:"service"`
	Redis   []Redis `json:"redis" yaml:"redis"`
	DB      []DB    `json:"db" yaml:"db"`
}

type Service struct {
	Env       string `json:"env" yaml:"env"`
	ServeType string `json:"serve_type" yaml:"serve_type"`
	ServePort int    `json:"serve_port" yaml:"serve_port"`
}

type Redis struct {
	Name     string `json:"name" yaml:"name"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	DB       int    `json:"db" yaml:"db"`
	UserName string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

type DB struct {
	Name     string `json:"name" yaml:"name"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
	UserName string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
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
