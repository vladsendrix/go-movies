package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/vladsendrix/go-movies/controller"
	"github.com/vladsendrix/go-movies/entities"
)

func StartGUI(movieController *controller.MovieController) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Movies")

	listButton := widget.NewButton("List Movies", func() {
		movies, err := movieController.GetAll()
		if err != nil {
			fmt.Println("Error listing movies:", err)
			return
		}
		for _, movie := range movies {
			fmt.Println(movie)
		}
	})

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Enter Movie Title")

	addButton := widget.NewButton("Add Movie", func() {
		title := titleEntry.Text
		if title == "" {
			fmt.Println("Please enter a movie title")
			return
		}

		// Create a new Movie object
		movie := &entities.Movie{
			Title: title,
		}

		err := movieController.Create(movie)
		if err != nil {
			fmt.Println("Error adding movie:", err)
			return
		}
		fmt.Println("Added movie:", title)
	})

	// Assuming you have a deleteEntry for entering the ID of the movie to delete
	deleteEntry := widget.NewEntry()
	deleteEntry.SetPlaceHolder("Enter Movie ID to delete")

	deleteButton := widget.NewButton("Delete Movie", func() {
		ID := deleteEntry.Text
		if ID == "" {
			fmt.Println("Please enter a movie ID")
			return
		}

		err := movieController.Delete(ID)
		if err != nil {
			fmt.Println("Error deleting movie:", err)
			return
		}
		fmt.Println("Deleted movie with ID:", ID)
	})

	// Assuming you have an updateEntry for entering the new title of the movie to update
	updateEntry := widget.NewEntry()
	updateEntry.SetPlaceHolder("Enter new Movie Title")

	updateButton := widget.NewButton("Update Movie", func() {
		idStr := deleteEntry.Text
		newTitle := updateEntry.Text
		if idStr == "" || newTitle == "" {
			fmt.Println("Please enter a movie ID and new title")
			return
		}

		// Create a new Movie object with the new title
		movie := &entities.Movie{
			Title: newTitle,
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid movie ID:", err)
			return
		}

		err = movieController.Update(id, movie)
		if err != nil {
			fmt.Println("Error updating movie:", err)
			return
		}
		fmt.Println("Updated movie with ID:", id)
	})

	myWindow.SetContent(container.NewVBox(
		listButton,
		titleEntry,
		addButton,
		deleteEntry,
		deleteButton,
		updateEntry,
		updateButton,
	))
	myWindow.ShowAndRun()
}
