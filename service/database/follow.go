package database

func (db *appdbimpl) Follow(userId string, follower string) error {

	rows, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?`, follower)
	if err != nil {
		return err
	}

	// Put the ID in a string
	var followerId string
	if rows.Next() {
		err = rows.Scan(&followerId)
		if err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT * FROM Segue WHERE Utente = ? AND Follower = ?`, userId, followerId)
	if err != nil {
		return err
	}
	if rows.Next() {
		return nil
	}
	if err := rows.Err(); err != nil {
		return err
	}

	// Check if the user is not banned
	err = db.CheckBan(userId, followerId)
	if err != nil {
		return err
	}

	rows.Close() // -> Required, otherwise the database remains locked
	_, err = db.c.Exec(`INSERT INTO Segue(Utente, Follower) VALUES (?,?);`, userId, followerId)
	if err != nil {
		return err
	}

	return nil
}
