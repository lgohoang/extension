package main

import (
	"extension/image_utils"
	"os"
)

// for tester
func main() {
	gif, err := image_utils.GIFFromFile("/Volumes/Storage/Workspace/Tools/QrCodeInvite/bg2.gif")
	if err != nil {
		panic(err)
	}

	img, err := gif.AddText([]image_utils.Text{
		{
			Text:           "https://example.vn",
			FontFamilyPath: "/Volumes/Storage/Workspace/Tools/QrCodeInvite/fonts/Lato-Regular.ttf",
			FontSize:       20,
			FontColor:      "#de32ab",
			Point: image_utils.Point{
				X: 0,
				Y: 0,
			},
		},
		{
			Text:           "Nguyễn Văn Hoàng",
			FontFamilyPath: "/Volumes/Storage/Workspace/Tools/QrCodeInvite/fonts/Lato-Regular.ttf",
			FontSize:       40,
			FontColor:      "#fcba03",
			Point: image_utils.Point{
				X: 0,
				Y: 100,
			},
		},
	})

	if err != nil {
		panic(err)
	}

	out, err := os.Create("image.gif")
	if err != nil {
		panic(err)
	}

	_, err = out.Write(img.Bytes())
	if err != nil {
		panic(err)
	}

	img, err = gif.AddText([]image_utils.Text{
		{
			Text:           "Test Text 2",
			FontFamilyPath: "/Volumes/Storage/Workspace/Tools/QrCodeInvite/fonts/Lato-Regular.ttf",
			FontSize:       40,
			FontColor:      "#fcba03",
			Point: image_utils.Point{
				X: 0,
				Y: 100,
			},
		},
	})

	if err != nil {
		panic(err)
	}

	out, err = os.Create("image2.gif")
	if err != nil {
		panic(err)
	}

	_, err = out.Write(img.Bytes())
	if err != nil {
		panic(err)
	}
}
