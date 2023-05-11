package biz

import (
	"context"
	"sync"

	"github.com/MIMONATCH/xysbtnProfileGetter/internal/data"
	pkgdata "github.com/MIMONATCH/xysbtnProfileGetter/internal/pkg/data"
)

type App struct {
	download *Download
	repo     *data.Repo
	wg       sync.WaitGroup
	failTask chan int
}

func NewApp(download *Download, repo *data.Repo) *App {
	return &App{
		download: download,
		repo:     repo,
	}
}

func (a *App) Run(ctx context.Context) error {
	data, err := a.repo.ListBids()
	if err != nil {
		return err
	}

	a.wg.Add(len(data.Supports))
	// 下载头像
	for index, one := range data.Supports {
		go func(one pkgdata.Support, index int) {
			defer a.wg.Done()
			if err := a.download.ProfileDownload(&one); err != nil {
				a.failTask <- index
			}
		}(one, index)
	}
	a.wg.Wait()

	// 下载失败的头像重试
	if len(a.failTask) == 0 {
		return nil
	}
	for task := range a.failTask {
		if err := a.download.ProfileDownload(&data.Supports[task]); err != nil {
			return err
		}
	}
	return nil
}
