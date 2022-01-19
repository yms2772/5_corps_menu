package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"regexp"
	"strconv"
)

func settingsItem(a fyne.App, w fyne.Window) *container.TabItem {
	enterArmyDateEntry := widget.NewEntry()
	enterArmyDateEntry.PlaceHolder = "예시: 20210315"

	dischargeArmyDateEntry := widget.NewEntry()
	dischargeArmyDateEntry.PlaceHolder = "예시: 20210315"

	if data := a.Preferences().StringWithFallback("enter_army_date", "none"); data != "none" {
		enterArmyDateEntry.SetText(data)
	}

	if data := a.Preferences().StringWithFallback("discharge_army_date", "none"); data != "none" {
		dischargeArmyDateEntry.SetText(data)
	}

	armySelect := widget.NewSelect([]string{"5군단"}, func(s string) {

	})
	armySelect.Selected = "5군단"

	form := &widget.Form{
		SubmitText: "저장",
		OnSubmit: func() {
			if _, err := strconv.Atoi(enterArmyDateEntry.Text); err != nil {
				dialog.ShowError(fmt.Errorf("입대일이 숫자가 아닙니다"), w)

				return
			}

			if _, err := strconv.Atoi(dischargeArmyDateEntry.Text); err != nil {
				dialog.ShowError(fmt.Errorf("전역일이 숫자가 아닙니다"), w)

				return
			}

			r, _ := regexp.Compile(`\b\d{8}\b`)

			if !r.MatchString(enterArmyDateEntry.Text) {
				dialog.ShowError(fmt.Errorf("입대일 입력 예시: 20210315"), w)

				return
			} else if !r.MatchString(dischargeArmyDateEntry.Text) {
				dialog.ShowError(fmt.Errorf("전역일 입력 예시: 20210315"), w)

				return
			}

			a.Preferences().SetString("enter_army_date", enterArmyDateEntry.Text)
			a.Preferences().SetString("discharge_army_date", dischargeArmyDateEntry.Text)

			dialog.ShowInformation("설정", "저장완료", w)
		},
	}
	form.Append("입대일", enterArmyDateEntry)
	form.Append("전역일", dischargeArmyDateEntry)
	form.Append("부대", armySelect)

	return container.NewTabItemWithIcon("설정", theme.SettingsIcon(),
		container.NewVBox(form),
	)
}
