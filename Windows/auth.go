package windows

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Auth() {
	a := app.New()

	w := a.NewWindow("Авторизация")

	w.Resize(fyne.NewSize(750, 300))

	login := widget.NewEntry()
	login.SetPlaceHolder("login")

	password := widget.NewEntry()
	password.SetPlaceHolder("password")
	password.Password = true

	button_login := widget.NewButton("login", func() {
		log.Println("login:", login.Text)
		log.Println("password:", password.Text)
	})
	button_login.Icon = theme.ConfirmIcon()

	button_new_user := widget.NewButton("new user", func() { log.Println("new user") })
	button_new_user.Icon = theme.AccountIcon()

	content := container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			layout.NewSpacer(),
			container.NewVBox(login, password, button_login, button_new_user),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	w.SetContent(content)

	w.ShowAndRun()

}
