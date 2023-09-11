package models

import "go-school-library/database"

func GetAccountID(a Account) int {
	return a.id
}

func GetAccountName(id int) string {
	var account Account
	db, err := database.GetDB()
	if err != nil {
		return err.Error()
	}
	accountRow, err := db.Query("SELECT name FROM accounts WHERE id = ?", id)
	for accountRow.Next() {

		err = accountRow.Scan(&account.name)
		if err != nil {
			return err.Error()
		}
	}
	return account.name
}

func GetAccountCPF(a Account) string {
	return a.cpf
}

func CreateAccount(name string, cpf string) string {
	db, err := database.GetDB()
	err = database.CreateAccountsTable()
	if len(cpf) == 11 {
		if err != nil {
			return err.Error()
		}
		_, err = db.Exec(
			"INSERT INTO accounts (name, cpf) VALUES (?, ?)",
			name,
			cpf,
		)
	}
	return "Conta criada com sucesso"
}

func DeleteAccount(id int) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM accounts WHERE id = ?", id)
	return err
}

func GetAllAccounts() ([]Account, error) {
	accounts := []Account{}
	db, err := database.GetDB()
	if err != nil {
		return accounts, err
	}

	rows, err := db.Query("SELECT id, name, cpf FROM accounts")
	if err != nil {
		return accounts, err
	}

	for rows.Next() {
		var account Account

		err = rows.Scan(&account.id, &account.name, &account.cpf)
		if err != nil {
			return accounts, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
