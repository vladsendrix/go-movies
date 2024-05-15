package gui

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/vladsendrix/go-movies/controller"
	"github.com/vladsendrix/go-movies/entities"
)

func StartGUI(movieController *controller.MovieController) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Movies")

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Enter Movie Title")

	addButton := widget.NewButton("Add Movie", func() {
		title := titleEntry.Text
		if title == "" {
			dialog.ShowInformation("Error", "Please enter a movie title", myWindow)
			return
		}

		// Create a new Movie object
		movie := &entities.Movie{
			Title: title,
		}

		err := movieController.Create(movie)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		dialog.ShowInformation("Success", "Added movie: "+title, myWindow)
	})

	deleteEntry := widget.NewEntry()
	deleteEntry.SetPlaceHolder("Enter Movie ID to delete")

	deleteButton := widget.NewButton("Delete Movie", func() {
		idStr := deleteEntry.Text
		if idStr == "" {
			dialog.ShowInformation("Error", "Please enter a movie ID", myWindow)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		err = movieController.Delete(id)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		dialog.ShowInformation("Success", "Deleted movie with ID: "+idStr, myWindow)
	})

	updateEntry := widget.NewEntry()
	updateEntry.SetPlaceHolder("Enter Movie ID")

	newTitleEntry := widget.NewEntry()
	newTitleEntry.SetPlaceHolder("Enter new Movie Title")

	updateButton := widget.NewButton("Update Movie", func() {
		idStr := updateEntry.Text
		newTitle := newTitleEntry.Text
		if idStr == "" || newTitle == "" {
			dialog.ShowInformation("Error", "Please enter a movie ID and new title", myWindow)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		// Create a new Movie object with the updated title
		movie := &entities.Movie{
			Title: newTitle,
		}

		err = movieController.Update(id, movie)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		dialog.ShowInformation("Success", "Updated movie with ID: "+idStr, myWindow)
	})

	listButton := widget.NewButton("List Movies", func() {
		movies, err := movieController.GetAll()
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		var movieList string
		for _, movie := range movies {
			movieList += movie.Title + "\n"
		}

		if movieList == "" {
			movieList = "No movies found"
		}

		dialog.ShowInformation("Movie List", movieList, myWindow)
	})

	// Container for input fields and buttons
	inputContainer := container.NewVBox(
		widget.NewLabel("Add Movie:"),
		titleEntry,
		addButton,
		widget.NewLabel("Delete Movie:"),
		deleteEntry,
		deleteButton,
		widget.NewLabel("Update Movie:"),
		updateEntry,
		newTitleEntry,
		updateButton,
		widget.NewLabel("List Movies:"),
		listButton,
	)

	myWindow.SetContent(inputContainer)
	myWindow.ShowAndRun()
}
