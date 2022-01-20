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

	remainedVacationWarn := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
	remainedVacationWarn.Hide()

	remainedVacation := widget.NewEntry()
	remainedVacation.SetText(fmt.Sprintf("%d", a.Preferences().IntWithFallback("vacation", 0)))
	remainedVacation.Validator = func(s string) error {
		if len(s) == 0 {
			return nil
		}

		if _, err := strconv.Atoi(s); err != nil {
			return fmt.Errorf("숫자만 입력할 수 있습니다")
		}

		return nil
	}
	remainedVacation.SetOnValidationChanged(func(err error) {
		if err != nil {
			remainedVacationWarn.SetText(fmt.Sprintf("* %s", err.Error()))
			remainedVacationWarn.Show()
		} else {
			remainedVacationWarn.Hide()
		}
	})

	remainedVacationBox := container.NewHBox(remainedVacation, widget.NewLabel("일"), remainedVacationWarn)

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

			if data, err := strconv.Atoi(remainedVacation.Text); err != nil {
				dialog.ShowError(fmt.Errorf("남은 휴가일은 숫자만 입력할 수 있습니다"), w)

				return
			} else {
				a.Preferences().SetInt("vacation", data)
			}

			a.Preferences().SetString("enter_army_date", enterArmyDateEntry.Text)
			a.Preferences().SetString("discharge_army_date", dischargeArmyDateEntry.Text)

			reloadApp := dialog.NewConfirm("설정", "앱을 재시작해야 설정이 반영됩니다", func(b bool) {
				if b {
					a.Quit()
				}
			}, w)
			reloadApp.SetConfirmText("재시작")
			reloadApp.SetDismissText("취소")
			reloadApp.Show()
		},
	}

	form.Append("입대일", enterArmyDateEntry)
	form.Append("전역일", dischargeArmyDateEntry)
	form.Append("남은 휴가일", remainedVacationBox)
	form.Append("부대", armySelect)

	return container.NewTabItemWithIcon("설정", theme.SettingsIcon(),
		container.NewVBox(form),
	)
}
