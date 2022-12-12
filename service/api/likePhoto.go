package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) LikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	username := ps.ByName("user_name")
	var id ID
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		w.WriteHeader(http.StatusBadRequest)
	}

	var user UserName
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	photoId := ps.ByName("photo_id")
	userLike := ps.ByName("user")

	if user.Name != userLike {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.LikePhoto(photoId, userLike)

	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
	}

}
