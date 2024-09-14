package biz

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/pkg/errors"

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
	profileUrl, err := SignAndGenerateURL(fmt.Sprint(d.config.ProfileInfoAPI.Url, support.Uid))
	if err != nil {
		return err
	}
	resp, err := d.profile.Check(profileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		return nil
	}

	outFile, err := os.Create(fmt.Sprint(support.Uid))
	defer outFile.Close()
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response body error")
	}
	content, err := os.ReadFile(fmt.Sprint(support.Uid))
	if err != nil {
		return errors.Wrap(err, "read file error")
	}

	body, err := io.ReadAll(resp.Body)
	fmt.Print(string(body))

	compileRegex := regexp.MustCompile("\"face\":\"(.*?)\",")
	face := compileRegex.FindStringSubmatch(string(content))

	fmt.Println(face)
	if len(face) == 0 {
		return nil
	}

	// 检查 profile url是否可达
	profileResp, err := d.profile.Check(face[1])
	if err != nil {
		return err
	}
	defer profileResp.Body.Close()

	fmt.Print(resp.Body)

	if err := d.compress.ProfileCompress(profileResp.Body, support.Uid); err != nil {
		return err
	}
	return nil
}
