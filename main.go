package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"
)

func main() {
	a := app.New()
	w := a.NewWindow("Demo")
	if icon, err := fyne.LoadResourceFromPath("demo.png"); err == nil {
		w.SetIcon(icon)
	}

	text := binding.NewString()
	label := widget.NewLabel("Demo Label")
	textFloat := binding.StringToFloat(text)
	w.SetContent(container.NewVBox(
		label,
		widget.NewButton("Click Me", func() {
			label.SetText("You Clicked")
			if v, e := textFloat.Get(); e == nil {
				text.Set(fmt.Sprintf("%f", v + 1))
			}
		}),
		widget.NewLabelWithData(text),
		widget.NewEntryWithData(text),
		widget.NewSliderWithData(0, 1000.0, textFloat),
	))
	w.ShowAndRun()
}
