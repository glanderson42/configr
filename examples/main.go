package main

import (
	"fmt"

	"github.com/glanderson42/configr"
)

type Config struct {
	Host       string `configr:"{'default': 'localhost', 'required': true, 'env': 'HOST'}"`
	Port       int    `configr:"{'default': 8080, 'required': true, 'env': 'PORT'}"`
	Production bool   `configr:"{'default': false, 'required': false, 'env': 'PRODUCTION'}"`
}

func main() {
	var cfg Config
	err := configr.ParseConfig(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg.Host)
	fmt.Println(cfg.Port)
	fmt.Println(cfg.Production)
}
