package main

import (
	"fmt"
	"log"

	"github.com/sunliang711/aliddns/config"
	"github.com/sunliang711/aliddns/recordOperation"
)

var (
	sha1ver   string
	buildTime string
)

func main() {
	fmt.Printf("Version: %v\n", sha1ver)
	fmt.Printf("Build time: %v\n", buildTime)
	cfg, err := config.NewConfig("config.toml")
	if err != nil {
		log.Fatalf("NewConfig error: %s", err)
	}

	operator, err := recordOperation.NewOperator(cfg)
	if err != nil {
		log.Fatal(err)
	}

	operator.AutomaticUpdate()
}
