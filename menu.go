package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func menuItem(w fyne.Window) *container.TabItem {
	data, err := LoadMenu()
	if err != nil {
		fmt.Println(err)
	}

	menu := GenerateMenuCard(data, time.Now().Format("2006-01-02"))

	breakfastCard := &widget.Card{
		Title:   "아침",
		Content: widget.NewLabel(strings.Join(menu[time.Now().Format("2006-01-02")]["아침"], "\n")),
	}
	lunchCard := &widget.Card{
		Title:   "점심",
		Content: widget.NewLabel(strings.Join(menu[time.Now().Format("2006-01-02")]["점심"], "\n")),
	}
	dinnerCard := &widget.Card{
		Title:   "저녁",
		Content: widget.NewLabel(strings.Join(menu[time.Now().Format("2006-01-02")]["저녁"], "\n")),
	}

	var nowDateBtn *widget.Button

	nowDateBtn = widget.NewButton(time.Now().Format("2006-01-02"), func() {
		nowDateBtn.SetText(time.Now().Format("2006-01-02"))

		menu = GenerateMenuCard(data, time.Now().Format("2006-01-02"))

		breakfastCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["아침"], "\n")))
		lunchCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["점심"], "\n")))
		dinnerCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["저녁"], "\n")))
	})

	beforeDateBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		nowDateBtnText, _ := time.Parse("2006-01-02", nowDateBtn.Text)
		dateMove := nowDateBtnText.AddDate(0, 0, -1)

		nowDateBtn.SetText(dateMove.Format("2006-01-02"))

		menu = GenerateMenuCard(data, nowDateBtn.Text)

		breakfastCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["아침"], "\n")))
		lunchCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["점심"], "\n")))
		dinnerCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["저녁"], "\n")))
	})

	afterDateBtn := widget.NewButtonWithIcon("", theme.NavigateNextIcon(), func() {
		nowDateBtnText, _ := time.Parse("2006-01-02", nowDateBtn.Text)
		dateMove := nowDateBtnText.AddDate(0, 0, 1)

		nowDateBtn.SetText(dateMove.Format("2006-01-02"))

		menu = GenerateMenuCard(data, nowDateBtn.Text)

		breakfastCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["아침"], "\n")))
		lunchCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["점심"], "\n")))
		dinnerCard.SetContent(widget.NewLabel(strings.Join(menu[nowDateBtn.Text]["저녁"], "\n")))
	})

	refreshBtn := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		updateProgress := widget.NewProgressBarInfinite()
		updateDialog := dialog.NewCustom("식단표 업데이트 중...", "취소", updateProgress, w)

		updateDialog.Show()

		data, err = LoadMenu()
		if err != nil {
			fmt.Println(err)
		}

		updateDialog.Hide()
	})

	dateButton := container.NewBorder(nil, nil, beforeDateBtn, afterDateBtn, nowDateBtn)

	return container.NewTabItemWithIcon("식단표", theme.SearchIcon(),
		container.NewVBox(dateButton, breakfastCard, lunchCard, dinnerCard, layout.NewSpacer(), refreshBtn),
	)
}
