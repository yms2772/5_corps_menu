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
	enterArmyDateEntry := widget.NewEntry()
	enterArmyDateEntry.PlaceHolder = "예시: 20210315"
	enterArmyDateEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return nil
		}

		r, _ := regexp.Compile(`\b(20\d{2})(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])\b`)

		if !r.MatchString(s) {
			return fmt.Errorf("날짜 형식이 올바르지 않습니다")
		}

		return nil
	}
	enterArmyDateEntry.SetOnValidationChanged(func(err error) {
		if err != nil {
			dialog.ShowError(err, w)
		}
	})

	dischargeArmyDateEntry := widget.NewEntry()
	dischargeArmyDateEntry.PlaceHolder = "예시: 20210315"
	dischargeArmyDateEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return nil
		}

		r, _ := regexp.Compile(`\b(20\d{2})(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])\b`)

		if !r.MatchString(s) {
			return fmt.Errorf("날짜 형식이 올바르지 않습니다")
		}

		return nil
	}
	dischargeArmyDateEntry.SetOnValidationChanged(func(err error) {
		if err != nil {
			dialog.ShowError(err, w)
		}
	})

	if data := a.Preferences().StringWithFallback("enter_army_date", "none"); data != "none" {
		enterArmyDateEntry.SetText(data)
	}

	if data := a.Preferences().StringWithFallback("discharge_army_date", "none"); data != "none" {
		dischargeArmyDateEntry.SetText(data)
	}

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
			dialog.ShowError(err, w)
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

	remainedVacationBox := container.NewHBox(remainedVacation, widget.NewLabel("일"), earlyDischargeCheck)

	armySelect := widget.NewSelect([]string{"5군단", "육군훈련소"}, func(s string) {

	})
	armySelect.Selected = "5군단"

	progressTickerUpdate := widget.NewSelect([]string{"10", "50", "100", "500", "1000"}, func(s string) {
		t, _ := strconv.ParseFloat(s, 64)

		ticker.Reset(time.Duration(t) * time.Millisecond)
		a.Preferences().SetInt("progress_update_ticker", int(t))
	})
	progressTickerUpdate.Selected = fmt.Sprintf("%d", a.Preferences().IntWithFallback("progress_update_ticker", 100))
	tickerUpdateBox := container.NewHBox(progressTickerUpdate, widget.NewLabel("ms"))

	progressDecimal := widget.NewSelect([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, func(s string) {
		point, _ := strconv.Atoi(s)

		a.Preferences().SetInt("progress_decimal_point", point)
	})
	progressDecimal.SetSelected(fmt.Sprintf("%d", a.Preferences().IntWithFallback("progress_decimal_point", 6)))

	progressDecimalBox := container.NewHBox(progressDecimal, widget.NewLabel("자리"))

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

	form.Append("입대일", enterArmyDateEntry)
	form.Append("전역일", dischargeArmyDateEntry)
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
