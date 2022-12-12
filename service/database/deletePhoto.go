package database

import (
	"errors"
)

func (db *appdbimpl) DeletePhoto(photoId string, userId string) error {

	rows, err := db.c.Query(`SELECT Utente FROM Foto WHERE FotoID = ?`, photoId)

	var output string

	if rows.Next() {
		rows.Scan(&output)

		if output != userId {
			return errors.New("Permission denied")
		}
	}
	rows.Close()

	_, err = db.c.Exec(`DELETE FROM Foto WHERE FotoID = ? AND Utente = ?;`, photoId, userId)

	if err != nil {
		return err
	}

	return nil
}
