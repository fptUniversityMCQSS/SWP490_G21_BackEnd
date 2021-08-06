package utility

import (
	"SWP490_G21_Backend/model"
	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Config struct {
	DbServer     string
	DBDriver     string
	DbPort       string
	DbUser       string
	DbPassword   string
	Database     string
	Server       string
	AIServer     string
	PortBackend  string
	PortFrontend string
}

var ConfigData = ReadConfig()
var DB orm.Ormer

func init() {
	stringConfig := ConfigData.DbUser + ":" +
		ConfigData.DbPassword + "@tcp(" +
		ConfigData.DbServer + ":" +
		ConfigData.DbPort + ")/" +
		ConfigData.Database + "?charset=utf8"
	orm.RegisterModel(
		new(model.Knowledge),
		new(model.Option),
		new(model.Question),
		new(model.User),
		new(model.ExamTest),
	)
	err := orm.RegisterDriver(ConfigData.DBDriver, orm.DRMySQL)
	if err != nil {
		return
	}

	err = orm.RegisterDataBase("default", ConfigData.DBDriver, stringConfig)
	if err != nil {
		log.Printf("false %v", err)
	}

	DB = orm.NewOrm()
}

func ReadConfig() Config {
	var configFile = "properties.ini"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("property file is missing: ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
