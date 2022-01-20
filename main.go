package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	mainApp := app.NewWithID("mokky.mnd.menu")
	mainApp.Settings().SetTheme(&myTheme{})

	mainWindow := mainApp.NewWindow("5군단 식단")

	enterArmyDate := mainApp.Preferences().StringWithFallback("enter_army_date", "none")
	dischargeArmyDate := mainApp.Preferences().StringWithFallback("discharge_army_date", "none")
	vacationDay := mainApp.Preferences().IntWithFallback("vacation", 0)

	if enterArmyDate == "none" || dischargeArmyDate == "none" {
		dialog.ShowError(fmt.Errorf("설정에서 입대일과 전역일을 설정하세요"), mainWindow)
	}

	enterArmyDateTime, _ := time.Parse("20060102", enterArmyDate)
	dischargeArmyDateTime, _ := time.Parse("20060102", dischargeArmyDate)

	totalArmyDateNanoSeconds := dischargeArmyDateTime.Sub(enterArmyDateTime).Nanoseconds()
	totalWOVArmyDateNanoSeconds := dischargeArmyDateTime.AddDate(0, 0, -vacationDay).Sub(enterArmyDateTime).Nanoseconds()

	diff := float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalArmyDateNanoSeconds)
	diffWOV := float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalWOVArmyDateNanoSeconds)

	diffData := binding.BindFloat(&diff)
	diffWOVData := binding.BindFloat(&diffWOV)

	go func() {
		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			if err := diffData.Set(float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalArmyDateNanoSeconds)); err != nil {
				fmt.Println(err)
			}

			if err := diffWOVData.Set(float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalWOVArmyDateNanoSeconds)); err != nil {
				fmt.Println(err)
			}
		}
	}()

	totalProgress := widget.NewProgressBarWithData(diffData)
	totalProgress.TextFormatter = func() string {
		return fmt.Sprintf("%f%% (D-%d)", totalProgress.Value*100, int(dischargeArmyDateTime.Sub(enterArmyDateTime).Hours()/24)-int(time.Now().Sub(enterArmyDateTime).Hours()/24))
	}

	totalWOV := widget.NewProgressBarWithData(diffWOVData)
	totalWOV.TextFormatter = func() string {
		return fmt.Sprintf("%f%% (D-%d)", totalWOV.Value*100, int(dischargeArmyDateTime.AddDate(0, 0, -vacationDay).Sub(enterArmyDateTime).Hours()/24)-int(time.Now().Sub(enterArmyDateTime).Hours()/24))
	}

	remainProgress := container.NewVBox(totalProgress, totalWOV)
	tabItem := container.NewAppTabs(menuItem(mainWindow), settingsItem(mainApp, mainWindow))

	mainWindow.SetContent(container.NewBorder(remainProgress, nil, nil, nil, tabItem))
	mainWindow.ShowAndRun()
}
