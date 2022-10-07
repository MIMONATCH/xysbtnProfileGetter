package config

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ProfileConfig struct {
	DownloadTime   int            `json:"downloadTime"`
	ProfileInfoAPI profileInfoAPI `json:"profileAPI"`
}

type profileInfoAPI struct {
	Url     string        `json:"url"`
	Timeout time.Duration `json:"timeout"`
}

// 解析config配置
func NewConfig(configFile string) (*ProfileConfig, error) {
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "read config failed")
	}
	config := &ProfileConfig{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "unmarshal failed")
	}
	return config, nil
}
