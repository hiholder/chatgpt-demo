package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)


func MainWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ChatGPT Demo")
	myWindow.Resize(fyne.NewSize(800, 500))
	tabs := container.NewAppTabs(
		container.NewTabItem("chat", NewChatWindow().chatUi()),
		container.NewTabItem("text gen", NewTextWindow().textGenUi()),
		container.NewTabItem("image gen", NewImageWindow().imageGenUi()),
		container.NewTabItem("edit text", NewEditWindow().EditUi()),
		)
	myWindow.SetContent(tabs)
	myWindow.Show()
	myApp.Run()
}

