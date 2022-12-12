package database

import "errors"

func (db *appdbimpl) LikePhoto(photoId string, user string) error {

	rows, _ := db.c.Query(`SELECT Utente FROM Foto WHERE FotoID = ?`, photoId)

	var photoOwner string

	if rows.Next() {
		rows.Scan(&photoOwner)
	} else {
		return errors.New("The photo doesn't exist")
	}
	rows.Close()

	rows, _ = db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, user)

	var userId string

	if rows.Next() {
		rows.Scan(&userId)
	}
	rows.Close()

	err := db.CheckBan(userId, photoOwner)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`INSERT INTO Like(FotoReference, Utente) VALUES (?, ?);`, photoId, userId)
	if err != nil {
		return err
	}

	return nil
}
