package easyserve

import (
	"flag"
	"fmt"

	"github.com/Sunqi43797189/easy_serve/config"
)

var configFile = flag.String("i", "config.yaml", "config file")

func New() {
	flag.Parse()
	config.InitConf(*configFile)
	fmt.Println(config.C)
}

func Serve(){
	
}
