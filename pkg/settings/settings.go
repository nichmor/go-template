package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type Server struct {
	Compress bool
	RunMode  string
	HttpPort int
}

var ServerSetting = &Server{}

type Mongo struct {
	Address        string
	Database       string
	KeysCollection string
}

var MongoSetting = &Mongo{}

type Redis struct {
	Address  string
	Database int
}

var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("mongo", MongoSetting)
	mapTo("redis", RedisSetting)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}

}
