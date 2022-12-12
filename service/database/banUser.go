package database

func (db *appdbimpl) BanUser(userId string, banned_user string) error {

	rows, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, banned_user)

	var output string

	if rows.Next() {
		rows.Scan(&output)
	}
	rows.Close()

	err = db.CheckBan(userId, output)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`INSERT INTO Ban(Utente, Bannato) VALUES (?,?);`, userId, output)

	if err != nil {
		return err
	}

	return nil
}
