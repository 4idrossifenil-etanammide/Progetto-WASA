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

		stream.Photos = append(stream.Photos, tmp)
	}
	if err := rows.Err(); err != nil {
		return Stream{}, err
	}

	sort.Slice(stream.Photos, func(i, j int) bool {
		return stream.Photos[i].UploadingDate.After(stream.Photos[j].UploadingDate)
	})

	return stream, nil
}
