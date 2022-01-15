package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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
