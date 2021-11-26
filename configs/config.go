package configs

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type Config struct {
	HTTP HTTPConf
	DB   DBConf
}

type DBConf struct {
	User     string `envconfig:"USER" default:"root"`
	Password string `envconfig:"PASSWORD" default:"root"`
	URL      string `envconfig:"URL" default:"localhost"`
	DBName   string `envconfig:"DBNAME" default:"library"`
	DSN      string `envconfig:"DSN" default:"root:root@tcp(localhost:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"`
}

type HTTPConf struct {
	TLSEnable      bool   `envconfig:"TLS_ENABLE"`
	Enable         bool   `envconfig:"ENABLE" default:"true"`
	Port           string `envconfig:"PORT" default:":8000"`
	TLSKeyAddress  string `envconfig:"TLS_KEY_ADDRESS"`
	TLSCertAddress string `envconfig:"TLS_CERT_ADDRESS"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("%+v\n", config)
	})
	return &config
}
