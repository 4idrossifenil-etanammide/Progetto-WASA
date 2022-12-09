package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var user UserName
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := ps.ByName("user_name")
	var id ID
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err = rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		w.WriteHeader(http.StatusBadRequest)
	}

	err = rt.db.SetName(id.Id, user.Name)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't change the name of the user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
