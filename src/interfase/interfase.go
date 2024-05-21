package interfase

import (
	"GoMod/src/initfunc"
	"image/color"
	"strconv"

	"fmt"

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
	answerRect.Resize(fyne.NewSize(250, 50))
	answerRect.Move(fyne.NewPos(21, 20))

	answerRadio := widget.NewRadioGroup([]string{"+", "-", "*", "/"}, func(s string) {})

	accARI := widget.NewAccordionItem("", answerRadio)

	accAR := widget.NewAccordion(accARI)
	accAR.Resize(fyne.NewSize(310, 50))
	accAR.Move(fyne.NewPos(20, 220))

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

	answerBtn := widget.NewButton("проверить", func() {
		entry1Str := entry1.Text
		entry1Fl, err1 := strconv.ParseFloat(entry1Str, 32)

		entry2Str := entry2.Text
		entry2Fl, err2 := strconv.ParseFloat(entry2Str, 32)
		if err1 != nil || err2 != nil {
			answer.SetText("error")
			fmt.Println(err1)
			fmt.Println(err2)
			answerRect.StrokeColor = color.NRGBA{255, 0, 0, 255}
		} else {
			answer.SetText(initfunc.Initfunc(float32(entry1Fl), float32(entry2Fl), answerRadio.Selected))
		}
	})
	answerBtn.Resize(fyne.NewSize(310, 50))
	answerBtn.Move(fyne.NewPos(20, 150))

	w.SetContent(container.NewWithoutLayout(accAR, answerBtn, entry2, entry1, answerRect, answer))
	w.ShowAndRun()
}
