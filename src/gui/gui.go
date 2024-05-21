package gui

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

func Gui() {
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

	operationsList := widget.NewRadioGroup([]string{"+", "-", "*", "/"}, func(s string) {}) //radio group

	listItem := widget.NewAccordionItem("", operationsList) //radio group menu item

	list := widget.NewAccordion(listItem) //radio group menu
	list.Resize(fyne.NewSize(310, 50))
	list.Move(fyne.NewPos(20, 220))

	userInputleft := widget.NewMultiLineEntry()
	userInputleft.Wrapping = fyne.TextWrapOff
	userInputleft.SetPlaceHolder("first number")
	userInputleft.Resize(fyne.NewSize(140, 50))
	userInputleft.Move(fyne.NewPos(20, 80))

	userInputRight := widget.NewMultiLineEntry()
	userInputRight.Wrapping = fyne.TextWrapOff
	userInputRight.SetPlaceHolder("second number")
	userInputRight.Resize(fyne.NewSize(140, 50))
	userInputRight.Move(fyne.NewPos(187, 80))

	answerBtn := widget.NewButton("проверить", func() {
		userInputleftStr := userInputleft.Text
		userInputleftFl, err1 := strconv.ParseFloat(userInputleftStr, 32)

		userInputRightStr := userInputRight.Text
		userInputRightFl, err2 := strconv.ParseFloat(userInputRightStr, 32)
		if err1 != nil || err2 != nil {
			answer.SetText("error")
			fmt.Println(err1)
			fmt.Println(err2)
			answerRect.StrokeColor = color.NRGBA{255, 0, 0, 255}
		} else {
			answer.SetText(initfunc.Initfunc(float32(userInputleftFl), float32(userInputRightFl), operationsList.Selected))
		}
	})
	answerBtn.Resize(fyne.NewSize(310, 50))
	answerBtn.Move(fyne.NewPos(20, 150))

	w.SetContent(container.NewWithoutLayout(list, answerBtn, userInputRight, userInputleft, answerRect, answer))
	w.ShowAndRun()
}