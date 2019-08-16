package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	//Cfg this is Cfg
	Cfg *ini.File
	//RunMode this is RunMode
	RunMode string
	//HTTPPort this is HTTPPort
	HTTPPort int
	//ReadTimeout this is ReadTimeout
	ReadTimeout time.Duration
	//WriteTimeout this is WriteTimeout
	WriteTimeout time.Duration
	//PageSize this is PageSize
	PageSize int
	// JwtSecret this is JwtSecret
	JwtSecret string
	// TYPE this is TYPE
	TYPE string
	// USER this is USER
	USER string
	// PASSWORD this is PASSWORD
	PASSWORD string
	// HOST this is HOST
	HOST string
	// DBNAME this is DBNAME
	DBNAME string
	// APIKEY this is APIKEY
	APIKEY string
)

func init() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Println("Load conf/app.ini failed", err)
	}
	loadServer(Cfg)
	loadApp(Cfg)
	loadDatabase(Cfg)
}

func loadServer(Cfg *ini.File) {
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	HTTPPort = server.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp(Cfg *ini.File) {
	server, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = server.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	APIKEY = server.Key("API_KEY").String()
	PageSize = server.Key("PAGE_SIZE").MustInt(10)
}
func loadDatabase(Cfg *ini.File) {
	database, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	TYPE = database.Key("TYPE").String()
	USER = database.Key("USER").String()
	PASSWORD = database.Key("PASSWORD").String()
	HOST = database.Key("HOST").String()
	DBNAME = database.Key("DBNAME").String()
}
