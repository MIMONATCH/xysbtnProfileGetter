package main

import (
	"context"
	"flag"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "conf", "config/config.yaml", "path to your profileGetter config eg. /path/to/xiaoxi/config.yaml or ../config.yaml")
}

func main() {
	flag.Parse()
	app, err := InitApp(configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
