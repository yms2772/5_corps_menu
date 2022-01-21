package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func settingsItem(a fyne.App, w fyne.Window) *container.TabItem {
	//TODO: warn을 dialog로 변경하기
	enterArmyDateWarn := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
	enterArmyDateWarn.Hide()

	enterArmyDateEntry := widget.NewEntry()
	enterArmyDateEntry.PlaceHolder = "예시: 20210315"
	enterArmyDateEntry.Validator = func(s string) error {
		r, _ := regexp.Compile(`\b(20\d{2})(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])\b`)

		if !r.MatchString(s) {
			return fmt.Errorf("날짜 예: 20210315")
		}

		return nil
	}
	enterArmyDateEntry.SetOnValidationChanged(func(err error) {
		if err != nil {
			enterArmyDateWarn.SetText(fmt.Sprintf("* %s", err.Error()))
			enterArmyDateWarn.Show()
		} else {
			enterArmyDateWarn.Hide()
		}
	})

	enterArmyDateBox := container.NewHBox(enterArmyDateEntry, enterArmyDateWarn)

	dischargeArmyDateWarn := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
	dischargeArmyDateWarn.Hide()

	dischargeArmyDateEntry := widget.NewEntry()
	dischargeArmyDateEntry.PlaceHolder = "예시: 20210315"
	dischargeArmyDateEntry.Validator = func(s string) error {
		r, _ := regexp.Compile(`\b(20\d{2})(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])\b`)

		if !r.MatchString(s) {
			return fmt.Errorf("날짜 예: 20210315")
		}

		return nil
	}
	dischargeArmyDateEntry.SetOnValidationChanged(func(err error) {
		if err != nil {
			dischargeArmyDateWarn.SetText(fmt.Sprintf("* %s", err.Error()))
			dischargeArmyDateWarn.Show()
		} else {
			dischargeArmyDateWarn.Hide()
		}
	})

	dischargeArmyDateBox := container.NewHBox(dischargeArmyDateEntry, dischargeArmyDateWarn)

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
			remainedVacation.SetText("")

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

	earlyDischargeCheck := widget.NewCheck("조기전역", func(b bool) {
		if b {
			a.Preferences().SetBool("early_discharge", true)

			totalWOVProgress.Show()
		} else {
			a.Preferences().SetBool("early_discharge", false)

			totalWOVProgress.Hide()
		}
	})
	earlyDischargeCheck.Checked = a.Preferences().BoolWithFallback("early_discharge", true)

	remainedVacationBox := container.NewHBox(remainedVacation, widget.NewLabel("일"), earlyDischargeCheck, remainedVacationWarn)

	armySelect := widget.NewSelect([]string{"5군단"}, func(s string) {

	})
	armySelect.Selected = "5군단"

	progressTickerUpdate := widget.NewSelect([]string{"10", "50", "100", "500", "1000"}, func(s string) {
		t, _ := strconv.ParseFloat(s, 64)

		ticker.Reset(time.Duration(t) * time.Millisecond)
	})
	progressTickerUpdate.Selected = "100"
	tickerUpdateBox := container.NewHBox(progressTickerUpdate, widget.NewLabel("ms"))

	progressDecimalWarn := widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
	progressDecimalWarn.Hide()

	progressDecimal := widget.NewEntry()
	progressDecimal.SetText(fmt.Sprintf("%d", a.Preferences().IntWithFallback("progress_decimal_point", 6)))
	progressDecimal.Validator = func(s string) error {
		if len(s) == 0 {
			return nil
		}

		if point, err := strconv.Atoi(s); err != nil {
			progressDecimal.SetText("")

			return fmt.Errorf("숫자만 입력할 수 있습니다")
		} else if point < 0 || point > 10 {
			progressDecimal.SetText("")

			return fmt.Errorf("소수점은 0부터 10까지만 가능합니다")
		} else {
			a.Preferences().SetInt("progress_decimal_point", point)
		}

		return nil
	}
	progressDecimal.SetOnValidationChanged(func(err error) {
		if err != nil {
			progressDecimalWarn.SetText(fmt.Sprintf("* %s", err.Error()))
			progressDecimalWarn.Show()
		} else {
			progressDecimalWarn.Hide()
		}
	})

	progressDecimalBox := container.NewHBox(progressDecimal, widget.NewLabel("자리"), progressDecimalWarn)

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

			dialog.ShowInformation("설정", "저장되었습니다", w)

			UpdateDate(a)
		},
	}

	form.Append("입대일", enterArmyDateBox)
	form.Append("전역일", dischargeArmyDateBox)
	form.Append("남은 휴가일", remainedVacationBox)
	form.Append("부대", armySelect)

	form2 := &widget.Form{}
	form2.Append("진행률 업데이트 주기", tickerUpdateBox)
	form2.Append("진행률 소수점", progressDecimalBox)

	armySettings := widget.NewCard("복무 설정", "", form)
	generalSettings := widget.NewCard("일반 설정", "", form2)

	return container.NewTabItemWithIcon("설정", theme.SettingsIcon(),
		container.NewVBox(armySettings, generalSettings),
	)
}
