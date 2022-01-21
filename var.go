package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

//go:embed src/NanumGothicBold.ttf
var fontNanumGothicBold []byte

var (
	resourceNanumGothicBold = &fyne.StaticResource{
		StaticName:    "NanumGothicBold.ttf",
		StaticContent: fontNanumGothicBold,
	}
)

var (
	ticker *time.Ticker

	totalWOVProgress *widget.ProgressBar

	enterArmyDate     string
	dischargeArmyDate string
	vacationDay       int

	enterArmyDateTime     time.Time
	dischargeArmyDateTime time.Time

	totalArmyDateNanoSeconds    int64
	totalWOVArmyDateNanoSeconds int64

	diff    float64
	diffWOV float64
)
