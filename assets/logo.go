package assets

import (
	"embed"
	"fmt"
	"image"
)

import (
	_ "golang.org/x/image/bmp" // allow decoding/loading .bmp images
	_ "image/jpeg"             // allow decoding/loading .jpg images
)

//go:embed amadej-tauses-xWOTojs1eg4-unsplash.jpg
var logoFS embed.FS

func GetLogo() (image.Image, error) {
	path := "amadej-tauses-xWOTojs1eg4-unsplash.jpg"
	logo, err := logoFS.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed opening image file: %w", err)
	}
	defer logo.Close()
	imgData, _ /* format */, err := image.Decode(logo)
	if err != nil {
		return nil, fmt.Errorf("failed decoding image data: %w", err)
	}
	return imgData, nil
}
