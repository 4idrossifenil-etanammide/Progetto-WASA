package database

import (
	"sort"
)

func (db *appdbimpl) GetStream(userId string) (Stream, error) {

	var stream Stream

	rows, err := db.c.Query(`SELECT FotoID, Data, nLikes, nCommenti FROM Foto WHERE Foto.Utente IN (SELECT Follower FROM Segue WHERE Utente = ?);`, userId)
	if err != nil {
		return Stream{}, err
	}

	var tmp Photo
	for rows.Next() {
		err = rows.Scan(&tmp.PhotoID, &tmp.UploadingDate, &tmp.LikeNumber, &tmp.CommentNumber)
		if err != nil {
			return Stream{}, err
		}

		rows1, err := db.c.Query(`SELECT Nome From Utente WHERE Utente.ID IN (SELECT Utente FROM Foto WHERE FotoID = ?);`, tmp.PhotoID)
		if err != nil {
			return Stream{}, err
		}

		for rows1.Next() {
			err = rows1.Scan(&tmp.Name)
			if err != nil {
				return Stream{}, err
			}
		}
		if err := rows1.Err(); err != nil {
			return Stream{}, err
		}
		rows1.Close()

		tmp.Comments, err = getComments(tmp.PhotoID, db)
		if err != nil {
			return Stream{}, err
		}

		stream.Photos = append(stream.Photos, tmp)
		tmp.Comments = []Comment{}
	}
	if err := rows.Err(); err != nil {
		return Stream{}, err
	}

	rows.Close()
	sort.Slice(stream.Photos, func(i, j int) bool {
		return stream.Photos[i].UploadingDate.After(stream.Photos[j].UploadingDate)
	})

	return stream, nil
}

func getComments(photoID string, db *appdbimpl) ([]Comment, error) {
	rows, err := db.c.Query(`SELECT CommentiID, Utente, Testo FROM Commenti WHERE FotoReference = ?;`, photoID)
	if err != nil {
		return []Comment{}, err
	}

	var commentId int
	var id string
	var text string
	var name string
	var comment Comment
	var toReturn []Comment
	for rows.Next() {
		err = rows.Scan(&commentId, &id, &text)
		if err != nil {
			return []Comment{}, err
		}

		rows1, err := db.c.Query(`SELECT Nome FROM Utente WHERE ID = ?;`, id)
		if err != nil {
			return []Comment{}, err
		}

		for rows1.Next() {
			err = rows1.Scan(&name)
			if err != nil {
				return []Comment{}, err
			}
		}
		if err := rows1.Err(); err != nil {
			return []Comment{}, err
		}
		rows1.Close()

		comment.ID = commentId
		comment.Name = name
		comment.Text = text

		toReturn = append(toReturn, comment)
	}

	if err := rows.Err(); err != nil {
		return []Comment{}, err
	}
	rows.Close()
	return toReturn, nil
}
