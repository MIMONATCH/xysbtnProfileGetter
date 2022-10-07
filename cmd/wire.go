//go:build wireinject
// +build wireinject

package main

import (
	"github.com/MIMONATCH/xysbtnProfileGetter/internal/biz"
	"github.com/MIMONATCH/xysbtnProfileGetter/internal/config"
	"github.com/MIMONATCH/xysbtnProfileGetter/internal/data"
	"github.com/google/wire"
)

func InitApp(configFile string) (*biz.App, error) {
	panic(wire.Build(
		biz.NewApp,
		config.NewConfig,
		data.ProviderSet,
		biz.ProviderSet,
	))
}
