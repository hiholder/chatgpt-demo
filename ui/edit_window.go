package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/hiholder/chatgpt-demo/openai"
)

type EditWindow struct {
	editInput *widget.Entry
	editInstruction *widget.Entry
	editOutput *widget.Label
}

func NewEditWindow() *EditWindow {
	return &EditWindow{
		editInput: widget.NewEntry(),
		editInstruction: widget.NewEntry(),
		editOutput: widget.NewLabelWithStyle("wait...", fyne.TextAlignCenter, fyne.TextStyle{}),
	}
}

func (w *EditWindow) EditUi() *fyne.Container {
	w.editInput.SetPlaceHolder("Please input...")
	w.editInstruction.SetPlaceHolder("Please input...")
	w.editOutput.Resize(fyne.NewSize(150, 100))
	w.editOutput.Wrapping = 3
	grid := container.New(layout.NewGridLayout(2), widget.NewButton("submit", w.editSubmitAction), widget.NewButton("clear", w.editClearAction))
	content := container.NewVBox(w.editInput, w.editInstruction, grid, w.editOutput)
	return content
}

func (w *EditWindow) editSubmitAction()  {
	text, err := openai.EditClient.GetEdit(w.editInput.Text, w.editInstruction.Text)
	if err != nil {
		return
	}
	w.editOutput.SetText(fmt.Sprintf("%s", text))
	w.editOutput.Refresh()
}

func (w *EditWindow) editClearAction()  {
	w.editInput.SetText("")
	w.editInput.Refresh()
	w.editInstruction.SetText("")
	w.editInstruction.Refresh()
	w.editOutput.SetText("")
	w.editOutput.Refresh()
}

