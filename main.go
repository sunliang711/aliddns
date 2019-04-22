package main

import (
	"github.com/sunliang711/aliddns/config"
	"github.com/sunliang711/aliddns/recordOperation"
	"log"
)

func main() {
	cfg, err := config.NewConfig("config.toml")
	if err != nil {
		log.Fatalf("NewConfig error: %s",err)
	}

	operator, err := recordOperation.NewOperator(cfg)
	if err != nil {
		log.Fatal(err)
	}

    operator.AutomaticUpdate()
}

