package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/hiholder/chatgpt-demo/openai"
)

type ChatWindow struct {
	chatInput *widget.Entry
	chatOutput *widget.Label
	chatScroll *container.Scroll
}

func NewChatWindow() *ChatWindow {
	return &ChatWindow{
		chatInput: widget.NewEntry(),
		chatOutput: widget.NewLabelWithStyle("wait...", fyne.TextAlignCenter, fyne.TextStyle{}),
	}
}

func (w *ChatWindow)chatUi() *fyne.Container {
	w.chatOutput.Wrapping = 3
	w.chatInput.SetPlaceHolder("Please input...")
	grid := container.New(layout.NewGridLayout(2), widget.NewButton("submit", w.chatSubmitAction), widget.NewButton("clear", w.chatClearAction))
	w.chatScroll = container.NewScroll(w.chatOutput)
	w.chatScroll.SetMinSize(fyne.NewSize(250, 400))
	content := container.NewVBox(w.chatInput, grid, w.chatScroll)
	//content := container.NewVBox(form)
	return content
}


func (w *ChatWindow)chatSubmitAction()  {
	answer := openai.ChatClient.GetAnswer(w.chatInput.Text)
	w.chatOutput.SetText(fmt.Sprintf("answer is %s", answer))
	w.chatOutput.Resize(fyne.NewSize(300,200))
	w.chatOutput.Refresh()
}

func (w *ChatWindow)chatClearAction()  {
	w.chatInput.SetText("")
	w.chatInput.Refresh()
	w.chatOutput.SetText("")
	w.chatOutput.Refresh()
}
