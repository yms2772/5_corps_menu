package main

import (
	_ "embed"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

//go:embed src/NanumGothicBold.ttf
var fontNanumGothicBold []byte

//go:embed src/tab_food_icon.png
var iconTabFood []byte

var (
	resourceNanumGothicBold = &fyne.StaticResource{
		StaticName:    "NanumGothicBold.ttf",
		StaticContent: fontNanumGothicBold,
	}
	resourceIconTabFood = &fyne.StaticResource{
		StaticName:    "tab_food_icon.png",
		StaticContent: iconTabFood,
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
