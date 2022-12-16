package database

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

		stream.Photos = append(stream.Photos, tmp)
	}

	return stream, nil
}
