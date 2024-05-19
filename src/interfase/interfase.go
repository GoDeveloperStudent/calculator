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

	entry1 := widget.NewMultiLineEntry()
	entry1.Wrapping = fyne.TextWrapOff
	entry1.SetPlaceHolder("first number")
	entry1.Resize(fyne.NewSize(140, 50))
	entry1.Move(fyne.NewPos(20, 80))

	entry2 := widget.NewMultiLineEntry()
	entry2.Wrapping = fyne.TextWrapOff
	entry2.SetPlaceHolder("second number")
	entry2.Resize(fyne.NewSize(140, 50))
	entry2.Move(fyne.NewPos(187, 80))

	w.SetContent(container.NewWithoutLayout(entry2, entry1, answerRect, answer))
	w.ShowAndRun()
}
