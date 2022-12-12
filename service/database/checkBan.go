package database

import "errors"

func (db *appdbimpl) CheckBan(userId string, banned_user string) error {

	rows, err := db.c.Query(`SELECT * FROM Ban WHERE Utente = ? AND Bannato = ?`, userId, banned_user)

	if rows.Next() {
		return errors.New("The selected user is banned")
	}
	rows.Close()

	return err
}
