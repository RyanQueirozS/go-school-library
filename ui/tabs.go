package ui

import (
	"fmt"
	"go-school-library/database"
	"go-school-library/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeTabs() fyne.Widget {
	tabs := container.NewAppTabs(
		container.NewTabItem("Rent Tab", container.NewWithoutLayout(
			makeRentDetail(getRentDetail()),
			makeRentCreation(),
			makeBookReturn(),
		),
		),
		container.NewTabItem(
			"Creation Tab",
			container.NewVBox(
				container.NewHBox(
					makeAccountCreation(),
					makeBookCreation(),
				),
				container.NewHBox(
					makeAccountDeletion(),
					makeBookDeletion(),
				),
			),
		),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	return tabs
}

type rent struct {
	accountID    int
	accountName  string
	bookID       int
	bookName     string
	deliveryDate string
}

func makeRentDetail(rentDetail [][]string, err error) fyne.Widget {
	if err != nil {
		return widget.NewLabel(err.Error())
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(rentDetail), len(rentDetail[0])
		},
		func() fyne.CanvasObject {
			item := widget.NewLabel("Rent Tab")
			item.Resize(fyne.Size{
				Width:  400,
				Height: 100,
			})

			return item
		},
		func(i widget.TableCellID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(
				fmt.Sprintf("%v", rentDetail[i.Row][i.Col]),
			)
			item.Resize(fyne.Size{
				Width:  200,
				Height: 20,
			})
		},
	)
	table.Resize(fyne.NewSize(400, 200))
	table.Move(table.Position().AddXY(10, 30))
	return table
}

func getRentDetail() ([][]string, error) {
	var allRentDetail [][]string

	err := database.CreateAccountBooksTable()
	if err != nil {
		return allRentDetail, err
	}
	db, err := database.GetDB()
	if err != nil {
		return allRentDetail, err
	}
	rows, err := db.Query("SELECT * FROM accountBooks")
	if err != nil {
		return allRentDetail, err
	}
	for rows.Next() {
		var rentDetail rent
		err = rows.Scan(
			&rentDetail.accountID,
			&rentDetail.accountName,
			&rentDetail.bookID,
			&rentDetail.bookName,
			&rentDetail.deliveryDate,
		)
		var rentString []string
		rentString = append(
			rentString,
			strconv.Itoa(rentDetail.accountID),
			rentDetail.accountName,
			strconv.Itoa(rentDetail.bookID),
			rentDetail.bookName,
			rentDetail.deliveryDate,
		)
		allRentDetail = append(allRentDetail, rentString)
		if err != nil {
			return allRentDetail, err
		}
	}
	if len(allRentDetail) == 0 {
		allRentDetail = append(allRentDetail, []string{"EMPTY"})
	}
	return allRentDetail, err
}

func makeAccountCreation() fyne.Widget {
	NameEntry := widget.NewEntry()
	CPFEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("Name:", NameEntry),
		widget.NewFormItem("CPF:", CPFEntry),
	)

	form.OnCancel = func() {
		NameEntry.Text = ""
		CPFEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		models.CreateAccount(NameEntry.Text, CPFEntry.Text)
		NameEntry.Text = ""
		CPFEntry.Text = ""
		form.Refresh()
	}

	btn := widget.NewButton("Create Account", func() {
		w := a.NewWindow("Account Creation")
		w.SetContent(form)
		w.Show()
	})
	return btn
}

func makeBookCreation() fyne.Widget {
	NameEntry := widget.NewEntry()
	DescriptionEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("Name:", NameEntry),
		widget.NewFormItem("Description:", DescriptionEntry),
	)

	form.OnCancel = func() {
		NameEntry.Text = ""
		DescriptionEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		models.CreateBook(NameEntry.Text, DescriptionEntry.Text)
		NameEntry.Text = ""
		DescriptionEntry.Text = ""
		form.Refresh()
	}
	btn := widget.NewButton("Create Book", func() {
		w := a.NewWindow("Book Creation")
		w.SetContent(form)
		w.Show()
	})
	return btn
}

func makeAccountDeletion() fyne.Widget {
	IDEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("ID:", IDEntry),
	)

	form.OnCancel = func() {
		IDEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		id, _ := strconv.Atoi(IDEntry.Text)
		models.DeleteAccount(id)
		IDEntry.Text = ""
		form.Refresh()
	}

	btn := widget.NewButton("Delete Account", func() {
		w := a.NewWindow("Account Deletion")
		w.SetContent(form)
		w.Show()
	})
	return btn
}

func makeBookDeletion() fyne.Widget {
	IDEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("ID:", IDEntry),
	)

	form.OnCancel = func() {
		IDEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		id, _ := strconv.Atoi(IDEntry.Text)
		models.DeleteBook(id)
		IDEntry.Text = ""
		form.Refresh()
	}

	btn := widget.NewButton("Delete Book", func() {
		w := a.NewWindow("Book Deletion")
		w.SetContent(form)
		w.Show()
	})
	return btn
}

func makeRentCreation() fyne.Widget {
	AccountEntry := widget.NewEntry()
	BookEntry := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("Account ID:", AccountEntry),
		widget.NewFormItem("Book ID:", BookEntry),
	)

	form.OnCancel = func() {
		AccountEntry.Text = ""
		BookEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		accountId, _ := strconv.Atoi(AccountEntry.Text)
		bookId, _ := strconv.Atoi(BookEntry.Text)
		models.RentBook(accountId, bookId)
		AccountEntry.Text = ""
		BookEntry.Text = ""
		form.Refresh()
	}

	btn := widget.NewButton("Make Rent", func() {
		w := a.NewWindow("Book Rent")
		w.SetContent(form)
		w.Show()
	})
	btn.Resize(fyne.NewSize(150, 50))
	btn.Move(btn.Position().AddXY(420, 30))
	return btn
}

func makeBookReturn() fyne.Widget {
	AccountEntry := widget.NewEntry()
	BookEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Account ID:", AccountEntry),
		widget.NewFormItem("Book ID:", BookEntry),
	)

	form.OnCancel = func() {
		AccountEntry.Text = ""
		BookEntry.Text = ""
		form.Refresh()
	}
	form.OnSubmit = func() {
		accountId, _ := strconv.Atoi(AccountEntry.Text)
		bookId, _ := strconv.Atoi(BookEntry.Text)
		models.ReturnBook(accountId, bookId)
		AccountEntry.Text = ""
		BookEntry.Text = ""
		form.Refresh()
	}

	btn := widget.NewButton("Return Book", func() {
		w := a.NewWindow("Book Return")
		w.SetContent(form)
		w.Show()
	})

	btn.Resize(fyne.NewSize(150, 50))
	btn.Move(btn.Position().AddXY(420, 150))
	return btn
}
