package gui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/vladsendrix/go-movies/controller"
	"github.com/vladsendrix/go-movies/entities"
)

func StartGUI(movieController *controller.MovieController) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Movies")
	myWindow.Resize(fyne.NewSize(800, 600))

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Enter Movie Title")

	yearEntry := widget.NewEntry()
	yearEntry.SetPlaceHolder("Enter Release Year")

	directorIDEntry := widget.NewEntry()
	directorIDEntry.SetPlaceHolder("Enter Director ID")

	addButton := widget.NewButton("Add Movie", func() {
		title := titleEntry.Text
		yearStr := yearEntry.Text
		directorIDStr := directorIDEntry.Text

		if title == "" || yearStr == "" || directorIDStr == "" {
			dialog.ShowInformation("Error", "Please fill in all fields", myWindow)
			return
		}

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		directorID, err := strconv.Atoi(directorIDStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		// Create a new Movie object
		movie := &entities.Movie{
			Title:       title,
			ReleaseYear: year,
			DirectorID:  directorID,
		}

		err = movieController.Create(movie)
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

	updateYearEntry := widget.NewEntry()
	updateYearEntry.SetPlaceHolder("Enter new Release Year")

	updateDirectorIDEntry := widget.NewEntry()
	updateDirectorIDEntry.SetPlaceHolder("Enter new Director ID")

	updateButton := widget.NewButton("Update Movie", func() {
		idStr := updateEntry.Text
		newTitle := newTitleEntry.Text
		newYearStr := updateYearEntry.Text
		newDirectorIDStr := updateDirectorIDEntry.Text

		if idStr == "" || newTitle == "" || newYearStr == "" || newDirectorIDStr == "" {
			dialog.ShowInformation("Error", "Please fill in all fields", myWindow)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		newYear, err := strconv.Atoi(newYearStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		newDirectorID, err := strconv.Atoi(newDirectorIDStr)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		// Create a new Movie object with updated details
		movie := &entities.Movie{
			Title:       newTitle,
			ReleaseYear: newYear,
			DirectorID:  newDirectorID,
		}

		err = movieController.Update(id, movie)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		dialog.ShowInformation("Success", fmt.Sprintf("Updated movie with ID %d", id), myWindow)
	})

	listButton := widget.NewButton("List Movies", func() {
		movies, err := movieController.GetAll()
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		var movieList string
		for _, movie := range movies {
			movieList += fmt.Sprintf("ID: %d, Title: %s, Release Year: %d, Director ID: %d\n", movie.ID, movie.Title, movie.ReleaseYear, movie.DirectorID)
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
		yearEntry,
		directorIDEntry,
		addButton,
		widget.NewLabel("Delete Movie:"),
		deleteEntry,
		deleteButton,
		widget.NewLabel("Update Movie:"),
		updateEntry,
		newTitleEntry,
		updateYearEntry,
		updateDirectorIDEntry,
		updateButton,
		widget.NewLabel("List Movies:"),
		listButton,
	)

	myWindow.SetContent(inputContainer)
	myWindow.ShowAndRun()
}
