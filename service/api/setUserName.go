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

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Initialize variables
	var id ID
	var newName UserName

	// Take the name of the user from the body
	err := json.NewDecoder(r.Body).Decode(&newName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Take the name of the user from the path
	username := ps.ByName("user_name")

	// Check the authorization
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err = rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		if errors.Is(err, database.ErrAuthentication) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// If the new name is equal to the name in the path, we don't have to change anything because of idempotency
	// (The user is trying to put his old name as the new name)
	if username == newName.Name {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Finally, it enters the change into the database, checking for errors
	err = rt.db.SetName(id.Id, newName.Name)
	if errors.Is(err, database.ErrChangeName) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("can't change the name of the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
