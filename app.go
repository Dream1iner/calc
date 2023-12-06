package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)

func main() {

	CalcApp := app.New()
	window := CalcApp.NewWindow("Caclulator")

	window.Resize(fyne.Size{Width: 300, Height: 400})

	Result := widget.NewLabel("")

	Inputted := ""
	Result.SetText(Inputted)

	content := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		Result, layout.NewSpacer(),
	)

	// upper row

	button_7 := widget.NewButton("   7   ", func() {
		Inputted += "7"
	})

	button_8 := widget.NewButton("   8   ", func() {
		Inputted += "8"
	})

	button_9 := widget.NewButton("   9   ", func() {
		Inputted += "9"
	})

	row_3 := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button_7,
		layout.NewSpacer(),
		button_8,
		layout.NewSpacer(),
		button_9,
		layout.NewSpacer())

	// next row

	button_4 := widget.NewButton("   4   ", func() {
		Inputted += "4"
	})

	button_5 := widget.NewButton("   5   ", func() {
		Inputted += "5"
	})

	button_6 := widget.NewButton("   6   ", func() {
		Inputted += "6"
	})

	row_2 := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button_4,
		layout.NewSpacer(),
		button_5,
		layout.NewSpacer(),
		button_6,
		layout.NewSpacer())

	// next row

	button_1 := widget.NewButton("   1   ", func() {
		Inputted += "1"
	})

	button_2 := widget.NewButton("   2   ", func() {
		Inputted += "2"
	})

	button_3 := widget.NewButton("   3   ", func() {
		Inputted += "3"
	})

	row_1 := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button_1,
		layout.NewSpacer(),
		button_2,
		layout.NewSpacer(),
		button_3,
		layout.NewSpacer())

	// lower row

	button_mul := widget.NewButton("   *   ", func() {
		Inputted += "*"
	})

	button_0 := widget.NewButton("   0   ", func() {
		Inputted += "0"
	})

	button_sub := widget.NewButton("   -   ", func() {
		Inputted += "-"
	})

	// math row

	button_div := widget.NewButton("   /   ", func() {
		Inputted += "/"
	})

	button_dot := widget.NewButton("  .  ", func() {
		Inputted += "."
	})

	button_add := widget.NewButton("   +   ", func() {
		Inputted += "+"
	})

	button_clr := widget.NewButton("  CLR  ", func() {
		Inputted += ""
	})

	// Clear button is now working, actually

	button_eq := widget.NewButton("   =   ", func() {
		expression, err := govaluate.NewEvaluableExpression(Inputted)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				Inputted = strconv.FormatFloat(result.(float64), 'f', -1, 64)
			} else {
				Inputted = err.Error()
			}
		} else {
			Inputted = err.Error()
		}
	})

	row_0 := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button_0,
		layout.NewSpacer(),
		button_dot,
		layout.NewSpacer(),
		button_eq,
		layout.NewSpacer())

	math_row := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		button_mul,
		layout.NewSpacer(),
		button_div,
		layout.NewSpacer(),
		button_add,
		layout.NewSpacer(),
		button_sub,
		layout.NewSpacer(),
		button_clr,
		layout.NewSpacer())

	go func() {
		for {
			Result.SetText(Inputted)
		}
	}()

	window.SetContent(container.New(layout.NewVBoxLayout(), content, row_3, row_2, row_1, row_0, math_row))

	window.ShowAndRun()

}
