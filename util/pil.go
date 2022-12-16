/***************************
@File        : pil.go
@Time        : 2022/03/04 16:08:07
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 图片处理
****************************/

package util

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

// Draw 将图片B与图片A组合生成新图片
func Draw(A, B string, path string) error {
	img, err := openImage(A)
	if err != nil {
		return err
	}

	watermark, err := openImage(B)
	if err != nil {
		return err
	}
	offset := image.Pt(0, 0)
	b := img.Bounds()
	m := image.NewRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	imgw, _ := os.Create(path)
	jpeg.Encode(imgw, m, &jpeg.Options{Quality: 100})
	defer imgw.Close()
	return nil
}

func openImage(pic string) (image.Image, error) {
	var image image.Image
	imgb, err := os.Open(pic)
	defer imgb.Close()

	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(imgb)
	image, err = jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		image, err = png.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, err
		}
	}
	return image, nil
}
