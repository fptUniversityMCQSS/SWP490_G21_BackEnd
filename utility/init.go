package utility

import (
	"SWP490_G21_Backend/model/unity"
	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/natefinch/lumberjack"
	"log"
	"os"
)

type Config struct {
	DbServer         string
	DBDriver         string
	DbPort           string
	DbUser           string
	DbPassword       string
	Database         string
	AIServer         string
	PortHttps        string
	PortHttp         string
	StaticFolder     string
	HttpsKey         string
	HttpsCertificate string
}

var ConfigData = ReadConfig()
var DB orm.Ormer
var FileLog *log.Logger

func init() {
	stringConfig := ConfigData.DbUser + ":" +
		ConfigData.DbPassword + "@tcp(" +
		ConfigData.DbServer + ":" +
		ConfigData.DbPort + ")/" +
		ConfigData.Database + "?charset=utf8"
	orm.RegisterModel(
		new(unity.Knowledge),
		new(unity.Option),
		new(unity.Question),
		new(unity.User),
		new(unity.ExamTest),
	)
	err := orm.RegisterDriver(ConfigData.DBDriver, orm.DRMySQL)
	if err != nil {
		log.Fatal("Fail to register driver \"" + ConfigData.DBDriver + "\"")
	}

	err = orm.RegisterDataBase("default", ConfigData.DBDriver, stringConfig)
	if err != nil {
		log.Fatal("missing \"" + ConfigData.HttpsKey + "\"")
	}

	DB = orm.NewOrm()

	err = os.MkdirAll("examtest", os.ModePerm)
	if err != nil {
		log.Print(err)
	}

	err = os.MkdirAll("knowledge", os.ModePerm)
	if err != nil {
		log.Print(err)
	}

	err = os.MkdirAll("log", os.ModePerm)
	if err != nil {
		log.Print(err)
	}

	_, err = os.Stat(ConfigData.HttpsKey)
	if err != nil {
		log.Fatal("missing \"" + ConfigData.HttpsKey + "\"")
	}

	_, err = os.Stat(ConfigData.HttpsCertificate)
	if err != nil {
		log.Fatal("missing \"" + ConfigData.HttpsCertificate + "\"")
	}

	_, err = os.Stat(ConfigData.StaticFolder)
	if err != nil {
		log.Fatal("missing \"" + ConfigData.StaticFolder + "\" folder for static html files")
	}

	_, err = os.Stat("template")
	if err != nil {
		log.Fatal("missing \"template\" folder")
	}

	logFile, err := os.OpenFile("log/file.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Can not open or create log/file.log")
	}
	FileLog = log.New(logFile, "", log.Ldate|log.Ltime)
	FileLog.SetOutput(&lumberjack.Logger{
		Filename:   "log/file.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})
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
