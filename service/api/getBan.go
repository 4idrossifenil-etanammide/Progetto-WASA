package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) GetBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	banned_user := ps.ByName("banned_user")
	username := ps.ByName("user_name")

	isAlreadyBanned, err := rt.db.GetBan(username, banned_user)
	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ban Ban
	if isAlreadyBanned {
		ban.Banned = "true"
	} else {
		ban.Banned = "false"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ban)

}

type Ban struct {
	Banned string `json:"banned"`
}
