package database

import "errors"

var ErrPhoto = errors.New("The photo doesn't exist")
var ErrUserDoesntExist = errors.New("The user doesn't exist")

func (db *appdbimpl) LikePhoto(photoId string, user string) error {

	// Initialize variable
	var photoOwner string
	var userId string

	// Get the info from the database
	rows, err := db.c.Query(`SELECT Utente FROM Foto WHERE FotoID = ?`, photoId)
	if err != nil {
		return err
	}

	// Check and collect info about the owner of the photo
	isNotEmpty := rows.Next()
	if err := rows.Err(); err != nil {
		return err
	}

	if isNotEmpty {
		err = rows.Scan(&photoOwner)
		if err != nil {
			return err
		}
	} else {
		rows.Close()
		return ErrPhoto
	}
	rows.Close() // -> Required, otherwise the database remains locked

	// Collect the ID of the user who want to like
	rows, err = db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, user)
	if err != nil {
		return err
	}

	isNotEmpty = rows.Next()
	if err := rows.Err(); err != nil {
		return err
	}

	if isNotEmpty {
		err = rows.Scan(&userId)
		if err != nil {
			return err
		}
	} else {
		rows.Close()
		return ErrUserDoesntExist
	}

	rows.Close() // -> Required, otherwise the database remains locked
	rows, err = db.c.Query(`SELECT * FROM Like WHERE FotoReference = ? AND Utente = ?`, photoId, userId)
	if err != nil {
		return err
	}
	if rows.Next() {
		rows.Close()
		return nil
	}
	if err := rows.Err(); err != nil {
		return err
	}

	rows.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`INSERT INTO Like(FotoReference, Utente) VALUES (?, ?);`, photoId, userId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE Foto SET nLikes = nLikes + 1 WHERE FotoID = ?;`, photoId)
	if err != nil {
		return err
	}

	return nil
}
