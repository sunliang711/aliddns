package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
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
		logrus.Fatalf("NewConfig error: %s", err)
	}

	logrus.SetLevel(lvl(cfg.Loglevel))
	operator, err := recordOperation.NewOperator(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	operator.AutomaticUpdate()
}

func lvl(level string) logrus.Level {
	switch level {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
