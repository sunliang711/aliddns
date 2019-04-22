package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/sunliang711/aliddns/config"
	"github.com/sunliang711/aliddns/recordOperation"
	"log"
)

var (
	Build   string
	Version string
)

func main() {
	//TODO:配置文件中正则表达式需要两个反斜杠来转移,比较麻烦
	version := pflag.BoolP("version", "v", false, "show version")
	pflag.Parse()

	if *version {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("Build at: %s\n", Build)
		return
	}
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
