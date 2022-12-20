package main

// https://www.youtube.com/watch?v=4C1IgIOSxLM&list=PL-Jc9J83PIiGhwxgJGiYxGG7zChL1PmAx&index=6&ab_channel=Pepcoding
import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	output := ""
	input := widget.NewLabel(output)

	historyBtn := widget.NewButton("history", func(){})
	
	backBtn := widget.NewButton("back", func(){
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})

	clearBtn := widget.NewButton("clear", func(){
		output = ""
		input.SetText(output)
	})

    openBracketBtn := widget.NewButton("(", func(){
		output = output+"("
		input.SetText(output)
	})

	closeBracketBtn := widget.NewButton(")", func(){
		output = output+")"
		input.SetText(output)
	})

	starBtn := widget.NewButton("*", func(){
		output = output+"*"
		input.SetText(output)
	})

	doptBtn := widget.NewButton(".", func(){
		output = output+"."
		input.SetText(output)
	})

	plusBtn := widget.NewButton("+", func(){
		output = output+"+"
		input.SetText(output)
	})
	
	substractBtn := widget.NewButton("-", func(){
		output = output+"-"
		input.SetText(output)
	})
	
	slashBtn := widget.NewButton("/", func(){
		output = output+"/"
		input.SetText(output)
	})

	

	oneBtn := widget.NewButton("1", func(){
		output = output+"1"
		input.SetText(output)
	})
	
	twoBtn := widget.NewButton("2", func ()  {
		output = output+"2"
		input.SetText(output)
	})

	threeBtn := widget.NewButton("3", func ()  {
		output = output+"3"
		input.SetText(output)
	})

	fourBtn := widget.NewButton("4", func ()  {
		output = output+"4"
		input.SetText(output)
	})

	fiveBtn := widget.NewButton("5", func ()  {
		output = output+"5"
		input.SetText(output)
	})

	sixBtn := widget.NewButton("6", func ()  {
		output = output+"6"
		input.SetText(output)
	})

	sevenBtn := widget.NewButton("7", func ()  {
		output = output+"7"
		input.SetText(output)
	})

	eightBtn := widget.NewButton("8", func ()  {
		output = output+"8"
		input.SetText(output)
	})

	nineBtn := widget.NewButton("9",func ()  {
		output = output+"9"
		input.SetText(output)
	})

	zeroBtn := widget.NewButton("0",func ()  {
		output = output+"0"
		input.SetText(output)
	})

	equalBtn := widget.NewButton("=",func ()  {

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				output = strconv.FormatFloat(result.(float64),'f', -1, 64)
			}else {
				output = "error"
			}
		}else {
			output = "error"
		}

		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		
		input,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
			historyBtn,
			backBtn,
			),

			container.NewGridWithColumns(4,
			clearBtn,
			openBracketBtn,
			closeBracketBtn,
			slashBtn,
			),

			container.NewGridWithColumns(4,
			sevenBtn,
			eightBtn,
			nineBtn,
			starBtn,
			),

			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				substractBtn,
			),

			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),

			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					doptBtn,
				),

				container.NewGridWithColumns(1,
					equalBtn,
				),
			),


		),
	))

	w.ShowAndRun()
}
