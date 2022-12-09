package database

import (
	"hash/fnv"
	"strconv"
	"time"
)

func (db *appdbimpl) UploadPhoto(userId string, path string) (Photo, error) {
	var photo Photo
	CreateData(&photo)

	_, err := db.c.Exec(`INSERT INTO Foto (FotoID, Utente, Data, nLikes, nCommenti, Path) VALUES (?, ?, ?, ?, ?, ?)`, photo.PhotoID, userId, photo.Date, photo.LikeNumber, photo.CommentNumber, path+photo.PhotoID+".jpg")
	return photo, err
}

func CreateData(photoStruct *Photo) {
	h := fnv.New32a()
	photoStruct.Date = time.Now().String()
	h.Write([]byte(photoStruct.Date))
	photoStruct.PhotoID = strconv.Itoa(int(h.Sum32()))
	photoStruct.CommentNumber = 0
	photoStruct.LikeNumber = 0
}
