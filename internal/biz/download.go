package biz

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"

	"github.com/MIMONATCH/xysbtnProfileGetter/internal/config"
	"github.com/MIMONATCH/xysbtnProfileGetter/internal/pkg/data"
)

type Download struct {
	compress *Compress
	config   *config.ProfileConfig
	profile  *Profile
}

func NewDownload(config *config.ProfileConfig, compress *Compress, profile *Profile) *Download {
	return &Download{
		config:   config,
		compress: compress,
		profile:  profile,
	}
}

func (d *Download) ProfileDownload(support *data.Support) error {
	resp, err := d.profile.Check(fmt.Sprint(d.config.ProfileInfoAPI.Url, support.Uid))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 从返回的json中解析出头像url
	upData, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response body error")
	}

	face := gojsonq.New().FromString(string(upData)).Find("data.face")

	// 检查 profile url是否可达
	profileResp, err := d.profile.Check(face.(string))
	if err != nil {
		return err
	}
	defer profileResp.Body.Close()

	if err := d.compress.ProfileCompress(profileResp.Body, support.Uid); err != nil {
		return err
	}
	return nil
}
