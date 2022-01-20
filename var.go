package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed src/NanumGothicBold.ttf
var fontNanumGothicBold []byte

var (
	resourceNanumGothicBold = &fyne.StaticResource{
		StaticName:    "NanumGothicBold.ttf",
		StaticContent: fontNanumGothicBold,
	}
)
