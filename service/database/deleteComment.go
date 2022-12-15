package database

import (
	"errors"
)

var ErrCommentDeleteNotAllowed = errors.New("You can delete only your own comments")

func (db *appdbimpl) DeleteComment(photoId string, userComment string) error {

	rows, err := db.c.Query(`SELECT Utente FROM Foto WHERE FotoID = ?`, photoId)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return ErrPhoto
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, userComment)
	if err != nil {
		return err
	}

	var userCommentId string
	if rows.Next() {
		err = rows.Scan(&userCommentId)
		if err != nil {
			return err
		}
	}

	rows.Close()
	_, err = db.c.Exec(`DELETE FROM Commenti WHERE Utente = ?`, userCommentId)
	if err != nil {
		return err
	}

	return nil
}
