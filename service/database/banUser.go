package database

import "errors"

func (db *appdbimpl) BanUser(userId string, banned_user string) error {

	// Collect the id of the user we want to ban
	rows, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, banned_user)
	if err != nil {
		return err
	}

	// Put the ID in a string
	var bannedId string
	if rows.Next() {
		err = rows.Scan(&bannedId)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	rows.Close() // -> Required, otherwise the database remains locked

	// Check if the user is not already banned (using the custom error) or if the operation on the database was not succesfully
	err = db.CheckBan(userId, bannedId)
	if errors.Is(err, ErrBan) {
		return nil
	} else if err != nil {
		return err
	}

	// Insert the information on the ban
	_, err = db.c.Exec(`INSERT INTO Ban(Utente, Bannato) VALUES (?,?);`, userId, bannedId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`DELETE FROM Segue WHERE Utente = ? AND Follower = ?;`, bannedId, userId)
	if err != nil {
		return err
	}

	// Return nil if everythin is fine
	return nil
}
