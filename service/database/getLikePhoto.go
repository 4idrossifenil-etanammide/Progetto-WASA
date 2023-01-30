package database

func (db *appdbimpl) GetLikePhoto(photoId string, user string) (bool, error) {

	// Collect the ID of the user who want to like
	rows, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?;`, user)
	if err != nil {
		return false, err
	}

	isNotEmpty := rows.Next()
	if err := rows.Err(); err != nil {
		return false, err
	}

	var userId string
	if isNotEmpty {
		err = rows.Scan(&userId)
		if err != nil {
			return false, err
		}
	} else {
		return false, ErrUserDoesntExist
	}

	rows.Close() // -> Required, otherwise the database remains locked
	rows, err = db.c.Query(`SELECT * FROM Like WHERE FotoReference = ? AND Utente = ?;`, photoId, userId)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		rows.Close()
		return true, nil
	}
	if err := rows.Err(); err != nil {
		return false, err
	}

	rows.Close()
	return false, nil

}
