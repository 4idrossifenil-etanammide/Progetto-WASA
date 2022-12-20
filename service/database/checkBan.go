package database

import "errors"

// Creation of a custom error for the ban operation
var ErrBan = errors.New("The selected user is banned")

func (db *appdbimpl) CheckBan(userId string, banned_user string) error {

	// Collect information from the database
	rows, err := db.c.Query(`SELECT * FROM Ban WHERE Utente = ? AND Bannato = ?`, userId, banned_user)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	// If the result is not empty, then the user selected is already banned, so we return an error
	if rows.Next() {
		return ErrBan
	}
	if err := rows.Err(); err != nil {
		return err
	}

	// Finally, we return nil if everything was fine
	return nil
}
