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

func (rt *_router) GetProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id ID
	var profile Profile

	username := ps.ByName("user_name")

	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]

	result, err := rt.db.GetProfile(id.Id, username)
	if err != nil {
		if errors.Is(err, database.ErrBan) {
			w.WriteHeader(http.StatusForbidden)
			return
		} else if errors.Is(err, database.ErrUserDoesntExist) {
			ctx.Logger.WithError(err).Error("Operation failed!")
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			ctx.Logger.WithError(err).Error("Operation failed!")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	profile.FromDatabase(result)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
