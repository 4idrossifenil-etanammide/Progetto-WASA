package database

func (db *appdbimpl) GetBan(username string, banned_user string) (bool, error) {

	rows, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?;`, username)
	if err != nil {
		return false, err
	}

	var usernameId string
	if rows.Next() {
		err = rows.Scan(&usernameId)
		if err != nil {
			return false, err
		}
	}
	if err := rows.Err(); err != nil {
		return false, err
	}

	rows.Close()

	rows, err = db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?;`, banned_user)
	if err != nil {
		return false, err
	}

	var bannedId string
	if rows.Next() {
		err = rows.Scan(&bannedId)
		if err != nil {
			return false, err
		}
	}
	if err := rows.Err(); err != nil {
		return false, err
	}

	rows.Close()

	err = db.CheckBan(usernameId, bannedId)
	if err != nil {
		return true, nil
	}

	return false, nil
}
