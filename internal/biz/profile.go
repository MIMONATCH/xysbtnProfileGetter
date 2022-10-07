package biz

import (
	"net/http"

	"github.com/MIMONATCH/xysbtnProfileGetter/internal/config"
	"github.com/pkg/errors"
)

type Profile struct {
	conf   *config.ProfileConfig
	client *http.Client
}

func NewProfile(conf *config.ProfileConfig) *Profile {
	client := http.Client{
		Timeout: conf.ProfileInfoAPI.Timeout,
	}
	return &Profile{
		conf:   conf,
		client: &client,
	}
}

// 接口探测是否可达
func (p *Profile) Check(url string) (*http.Response, error) {
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "url[%s] is redirect url or client get error", url)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Wrapf(errors.New("url is not avaliable"), "url:[%s]", url)
	}
	return resp, nil
}
