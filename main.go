package main

import (
	"fmt"
	"github.com/flopp/go-findfont"
	"github.com/hiholder/chatgpt-demo/ui"
	"os"
	"strings"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	fmt.Println("=============")
}
func main()  {
	ui.MainWindow()
	os.Unsetenv("FYNE_FONT")
}
