package driver

import (
	_ "github.com/mattn/go-oci8"
	"log"
	"ri/utils"
	"strings"
	"time"
	"xorm.io/xorm"
)

var Engine *xorm.Engine
var err error

func init() {
	config, err := utils.ParseConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	var db = config.Database
	var builder strings.Builder
	builder.WriteString(db.User)
	builder.WriteString("/")
	builder.WriteString(db.Password)
	builder.WriteString("@")
	builder.WriteString(db.Host)
	builder.WriteString(":")
	builder.WriteString(db.Port)
	builder.WriteString("/")
	builder.WriteString(db.ServiceName)
	dsn := builder.String()
	log.Printf("DB init, driver: %s, datasource: %s", db.Driver, dsn)

	Engine, err = xorm.NewEngine(db.Driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = Engine.Ping(); err != nil {
		log.Println("DB connect failed")
		log.Fatal(err)
	}
	log.Println("DB connect success")
	Engine.SetMaxOpenConns(10)
	Engine.SetConnMaxLifetime(time.Hour)
	Engine.SetMaxIdleConns(5)
	Engine.ShowSQL(true)
}
