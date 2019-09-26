package main

import (
	"archive/zip"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: zip2png <zipfile>\nConverts a remove.bg ZIP to a PNG")
	}
	filename := os.Args[1]

	fmt.Println("Extracting...")
	rgb, alpha := readZIP(filename)

	fmt.Println("Compositing...")
	composited := composite(rgb, alpha)

	fmt.Println("Saving...")
	savePNG(composited, filename+".png")
}

func readZIP(filename string) (image.Image, image.Image) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	var rgb image.Image
	var alpha image.Image

	for _, f := range r.File {
		if f.Name == "color.jpg" {
			rc, err := f.Open()
			if err != nil {
				panic(err)
			}
			rgb, err = jpeg.Decode(rc)
			if err != nil {
				panic(err)
			}
			rc.Close()
		}
		if f.Name == "alpha.png" {
			rc, err := f.Open()
			if err != nil {
				panic(err)
			}
			alpha, err = png.Decode(rc)
			if err != nil {
				panic(err)
			}
			rc.Close()
		}
	}

	return rgb, alpha
}

func composite(rgb image.Image, alpha image.Image) *image.NRGBA {
	dimensions := rgb.Bounds().Max
	width := dimensions.X
	height := dimensions.Y

	composited := image.NewNRGBA(image.Rect(0, 0, width, height))

	colorModel := composited.ColorModel()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			rgbColor := (colorModel.Convert(rgb.At(x, y))).(color.NRGBA)
			alphaColor := (alpha.At(x, y)).(color.Gray)
			rgbColor.A = alphaColor.Y

			composited.SetNRGBA(x, y, rgbColor)
		}
	}

	return composited
}

func savePNG(image *image.NRGBA, filename string) {
	resultPng, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer resultPng.Close()

	png.Encode(resultPng, image)
}
