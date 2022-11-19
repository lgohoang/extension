# extension
## GIF
### Add text

```
    gif, err := image.GIFFromFile("image.gif")
	if err != nil {
		panic(err)
	}

	img, err := gif.AddText([]image.Text{
		{
			Text:           "https://example.vn",
			FontFamilyPath: "Lato-Regular.ttf",
			FontSize:       20,
			FontColor:      "#de32ab",
			Point: image.Point{
				X: 0,
				Y: 0,
			},
		},
		{
			Text:           "First, Last Name",
			FontFamilyPath: "Lato-Regular.ttf",
			FontSize:       40,
			FontColor:      "#fcba03",
			Point: image.Point{
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
```