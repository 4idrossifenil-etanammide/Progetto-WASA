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
	var tmpId string
	for rows.Next() {
		err = rows.Scan(&tmp.PhotoID, &tmp.UploadingDate, &tmp.LikeNumber, &tmp.CommentNumber)
		if err != nil {
			return Stream{}, err
		}

		rows1, err := db.c.Query(`SELECT ID, Nome From Utente WHERE Utente.ID IN (SELECT Utente FROM Foto WHERE FotoID = ?);`, tmp.PhotoID)
		if err != nil {
			return Stream{}, err
		}

		for rows1.Next() {
			err = rows1.Scan(&tmpId, &tmp.Name)
			if err != nil {
				return Stream{}, err
			}
		}
		if err := rows1.Err(); err != nil {
			return Stream{}, err
		}
		rows1.Close()

		tmp.Comments, err = getComments(userId, tmp.PhotoID, db)
		if err != nil {
			return Stream{}, err
		}
		tmp.CommentNumber = len(tmp.Comments)

		toSubtract, err := filterLike(tmp.PhotoID, userId, db)
		if err != nil {
			return Stream{}, err
		}

		tmp.LikeNumber = tmp.LikeNumber - toSubtract

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

func filterLike(photoID string, id string, db *appdbimpl) (int, error) {
	rows, err := db.c.Query(`SELECT Utente FROM Like WHERE FotoReference = ?;`, photoID)
	if err != nil {
		return 0, err
	}

	var tmpId string
	bannedLikeCounter := 0
	for rows.Next() {
		err = rows.Scan(&tmpId)
		if err != nil {
			return 0, err
		}

		if db.CheckBan(tmpId, id) != nil || db.CheckBan(id, tmpId) != nil {
			bannedLikeCounter = bannedLikeCounter + 1
		}
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}
	rows.Close()

	return bannedLikeCounter, nil
}

func getComments(userId string, photoID string, db *appdbimpl) ([]Comment, error) {
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

		if db.CheckBan(id, userId) == nil && db.CheckBan(userId, id) == nil {
			comment.ID = commentId
			comment.Name = name
			comment.Text = text

			toReturn = append(toReturn, comment)
		}
	}

	if err := rows.Err(); err != nil {
		return []Comment{}, err
	}
	rows.Close()
	return toReturn, nil
}
