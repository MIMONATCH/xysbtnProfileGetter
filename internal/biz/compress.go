package biz

import (
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"path"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	"github.com/pkg/errors"
)

type Compress struct {
}

func NewCompress() *Compress {
	return &Compress{}
}

func (c *Compress) ProfileCompress(r io.Reader, name string) error {
	// 创建临时的jpg文件
	if err := os.MkdirAll("../webp", 0750); err != nil {
		return errors.Wrap(err, "mkdir webp failed")
	}
	file, err := os.CreateTemp("../webp", "temp")
	if err != nil {
		return errors.Wrap(err, "temp jpg create failed")
	}
	if _, err := io.Copy(file, r); err != nil {
		return errors.Wrap(err, "copy response reader to jpg reader failed")
	}

	defer file.Close()
	defer os.Remove(file.Name())

	// 临时文件reader
	jpgReader, err := os.Open(path.Join("../webp", file.Name()))
	if err != nil {
		return errors.Wrap(err, "open temp jpg file failed")
	}
	defer jpgReader.Close()

	// 解码jpg、转换为webp
	img, err := jpeg.Decode(jpgReader)
	if err != nil {
		return errors.Wrap(err, "image decode failed")
	}
	w, err := os.Create(fmt.Sprint("../webp/", name, ".webp"))
	if err != nil {
		return errors.Wrap(err, "create webp file failed")
	}
	defer w.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		return errors.Wrap(err, "create webp option failed")
	}
	if err := webp.Encode(w, img, options); err != nil {
		return errors.Wrap(err, "webp encoder is failed")
	}
	return nil
}
