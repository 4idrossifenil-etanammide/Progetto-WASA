package api

import "wasaphoto.uniroma1.it/wasaphoto/service/database"

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
	/*
		TODO: Implementare il check sulla regex implmentata nel file api.yaml
			  Per ora ancora non sono riuscito a tradurre la regex dallo standard di swaggere a quello del Go
	*/
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
	PhotoID       string `json:"photoID"`
	Date          string `json:"date"`
	LikeNumber    int    `json:"likeNumber"`
	CommentNumber int    `json:"commentNumber"`
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.PhotoID = photo.PhotoID
	p.Date = photo.Date
	p.CommentNumber = photo.CommentNumber
	p.LikeNumber = photo.LikeNumber
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		PhotoID:       p.PhotoID,
		Date:          p.Date,
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
