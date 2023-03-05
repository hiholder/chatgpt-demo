package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/hiholder/chatgpt-demo/openai"
)

type ImageWindow struct {
	imageInput *widget.Entry
	imageContent *fyne.Container
}

func NewImageWindow() *ImageWindow {
	return &ImageWindow{
		imageInput: widget.NewEntry(),
	}
}

func (w *ImageWindow)imageGenUi() *fyne.Container {
	w.imageInput.SetPlaceHolder("Please input...")
	w.imageContent = container.NewVBox(w.imageInput, widget.NewButton("gen", w.imageClickAction))
	return w.imageContent
}

func (w *ImageWindow)imageClickAction()  {
	uris, err  := openai.ImageClient.GetImage(w.imageInput.Text)
	if err != nil {
		return
	}
	//uri := "https://tse3-mm.cn.bing.net/th/id/OIP-C.-LK9fF9heGJYfhyakkrWlwHaE8?pid=ImgDet&rs=1"
	if err != nil {
		return
	}
	for _, img := range uris {
		urlString, err := fyne.LoadResourceFromURLString(img)
		if err != nil {
			return
		}
		genImage := canvas.NewImageFromImage(nil)
		genImage.FillMode = canvas.ImageFillOriginal
		genImage.Resource = urlString
		genImage.Resize(fyne.NewSize(512, 512))
		w.imageContent.Add(genImage)
	}
	w.imageContent.Layout = layout.NewCenterLayout()
	w.imageContent.Refresh()
}