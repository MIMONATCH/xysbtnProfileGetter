package biz

import (
	"fmt"
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
	// resp, err := p.client.Get(url)
	resq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request get error")
	}
	resq.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	resq.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36")
	resq.Header.Add("Referer", fmt.Sprint("https://space.bilibili.com"))
	resp, err := p.client.Do(resq)
	if err != nil {
		return nil, errors.Wrapf(err, "url[%s] is redirect url or client get error", url)
	}
	if resp.StatusCode != 200 {
		return nil, errors.Wrapf(errors.New("url is not avaliable"), "url:[%s]", url)
	}
	return resp, nil
}
