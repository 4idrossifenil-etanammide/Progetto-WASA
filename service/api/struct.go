package api

import "wasaphoto.uniroma1.it/wasaphoto/service/database"

type UserName struct {
	Name string `json:"name"`
}

func (u *UserName) FromDatabase(user database.UserName) {
	u.Name = user.Name
}

func (u *UserName) ToDatabase() database.UserName {
	return database.UserName{
		Name: u.Name,
	}
}

func (u *UserName) ValidUserName() bool {
	return len(u.Name) >= 3 && len(u.Name) <= 16
	/*
		TODO: Implementare il check sulla regex implmentata nel file api.yaml
			  Per ora ancora non sono riuscito a tradurre la regex dallo standard di swaggere a quello del Go
	*/
}

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
