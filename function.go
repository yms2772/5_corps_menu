package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
)

func UpdateDate(a fyne.App) {
	enterArmyDate = a.Preferences().StringWithFallback("enter_army_date", "none")
	dischargeArmyDate = a.Preferences().StringWithFallback("discharge_army_date", "none")
	vacationDay = a.Preferences().IntWithFallback("vacation", 0)

	enterArmyDateTime, _ = time.Parse("20060102", enterArmyDate)
	dischargeArmyDateTime, _ = time.Parse("20060102", dischargeArmyDate)

	totalArmyDateNanoSeconds = dischargeArmyDateTime.Sub(enterArmyDateTime).Nanoseconds()
	totalWOVArmyDateNanoSeconds = dischargeArmyDateTime.AddDate(0, 0, -vacationDay).Sub(enterArmyDateTime).Nanoseconds()

	diff = float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalArmyDateNanoSeconds)
	diffWOV = float64(time.Now().Sub(enterArmyDateTime).Nanoseconds()) / float64(totalWOVArmyDateNanoSeconds)
}

func FixedDecimal(num float64, point int) string {
	numSplit := strings.Split(fmt.Sprintf("%.10f", num), ".")

	if len(numSplit) != 2 {
		return "0"
	}

	return fmt.Sprintf("%s.%s", numSplit[0], numSplit[1][:point])
}

func DeleteEmptyMenu(menu []string) []string {
	var optimizedMenu []string
	for _, str := range menu {
		if str != "" {
			optimizedMenu = append(optimizedMenu, str)
		}
	}
	return optimizedMenu
}

func LoadMenu() (data OpenAPI, err error) {
	resp, err := http.Get("https://openapi.mnd.go.kr/3234313636323136353532303332303936/json/DS_TB_MNDT_DATEBYMLSVC_7369/0/2397/")
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &data)

	return
}

func CheckDayExist(data OpenAPI, date string) (ok bool) {
	ok = false

	for _, item := range data.Data.Row {
		if date == item.Dates {
			ok = true
		}
	}

	return
}

func GenerateMenuCard(data OpenAPI, date string) (menu map[string]map[string][]string) {
	log.Println("Get: " + date)

	menu = make(map[string]map[string][]string)

	for _, item := range data.Data.Row {
		if _, ok := menu[item.Dates]; !ok {
			menu[item.Dates] = make(map[string][]string)
		}

		if date == item.Dates {
			menu[item.Dates]["아침"] = append(menu[item.Dates]["아침"], item.Brst)
			menu[item.Dates]["점심"] = append(menu[item.Dates]["점심"], item.Lunc)
			menu[item.Dates]["저녁"] = append(menu[item.Dates]["저녁"], item.Dinr)
		}
	}

	menu[date]["아침"] = DeleteEmptyMenu(menu[date]["아침"])
	menu[date]["점심"] = DeleteEmptyMenu(menu[date]["점심"])
	menu[date]["저녁"] = DeleteEmptyMenu(menu[date]["저녁"])

	log.Println("Breakfast: " + strings.Join(menu[date]["아침"], "|"))
	log.Println("Lunch: " + strings.Join(menu[date]["점심"], "|"))
	log.Println("Dinner: " + strings.Join(menu[date]["저녁"], "|"))
	log.Println("Done: " + date)

	return
}
