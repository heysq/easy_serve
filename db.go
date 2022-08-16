package easy_serve

import (
	"fmt"
	"time"

	"github.com/Sunqi43797189/easy_serve/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMap = make(map[string]*dbobj)

type dbobj struct {
	db   *gorm.DB
	err  error
	name string
}

func initDB() {
	for _, conf := range config.C.DB {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds&readTimeout=%ds&writeTimeout=%ds",
			conf.UserName,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
			conf.ConnectTimeout,
			conf.ReadTimeout,
			conf.WriteTimeout,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		var dbobj = dbobj{name: conf.Name}
		if err != nil {
			dbobj.err = err
		} else {
			dbobj.db = db
		}
		sqlx, err := db.DB()
		if err != nil {
			dbobj.err = err
		}
		sqlx.SetConnMaxIdleTime(time.Duration(conf.ConnMaxIdleTime))
		sqlx.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime))
		sqlx.SetMaxIdleConns(conf.MaxIdleConn)
		sqlx.SetMaxOpenConns(conf.MaxLifeConn)
		dbMap[conf.Name] = &dbobj
	}
}

func (o *dbobj) close() {
	sqlx, err := o.db.DB()
	err = sqlx.Close()
	if err != nil {
		fmt.Printf("db exit failed, name: %s, err: %v\n", o.name, err)
	} else {
		fmt.Printf("db exited, name: %s\n", o.name)
	}
}
