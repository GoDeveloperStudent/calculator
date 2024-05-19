package interfase

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func interfase() {
	a := app.New()
	w := a.NewWindow("calculator")

	w.SetContent(container.NewWithoutLayout())
	w.ShowAndRun()
}
