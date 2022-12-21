package database

func (db *appdbimpl) UnlikePhoto(photoId string, user string) error {

	// Initialize variable
	var userId string

	// Check if the photo exists
	rows, err := db.c.Query(`SELECT * FROM Foto WHERE FotoID = ?`, photoId)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return ErrPhoto
	}
	if err := rows.Err(); err != nil {
		return err
	}

	// Collect the ID of the user who want to unlike
	rows1, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, user)
	if err != nil {
		return err
	}

	isNotEmpty := rows1.Next()
	if err := rows.Err(); err != nil {
		return err
	}

	if isNotEmpty {
		err = rows1.Scan(&userId)
		if err != nil {
			return err
		}
	} else {
		return ErrUserDoesntExist
	}

	/*
		err = db.CheckBan(userId, photoOwner)
		if err != nil {
			return err
		}
	*/

	rows.Close()  // -> Required, otherwise the database remains locked
	rows1.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`DELETE FROM Like WHERE FotoReference = ? AND Utente = ?;`, photoId, userId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE Foto SET nLikes = nLikes - 1 WHERE FotoID = ? AND Utente = ?;`, photoId, userId)
	if err != nil {
		return err
	}

	return nil
}
