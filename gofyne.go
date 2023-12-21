package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("點擊")

	score := 0
	scoreLabel := widget.NewLabel("Score: 0")

	clickButton := widget.NewButton("buttom", func() {
		score++
		scoreLabel.SetText("Score: " + fmt.Sprint(score))
	})

	content := container.New(layout.NewVBoxLayout(),
		scoreLabel,
		clickButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
