package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)


const memeURL = "http://apimeme.com/meme?meme=Bad-Advice-Cat&"

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Meme generator")

	//menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }))

	helpMenu := fyne.NewMenu("Help", fyne.NewMenuItem("About", func() {
		dialog.ShowCustom("About", "Close", container.NewVBox(widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
			widget.NewLabel("Version: v0.1"),
			widget.NewLabel("Author: Adrianna Kaminska"),
		), myWindow)
	}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)

	//when  you create new var and you assign something to it, then you use :=
	text := canvas.NewText("Hello Meme-lover!", color.White)
	//when you use feature, function of this variable and you assign value to it you use =
	text.Alignment = fyne.TextAlignCenter

	var resource, _ = fyne.LoadResourceFromURLString(memeURL)
	gopherImage := canvas.NewImageFromResource(resource)
	gopherImage.SetMinSize(fyne.Size{Width: 500, Height: 500})

	//input box
	inputTop := widget.NewEntry()
	inputTop.SetPlaceHolder("Type the top text for your meme here")

	inputBottom := widget.NewEntry()
	inputBottom.SetPlaceHolder("Type the bottom text for your meme here")

	//making button
	btn := widget.NewButton("Click me!", func() {
		resource, _ := fyne.LoadResourceFromURLString(memeURL + "top=" + inputTop.Text + "&bottom=" + inputBottom.Text)
		gopherImage.Resource = resource
		//Redrawn the image with the new path
		gopherImage.Refresh()
	})
	btn.Importance = widget.HighImportance

	box := container.NewVBox(
		text,
		gopherImage,
		inputTop,
		inputBottom,
		btn,
	)
	// Display our content
	myWindow.SetContent(box)

	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	//this actually displays
	myWindow.ShowAndRun()
}
