//go:build gui

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func launchGui() {
	a := app.New()
	w := a.NewWindow("Go Gin Project Initializer")

	moduleEntry := widget.NewEntry()
	moduleEntry.SetPlaceHolder("e.g., github.com/user/project")

	initGitCheck := widget.NewCheck("Initialize Git", nil)
	fullApiCheck := widget.NewCheck("Full API structure", nil)

	generateButton := widget.NewButton("Generate", func() {
		moduleName := moduleEntry.Text
		if moduleName == "" {
			dialog.ShowError(fmt.Errorf("module name is required"), w)
			return
		}

		err := generateProject(moduleName, fullApiCheck.Checked, initGitCheck.Checked)
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			dialog.ShowInformation("Success", "✅ Project structure initialized.", w)
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.Close()
	})

	buttons := container.NewHBox(generateButton, cancelButton)

	form := container.NewVBox(
		widget.NewLabel("Module Name:"),
		moduleEntry,
		initGitCheck,
		fullApiCheck,
		buttons,
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(400, 200))
	w.ShowAndRun()
}
