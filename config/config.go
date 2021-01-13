package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"time"
)

var Config = struct {
	DebugMode    bool          `default:"true"`
	TimeInterval time.Duration `default:"300" env:"TIME_INTERVAL"`

	DDNS struct {
		LoginToken   string `env:"LOGIN_TOKEN"`
		DomainID     string `env:"DOMAIN_ID"`
		RecordID     string `env:"RECORD_ID"`
		RecordLineID string `env:"RECORD_LINE_ID"`
		SubDomain    string `env:"SUB_DOMAIN"`
	}
}{}

func init() {
	if err := configor.Load(&Config); err != nil {
		panic(err)
	}
	fmt.Printf("config: %#v", Config)
}
