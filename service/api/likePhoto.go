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

func (rt *_router) LikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Initialize variable
	var id ID
	var user UserName

	// Collect information from the parameters
	photoId := ps.ByName("photo_id")
	userLike := ps.ByName("user")

	// Check authorization, in this case we check if the person who want to like
	// have the right to do so
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, userLike)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		if errors.Is(err, database.ErrAuthentication) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Collect the name of the user from the body
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If the name specified in the body is not equal to the one specified in the parameters, we return an error
	if user.Name != userLike {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Put the information into the database
	err = rt.db.LikePhoto(photoId, userLike)
	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
