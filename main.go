package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	mainApp := app.NewWithID("mokky.mnd.menu")
	mainApp.Settings().SetTheme(&myTheme{})

	mainWindow := mainApp.NewWindow("5군단 식단")

	enterArmyDate = mainApp.Preferences().StringWithFallback("enter_army_date", "none")
	dischargeArmyDate = mainApp.Preferences().StringWithFallback("discharge_army_date", "none")
	vacationDay = mainApp.Preferences().IntWithFallback("vacation", 0)

	if enterArmyDate == "none" || dischargeArmyDate == "none" {
		dialog.ShowError(fmt.Errorf("설정에서 입대일과 전역일을 설정하세요"), mainWindow)
	}

	UpdateDate(mainApp)

	diffData := binding.BindFloat(&diff)
	diffWOVData := binding.BindFloat(&diffWOV)

	go func() {
		ticker = time.NewTicker(time.Duration(mainApp.Preferences().IntWithFallback("progress_update_ticker", 100)) * time.Millisecond)

		for {
			select {
			case <-ticker.C:
				if err := diffData.Set(float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalArmyDateNanoSeconds)); err != nil {
					fmt.Println(err)
				}

				if err := diffWOVData.Set(float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalWOVArmyDateNanoSeconds)); err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	totalProgress := widget.NewProgressBarWithData(diffData)
	totalProgress.TextFormatter = func() string {
		if enterArmyDate == "none" || dischargeArmyDate == "none" {
			return "데이터 없음"
		}

		return fmt.Sprintf("%s%% (D-%d)", FixedDecimal(totalProgress.Value*100, mainApp.Preferences().IntWithFallback("progress_decimal_point", 6)), int(dischargeArmyDateTime.Sub(enterArmyDateTime).Hours()/24)-int(time.Now().Sub(enterArmyDateTime).Hours()/24))
	}

	totalWOVProgress = widget.NewProgressBarWithData(diffWOVData)
	totalWOVProgress.Hidden = !mainApp.Preferences().BoolWithFallback("early_discharge", true)
	totalWOVProgress.TextFormatter = func() string {
		if enterArmyDate == "none" || dischargeArmyDate == "none" {
			return "데이터 없음"
		}

		return fmt.Sprintf("%s%% (D-%d)", FixedDecimal(totalWOVProgress.Value*100, mainApp.Preferences().IntWithFallback("progress_decimal_point", 6)), int(dischargeArmyDateTime.AddDate(0, 0, -vacationDay).Sub(enterArmyDateTime).Hours()/24)-int(time.Now().Sub(enterArmyDateTime).Hours()/24))
	}

	remainProgress := container.NewVBox(totalProgress, totalWOVProgress)
	tabItem := container.NewAppTabs(menuItem(mainWindow), settingsItem(mainApp, mainWindow))

	mainWindow.SetContent(container.NewBorder(remainProgress, nil, nil, nil, tabItem))
	mainWindow.ShowAndRun()
}
