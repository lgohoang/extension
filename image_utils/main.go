package image_utils

import (
	"bytes"
	"extension/utilities"
	"image"
	"image/draw"
	"image/gif"
	"os"

	"github.com/andybons/gogif"
	"github.com/fogleman/gg"
)

type Point struct {
	X, Y float64
}

type Text struct {
	FontFamilyPath string
	FontSize       float64
	FontColor      string
	Text           string
	Point          Point
}

type Image struct {
}

type GIF struct {
	*gif.GIF
}

func GIFFromFile(filename string) (*GIF, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	imageGif, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}

	return &GIF{
		imageGif,
	}, nil
}

func (i *GIF) AddText(text []Text) (*bytes.Buffer, error) {
	imgs, err := i.SplitAnimatedGIF()
	if err != nil {
		return nil, err
	}

	// result := *i.GIF

	for j, img := range imgs {
		dc := gg.NewContextForRGBA(&img)

		for _, t := range text {
			color, err := utilities.ParseHexColorFast(t.FontColor)
			if err != nil {
				return nil, err
			}

			dc.SetColor(color)
			dc.LoadFontFace(t.FontFamilyPath, t.FontSize)
			_, h := dc.MeasureString(t.Text)
			dc.DrawString(t.Text, t.Point.X, t.Point.Y+h)
		}

		dc.Clip()

		palettedImage := image.NewPaletted(dc.Image().Bounds(), nil)
		quantizer := gogif.MedianCutQuantizer{NumColor: 256}
		quantizer.Quantize(palettedImage, dc.Image().Bounds(), dc.Image(), image.Point{})
		draw.Draw(palettedImage, palettedImage.Rect, dc.Image(), dc.Image().Bounds().Min, draw.Src)

		i.Image[j] = palettedImage
	}

	buf := new(bytes.Buffer)

	if err := gif.EncodeAll(buf, i.GIF); err != nil {
		return nil, err
	}

	return buf, nil
}

func (i *GIF) SplitAnimatedGIF() ([]image.RGBA, error) {
	imgWidth, imgHeight := i.GetGifDimensions()

	result := []image.RGBA{}
	var lastImage image.Image = i.Image[0]

	for _, srcImg := range i.Image {

		overpaintImage := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

		draw.Draw(overpaintImage, overpaintImage.Bounds(), lastImage, image.Point{}, draw.Src)
		draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.Point{}, draw.Over)

		lastImage = overpaintImage
		result = append(result, *overpaintImage)

	}

	return result, nil
}

func (i *GIF) GetGifDimensions() (x, y int) {
	var lowestX int
	var lowestY int
	var highestX int
	var highestY int

	for _, img := range i.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}

	return highestX - lowestX, highestY - lowestY
}

func MergeImage(bg image.Image, qr image.Image, p image.Point) *image.RGBA {

	// starting position of the second image (bottom left)
	// sp2 := image.Point{qr.Bounds().Dx(), 0}

	// new rectangle for the second image
	r2 := image.Rectangle{p, bg.Bounds().Size()}

	// rectangle for the big image
	r := image.Rectangle{image.Point{0, 0}, r2.Max}

	new := image.NewRGBA(r)

	draw.Draw(new, bg.Bounds(), bg, image.Point{0, 0}, draw.Src)
	draw.Draw(new, r2, qr, image.Point{0, 0}, draw.Over)

	return new
}
