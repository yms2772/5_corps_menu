package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
	"time"
)

func main() {
	mainApp := app.New()
	mainApp.Settings().SetTheme(&myTheme{})

	mainWindow := mainApp.NewWindow("5군단 식단")

	data, err := LoadMenu()
	if err != nil {
		fmt.Println(err)
	}

	breakfastCard := &widget.Card{}
	lunchCard := &widget.Card{}
	dinnerCard := &widget.Card{}

	menu := GenerateMenuCard(data, time.Now().Format("2006-01-02"))

	breakfastCard.SetContent(widget.NewLabel(strings.Join(menu[time.Now().Format("2006-01-02")]["아침"], "\n")))
	lunchCard.SetContent(widget.NewLabel(strings.Join(menu[time.Now().Format("2006-01-02")]["점심"], "\n")))
	dinnerCard.SetContent(widget.NewLabel(strings.Join(menu[time.Now().Format("2006-01-02")]["저녁"], "\n")))

	nowDateBtn := widget.NewButton(time.Now().Format("2006-01-02"), func() {
		updateProgress := widget.NewProgressBarInfinite()
		updateDialog := dialog.NewCustom("식단표 업데이트 중...", "취소", updateProgress, mainWindow)

		updateDialog.Show()

		data, err = LoadMenu()
		if err != nil {
			fmt.Println(err)
		}

		updateDialog.Hide()
	})

	beforeDateBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		log.Println("Before Btn Pressed")

		nowDateBtnText, _ := time.Parse("2006-01-02", nowDateBtn.Text)
		dateMove := nowDateBtnText.AddDate(0, 0, -1)

		nowDateBtn.SetText(dateMove.Format("2006-01-02"))

		menu = GenerateMenuCard(data, nowDateBtn.Text)

		breakfastCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["아침"], "\n")))
		lunchCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["점심"], "\n")))
		dinnerCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["저녁"], "\n")))
	})

	afterDateBtn := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		log.Println("After Btn Pressed")

		nowDateBtnText, _ := time.Parse("2006-01-02", nowDateBtn.Text)
		dateMove := nowDateBtnText.AddDate(0, 0, 1)

		nowDateBtn.SetText(dateMove.Format("2006-01-02"))

		menu = GenerateMenuCard(data, nowDateBtn.Text)

		breakfastCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["아침"], "\n")))
		lunchCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["점심"], "\n")))
		dinnerCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["저녁"], "\n")))
	})

	dateButton := container.NewBorder(nil, nil, beforeDateBtn, afterDateBtn, nowDateBtn)
	content := container.NewVBox(dateButton, breakfastCard, lunchCard, dinnerCard)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
