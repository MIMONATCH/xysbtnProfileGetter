package data

import (
	"os"
	"path"

	pkgdata "github.com/MIMONATCH/xysbtnProfileGetter/internal/pkg/data"
	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) loadData() (*pkgdata.Supports, error) {
	files, err := os.ReadDir("../assets")
	if err != nil {
		return nil, errors.Wrap(err, "scan assets dir is failed")
	}
	// 单个配置文件 不支持多个
	if len(files) != 1 {
		return nil, errors.New("must have one supports config, not more")
	}
	siteInfo, err := os.ReadFile(path.Join("../assets", files[0].Name()))
	if err != nil {
		return nil, errors.Wrapf(err, "read [%s] failed", files[0].Name())
	}
	supports := pkgdata.Supports{}
	if err := yaml.Unmarshal(siteInfo, &supports); err != nil {
		return nil, errors.Wrapf(err, "[%s] unmarshal failed", files[0].Name())
	}
	return &supports, nil
}

func (r *Repo) ListBids() (*pkgdata.Supports, error) {
	return r.loadData()
}
