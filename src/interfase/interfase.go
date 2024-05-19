package interfase

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Interfase() {
	a := app.New()
	w := a.NewWindow("calculator")
	w.Resize(fyne.NewSize(356, 450))

	answer := widget.NewLabel("ответ")
	answer.Move(fyne.NewPos(20, 20))

	w.SetContent(container.NewWithoutLayout(answer))
	w.ShowAndRun()
}
