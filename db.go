package easy_serve

import (
	"fmt"

	"github.com/Sunqi43797189/easy_serve/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbMap = make(map[string]*dbobj)

type dbobj struct {
	db   *gorm.DB
	err  error
	name string
}

func initDB() {
	for _, conf := range config.C.DB {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.UserName,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database)
		var dbobj = dbobj{name: conf.Name}
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			dbobj.err = err
		} else {
			dbobj.db = db
		}
		dbMap[conf.Name] = &dbobj
	}
}

func (o *dbobj) close() {
	err := o.db.Close()
	if err != nil {
		fmt.Printf("db exit failed, name: %s, err: %v\n", o.name, err)
	} else {
		fmt.Printf("db exited, name: %s\n", o.name)
	}
}
