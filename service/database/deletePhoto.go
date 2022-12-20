package database

import (
	"errors"
)

var ErrDeletePhoto = errors.New("You can't delete the photo of another user")

func (db *appdbimpl) DeletePhoto(photoId string, userId string) error {

	// Collect information from the database
	rows, err := db.c.Query(`SELECT Utente FROM Foto WHERE FotoID = ?`, photoId)
	if err != nil {
		return err
	}

	// check that the user who wants to delete the photo is the owner
	var photoOwner string
	if rows.Next() {
		err = rows.Scan(&photoOwner)
		if err != nil {
			return err
		}

		if photoOwner != userId {
			return ErrDeletePhoto
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	// Delete the photo
	rows.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`DELETE FROM Foto WHERE FotoID = ? AND Utente = ?;`, photoId, userId)
	if err != nil {
		return err
	}

	return nil
}
