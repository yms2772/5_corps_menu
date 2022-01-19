package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

////go:embed src/NanumGothic.ttf
//var fontNanumGothic []byte

//go:embed src/NanumGothicBold.ttf
var fontNanumGothicBold []byte

var (
	resourceNanumGothic = &fyne.StaticResource{
		StaticName:    "NanumGothicBold.ttf",
		StaticContent: fontNanumGothicBold,
	}
	resourceNanumGothicBold = &fyne.StaticResource{
		StaticName:    "NanumGothicBold.ttf",
		StaticContent: fontNanumGothicBold,
	}
)
