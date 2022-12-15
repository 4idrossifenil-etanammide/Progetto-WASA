package database

func (db *appdbimpl) CommentPhoto(photoId string, comment Comment) (UserName, error) {

	rows, err := db.c.Query(`SELECT * FROM Foto WHERE FotoID = ?`, photoId)
	if err != nil {
		return UserName{}, err
	}
	if !rows.Next() {
		return UserName{}, ErrPhoto
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, comment.Name)
	if err != nil {
		return UserName{}, err
	}

	var userId string
	if rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return UserName{}, err
		}
	}

	rows.Close()
	_, err = db.c.Exec(`INSERT INTO Commenti(FotoReference, Utente, Testo) VALUES (?, ?, ?);`, photoId, userId, comment.Text)
	if err != nil {
		return UserName{}, err
	}

	return UserName{Name: comment.Name}, nil
}
