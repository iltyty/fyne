// Package main loads a very basic Hello World graphical application.
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	// hello := widget.NewLabel("Hello Fyne!")
	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome 😀")
	// 	}),
	// ))

	l := widget.NewCustomList(func() int {
		return 1
	}, func() fyne.CanvasObject {
		return canvas.NewText("hello", nil)
	}, func(i widget.CustomListItemID) fyne.CanvasObject {
		// return widget.NewButton("button", func() {})
		return canvas.NewText("world", nil)
	})
	l.OnSelected = func(id widget.CustomListItemID) {
		fmt.Println(id)
	}
	w.SetContent(l)

	w.ShowAndRun()
}
