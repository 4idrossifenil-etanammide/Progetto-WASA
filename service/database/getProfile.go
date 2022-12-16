package database

func (db *appdbimpl) GetProfile(userId string, profileName string) (Profile, error) {

	var profile Profile
	var profileId string

	rows, err := db.c.Query(`SELECT ID FROM Utente WHERE Nome = ?;`, profileName)
	if err != nil {
		return Profile{}, err
	}
	if rows.Next() {
		err = rows.Scan(&profileId)
		if err != nil {
			return Profile{}, err
		}
	} else {
		return Profile{}, ErrUserDoesntExist
	}

	err = db.CheckBan(userId, profileId)
	if err != nil {
		return Profile{}, err
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT FotoID, Data, nLikes, nCommenti FROM Foto WHERE Utente = ?;`, profileId)
	if err != nil {
		return Profile{}, err
	}

	var tmpPhoto Photo
	for rows.Next() {
		err = rows.Scan(&tmpPhoto.PhotoID, &tmpPhoto.UploadingDate, &tmpPhoto.LikeNumber, &tmpPhoto.CommentNumber)
		if err != nil {
			return Profile{}, err
		}

		profile.Photos = append(profile.Photos, tmpPhoto)
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT Nome FROM Utente WHERE ID IN (SELECT Follower FROM Segue WHERE Utente = ?);`, profileId)
	if err != nil {
		return Profile{}, err
	}

	var tmpFollower UserName
	for rows.Next() {
		err = rows.Scan(&tmpFollower.Name)
		if err != nil {
			return Profile{}, err
		}

		profile.Follower = append(profile.Follower, tmpFollower)
	}

	rows.Close()
	rows, err = db.c.Query(`SELECT Nome FROM Utente WHERE ID IN (SELECT Utente FROM Segue WHERE Follower = ?);`, profileId)
	if err != nil {
		return Profile{}, err
	}

	var tmpFollowing UserName
	for rows.Next() {
		err = rows.Scan(&tmpFollowing.Name)
		if err != nil {
			return Profile{}, err
		}

		profile.Following = append(profile.Following, tmpFollowing)
	}
	rows.Close()

	return profile, nil
}
