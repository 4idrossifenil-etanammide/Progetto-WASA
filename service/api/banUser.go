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

func (rt *_router) BanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Initialize variable
	var body UserName
	var id ID

	// Collect information from the parameters
	banned_user := ps.ByName("banned_user")
	username := ps.ByName("user_name")

	//Check Authorization
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		if errors.Is(err, database.AuthenticationError) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// Collect information from the body
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the user is trying to ban himself and if the user to ban specified in the body is equal to the one in the parameter
	if banned_user != body.Name || banned_user == username {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Write the information on the database
	err = rt.db.BanUser(id.Id, banned_user)

	// Check for errors from database
	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
