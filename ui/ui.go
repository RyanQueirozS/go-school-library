package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"go-school-library/models"
)

var (
	a = app.New()
	w = a.NewWindow("Library")
)

func MakeUI() {
	w.Resize(fyne.NewSize(600, 150))
	w.CenterOnScreen()

	tabs := container.NewAppTabs(
		container.NewTabItem("Accounts", displayAccount()))
	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.ShowAndRun()
}

func displayAccount() fyne.Widget {
	accountString := widget.NewLabel(accountToString())

	tabs := container.NewAppTabs(
		container.NewTabItem("Accounts", accountString))
	tabs.SetTabLocation(container.TabLocationTop)

	content := widget.NewButton("Show all accounts", func() {
		w.SetContent(tabs)
	})

	content.Resize(fyne.NewSize(100, 60))
	return content
}

func accountToString() string {
	var accounts []models.Account
	accounts, err := models.GetAllAccounts()
	var accountString string = ""

	if err != nil || len(accounts) == 0 {
		return err.Error()
	} else {
		for _, v := range accounts {
			accountString = accountString + fmt.Sprintf("%d %s\n", models.GetAccountID(v), models.GetAccountName(v))
		}
		return accountString
	}
}
