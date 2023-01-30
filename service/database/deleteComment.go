package database

import "strconv"

func (db *appdbimpl) DeleteComment(photoId string, commentId string) error {

	var valueForQuery int
	valueForQuery, err := strconv.Atoi(commentId)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`DELETE FROM Commenti WHERE CommentiID = ?;`, valueForQuery)
	if err != nil {
		return err
	}

	_, err = db.c.Exec(`UPDATE Foto SET nCommenti = nCommenti - 1 WHERE FotoID = ?;`, photoId)
	if err != nil {
		return err
	}

	return nil
}
