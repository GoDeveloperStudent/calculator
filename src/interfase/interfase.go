package interfase

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func interfase() {
	a := app.New()
	w := a.NewWindow("calculator")
	w.Resize(fyne.NewSize(300, 356))

	w.SetContent(container.NewWithoutLayout())
	w.ShowAndRun()
}
