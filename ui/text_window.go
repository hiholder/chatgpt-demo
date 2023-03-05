package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/hiholder/chatgpt-demo/openai"
)

type  TextWindow struct {
	textInput *widget.Entry
	textOutput *widget.Label
}

func NewTextWindow() *TextWindow {
	return &TextWindow{
		textInput: widget.NewEntry(),
		textOutput: widget.NewLabelWithStyle("wait...", fyne.TextAlignCenter, fyne.TextStyle{}),
	}
}


func (w *TextWindow)textGenUi() *fyne.Container {
	w.textInput.SetPlaceHolder("Please input...")
	w.textOutput.Resize(fyne.NewSize(150, 100))
	w.textOutput.Wrapping = 3
	grid := container.New(layout.NewGridLayout(2), widget.NewButton("submit", w.textSubmitAction), widget.NewButton("clear", w.textClearAction))
	content := container.NewVBox(w.textInput, grid, w.textOutput)
	return content
}


func (w *TextWindow)textSubmitAction()  {
	answer, err := openai.TextClient.GetText(w.textInput.Text)
	if err != nil {
		return
	}
	w.textOutput.SetText(fmt.Sprintf("answer is %s", answer))
	w.textOutput.Refresh()
}

func (w *TextWindow)textClearAction()  {
	w.textInput.SetText("")
	w.textInput.Refresh()
	w.textOutput.SetText("")
	w.textOutput.Refresh()
}