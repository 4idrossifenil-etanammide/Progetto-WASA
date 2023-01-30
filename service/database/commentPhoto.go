package database

import "strconv"

func (db *appdbimpl) CommentPhoto(photoId string, comment Comment) (CommentID, error) {

	rows, err := db.c.Query(`SELECT * FROM Foto WHERE FotoID = ?;`, photoId)
	if err != nil {
		return CommentID{}, err
	}
	if !rows.Next() {
		return CommentID{}, ErrPhoto
	}
	if err := rows.Err(); err != nil {
		return CommentID{}, err
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?;`, comment.Name)
	if err != nil {
		return CommentID{}, err
	}

	var userId string
	if rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return CommentID{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return CommentID{}, err
	}

	rows.Close()
	_, err = db.c.Exec(`INSERT INTO Commenti(FotoReference, Utente, Testo) VALUES (?, ?, ?);`, photoId, userId, comment.Text)
	if err != nil {
		return CommentID{}, err
	}

	_, err = db.c.Exec(`UPDATE Foto SET nCommenti = nCommenti + 1 WHERE FotoID = ?;`, photoId)
	if err != nil {
		return CommentID{}, err
	}

	rows, err = db.c.Query(`SELECT CommentiID FROM Commenti WHERE FotoReference = ? AND Utente = ? AND Testo = ?;`, photoId, userId, comment.Text)
	if err != nil {
		return CommentID{}, err
	}

	var commentId string
	if rows.Next() {
		err = rows.Scan(&commentId)
		if err != nil {
			return CommentID{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return CommentID{}, err
	}

	rows.Close()

	var valueToReturn int
	valueToReturn, err = strconv.Atoi(commentId)
	if err != nil {
		return CommentID{}, err
	}

	return CommentID{ID: valueToReturn}, nil
}
