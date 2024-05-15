package gui

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartGUI() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Movies")

    listButton := widget.NewButton("List Movies", func() {
        fmt.Println("List Movies")
    })

    addButton := widget.NewButton("Add Movie", func() {
        fmt.Println("Add Movie")
    })

    idEntry := widget.NewEntry()
    idEntry.SetPlaceHolder("Enter Movie ID")

    deleteButton := widget.NewButton("Delete Movie", func() {
        id := idEntry.Text
        fmt.Println("Delete Movie", id)
    })

    content := container.NewVBox(listButton, addButton, idEntry, deleteButton)
    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}

func main() {
	StartGUI()
}