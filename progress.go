package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func progressBar(a fyne.App, w fyne.Window) *fyne.Container {
	enterArmyDate := a.Preferences().StringWithFallback("enter_army_date", "none")
	dischargeArmyDate := a.Preferences().StringWithFallback("discharge_army_date", "none")

	if enterArmyDate == "none" || dischargeArmyDate == "none" {
		dialog.NewError(fmt.Errorf("설정에서 입대일과 전역일을 설정하세요"), w).Show()

		return container.NewVBox()
	}

	enterArmyDateTime, _ := time.Parse("20060102", enterArmyDate)
	dischargeArmyDateTime, _ := time.Parse("20060102", dischargeArmyDate)

	totalArmyDateNanoSeconds := dischargeArmyDateTime.Sub(enterArmyDateTime).Nanoseconds()
	diff := float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalArmyDateNanoSeconds)

	diffData := binding.BindFloat(&diff)

	totalProgress := widget.NewProgressBarWithData(diffData)

	return container.NewVBox(totalProgress)
}
