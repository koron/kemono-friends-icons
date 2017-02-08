package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

var dstRect = image.Rect(0, 0, 128, 128)

func load(name string) (image.Image, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func save(name string, img image.Image) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	err = png.Encode(f, img)
	f.Close()
	if err != nil {
		os.Remove(name)
		return nil
	}
	return nil
}

func trim(src image.Image, x, y, w, h int) image.Image {
	r := image.Rect(0, 0, w, h)
	dst := image.NewRGBA(r)
	draw.Draw(dst, r, src, image.Pt(x, y), draw.Src)
	return dst
}

func conv(dst, src string) error {
	si, err := load(src)
	if err != nil {
		return err
	}
	mi := trim(si, 43, 3, 274, 274)
	di := resize.Resize(128, 128, mi, resize.Bilinear)
	return save(dst, di)
}

func convAll(dstDir, srcDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dst := filepath.Join(dstDir, info.Name())
		if !strings.HasSuffix(path, ".png") {
			return nil
		}
		fmt.Printf("%s <- %s\n", dst, path)
		return conv(dst, path)
	})
}

func main() {
	err := convAll("icon2", "src2")
	if err != nil {
		log.Fatal(err)
	}
}
