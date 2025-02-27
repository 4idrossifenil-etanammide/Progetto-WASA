package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/database"
)

func (rt *_router) CommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id ID
	var toReturn CommentID
	var comment Comment

	photoId := ps.ByName("photo_id")

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check authorization, in this case we check the that the user who want to comment have
	// the right to do so
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err = rt.db.CheckToken(id.Id, comment.Name)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		if errors.Is(err, database.ErrAuthentication) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	result, err := rt.db.CommentPhoto(photoId, comment.ToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	toReturn.FromDatabase(result)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(toReturn)

}

type CommentID struct {
	ID int `json:"id"`
}

func (c *CommentID) FromDatabase(comment database.CommentID) {
	c.ID = comment.ID
}
