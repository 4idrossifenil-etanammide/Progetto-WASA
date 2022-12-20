package database

func (db *appdbimpl) UnbanUser(userId string, banned_user string) error {

	// Collect the id of the user we want to unban
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

	// Insert the information on the unban
	rows.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`DELETE FROM Ban WHERE Utente = ? AND Bannato = ?;`, userId, bannedId)
	if err != nil {
		return err
	}

	//Return nil if everything is fine
	return nil
}
