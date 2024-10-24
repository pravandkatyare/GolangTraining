package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// w := a.NewWindow("Hello World")

	//button
	// content := widget.NewButton("click me", func() {
	// 	log.Println("Clicked...!!")
	// })
	// button with icon
	// content := widget.NewButtonWithIcon("click me", theme.HomeIcon(), func() {
	// 	log.Println("Clicked...!!")
	// })
	// content.Resize(fyne.NewSize(20, 20))
	// //label
	// w.SetContent(widget.NewLabel("Hello World!"))
	// w.SetContent(content)

	// form

	w := a.NewWindow("Form Widget")
	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Test", Widget: entry},
		}, OnSubmit: func() {
			log.Println("Submitted", entry.Text)
			w.Close()
		},
	}
	textArea := widget.NewMultiLineEntry()
	form.Append("Text", textArea)

	w.SetContent(form)

	w.ShowAndRun()
}
