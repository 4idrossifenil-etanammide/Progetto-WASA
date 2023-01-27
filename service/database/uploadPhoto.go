package database

import (
	"hash/fnv"
	"strconv"
	"time"
)

func (db *appdbimpl) UploadPhoto(userId string, path string, username string) (Photo, error) {

	// Initialize variable
	var photo Photo

	// Populate the struct with the information about the foto
	CreateData(&photo, username)

	// Put the information into the database
	_, err := db.c.Exec(`INSERT INTO Foto (FotoID, Utente, Data, nLikes, nCommenti, Path) VALUES (?, ?, ?, ?, ?, ?)`, photo.PhotoID, userId, photo.UploadingDate, photo.LikeNumber, photo.CommentNumber, path+photo.PhotoID+".jpg")
	if err != nil {
		return Photo{}, err
	}
	return photo, nil
}

// This function create the info about the photo
func CreateData(photoStruct *Photo, username string) {

	// Creation of the unique ID and insertion of the date
	h := fnv.New32a()
	photoStruct.UploadingDate = time.Now()
	h.Write([]byte(photoStruct.UploadingDate.Format("02-01-2006 15:04:05")))

	// Put the info into the struct
	photoStruct.PhotoID = strconv.Itoa(int(h.Sum32()))
	photoStruct.CommentNumber = 0
	photoStruct.LikeNumber = 0

	photoStruct.Name = username
}
