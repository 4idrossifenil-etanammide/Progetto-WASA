package api

import (
	"time"

	"wasaphoto.uniroma1.it/wasaphoto/service/database"
)

// Struct for collecting information about a username
type UserName struct {
	Name string `json:"name"`
}

// Convert the struct from the database to the one defined here (these are equal but from different package)
func (u *UserName) FromDatabase(user database.UserName) {
	u.Name = user.Name
}

// Convert this struct to the one defined (these are equal but from different package) in the database.go file
func (u *UserName) ToDatabase() database.UserName {
	return database.UserName{
		Name: u.Name,
	}
}

// Thi method check if the username is valid
func (u *UserName) ValidUserName() bool {
	return len(u.Name) >= 3 && len(u.Name) <= 16
}

// Same as above, but with the struct for the Id
type ID struct {
	Id string `json:"id"`
}

func (i *ID) FromDatabase(id database.ID) {
	i.Id = id.Id
}

func (i *ID) ToDatabase() database.ID {
	return database.ID{
		Id: i.Id,
	}
}

// Struct to collect information about a photo, same as above
type Photo struct {
	Name          string    `json:"name"`
	PhotoID       string    `json:"photoID"`
	UploadingDate time.Time `json:"date"`
	LikeNumber    int       `json:"likeNumber"`
	CommentNumber int       `json:"commentNumber"`
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.Name = photo.Name
	p.PhotoID = photo.PhotoID
	p.UploadingDate = photo.UploadingDate
	p.CommentNumber = photo.CommentNumber
	p.LikeNumber = photo.LikeNumber
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		Name:          p.Name,
		PhotoID:       p.PhotoID,
		UploadingDate: p.UploadingDate,
		LikeNumber:    p.LikeNumber,
		CommentNumber: p.CommentNumber,
	}
}

type Comment struct {
	Name string `json:"name"`
	Text string `json:"comment"`
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		Name: c.Name,
		Text: c.Text,
	}
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.Name = comment.Name
	c.Text = comment.Text
}

type Stream struct {
	Photos []Photo `json:"photos"`
}

func (s *Stream) FromDatabase(stream database.Stream) {
	var tmpPhoto Photo
	for _, photo := range stream.Photos {
		tmpPhoto.FromDatabase(photo)
		s.Photos = append(s.Photos, tmpPhoto)
	}
}

type Profile struct {
	Photos      []Photo
	PhotoNumber int
	Follower    []UserName
	Following   []UserName
}

func (p *Profile) FromDatabase(profile database.Profile) {
	var tmpPhoto Photo
	for _, photo := range profile.Photos {
		tmpPhoto.FromDatabase(photo)
		p.Photos = append(p.Photos, tmpPhoto)
	}

	p.PhotoNumber = profile.PhotoNumber

	var tmpFollower UserName
	for _, username := range profile.Follower {
		tmpFollower.FromDatabase(username)
		p.Follower = append(p.Follower, tmpFollower)
	}

	var tmpFollowing UserName
	for _, username := range profile.Following {
		tmpFollowing.FromDatabase(username)
		p.Following = append(p.Following, tmpFollowing)
	}
}
