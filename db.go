package easy_serve

import (
	"fmt"

	"github.com/Sunqi43797189/easy_serve/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMap = make(map[string]*dbobj)

type dbobj struct {
	db  *gorm.DB
	err error
}

func initDB() {
	for _, conf := range config.C.DB {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.UserName,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database)
		var dbobj dbobj
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			dbobj.err = err
		} else {
			dbobj.db = db
		}
		dbMap[conf.Name] = &dbobj
	}
}
