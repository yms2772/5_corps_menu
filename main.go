package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	mainApp := app.NewWithID("mokky.mnd.menu")
	mainApp.Settings().SetTheme(&myTheme{})

	mainWindow := mainApp.NewWindow("5군단 식단")

	mainWindow.SetContent(container.NewBorder(progressBar(mainApp, mainWindow), nil, nil, nil, container.NewAppTabs(menuItem(mainWindow), settingsItem(mainApp, mainWindow))))
	mainWindow.ShowAndRun()
}
