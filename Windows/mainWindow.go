package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func MainWindow() {
	a := app.New()
	w := a.NewWindow("r2")
	w.Resize(fyne.NewSize(500, 500))

	w.ShowAndRun()
}
