package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Driver   string
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}
func initConfig() *AppConfig {
	var defaultConfig AppConfig
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Cannot read configuration")
		return nil
	}
	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}
	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("DBname")
	defaultConfig.Username = os.Getenv("DBusername")
	defaultConfig.Password = os.Getenv("DBpassword")
	defaultConfig.Address = os.Getenv("DBhost")
	cnv, err = strconv.Atoi(os.Getenv("DBport"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.Port = cnv

	return &defaultConfig
}
