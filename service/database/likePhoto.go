package database

import "errors"

var LikePhotoError = errors.New("The photo doesn't exist")
var UserDoesntExistError = errors.New("The user doesn't exist")

func (db *appdbimpl) LikePhoto(photoId string, user string) error {

	// Initialize variable
	var photoOwner string
	var userId string

	// Get the info from the database
	rows, err := db.c.Query(`SELECT Utente FROM Foto WHERE FotoID = ?`, photoId)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	// Check and collect info about the owner of the photo
	if rows.Next() {
		err = rows.Scan(&photoOwner)
		if err != nil {
			return err
		}
	} else {
		return LikePhotoError
	}

	// Collect the ID of the user who want to like
	rows1, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, user)
	if err != nil {
		return err
	}
	defer func() { _ = rows1.Close() }()

	if rows1.Next() {
		err = rows1.Scan(&userId)
		if err != nil {
			return err
		}
	} else {
		return UserDoesntExistError
	}

	/*
		err = db.CheckBan(userId, photoOwner)
		if err != nil {
			return err
		}
	*/

	rows.Close()  // -> Required, otherwise the database remains locked
	rows1.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`INSERT INTO Like(FotoReference, Utente) VALUES (?, ?);`, photoId, userId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec(`UPDATE Foto SET nLikes = nLikes + 1 WHERE FotoID = ? AND Utente = ?;`, photoId, userId)
	if err != nil {
		return err
	}

	return nil
}
