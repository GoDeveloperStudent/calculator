package interfase

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Interfase() {
	a := app.New()
	w := a.NewWindow("calculator")
	w.Resize(fyne.NewSize(356, 450))

	answer := widget.NewLabel("ответ")
	answer.Move(fyne.NewPos(20, 20))

	answerRect := canvas.NewRectangle(color.NRGBA{255, 255, 255, 255})
	answerRect.StrokeColor = color.Black
	answerRect.StrokeWidth = 1
	answerRect.Resize(fyne.NewSize(260, 50))
	answerRect.Move(fyne.NewPos(21, 20))

	w.SetContent(container.NewWithoutLayout(answerRect, answer))
	w.ShowAndRun()
}
