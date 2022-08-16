package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	ServeType_HTTP = "http"
	ServeType_GRPC = "grpc"
	ServeType_CRON = "cronjob"
	ServeEnv_Pro   = "pro"
	ServeEnv_Test  = "test"
)

type Config struct {
	Service Service `json:"service" yaml:"service"`
	Redis   []Redis `json:"redis" yaml:"redis"`
	DB      []DB    `json:"db" yaml:"db"`
	Log     []Log   `json:"log" yaml:"log"`
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
	Name            string `json:"name" yaml:"name"`
	Host            string `json:"host" yaml:"host"`
	Port            int    `json:"port" yaml:"port"`
	Database        string `json:"database" yaml:"database"`
	UserName        string `json:"username" yaml:"username"`
	Password        string `json:"password" yaml:"password"`
	ConnMaxIdleTime int    `json:"max_idle_time" yaml:"max_idle_time"`
	ConnMaxLifeTime int    `json:"max_life_time" yaml:"max_life_time"`
	MaxIdleConn     int    `json:"max_idle_conn" yaml:"max_idle_conn"`
	MaxLifeConn     int    `json:"max_life_conn" yaml:"max_life_conn"`
	ConnectTimeout  int    `json:"connect_timeout" yaml:"connect_timeout"`
	ReadTimeout     int    `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout    int    `json:"write_timeout" yaml:"write_timeout"`
}

type Log struct {
	Name     string `json:"name" yaml:"name"`
	LogPath string `json:"logpath" yaml:"logpath"`
	Level    string `json:"level" yaml:"level"`
	Compress bool   `json:"compress" yaml:"compress"`
	MaxSize  int    `json:"max_size" yaml:"max_size"`
}

var C = &Config{}

func InitConf(configFile string, customConfig interface{}) error {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(conf, C)
	if err != nil {
		return err
	}

	if customConfig != nil {
		err = yaml.Unmarshal(conf, customConfig)
		if err != nil {
			return err
		}
	}
	return nil
}
