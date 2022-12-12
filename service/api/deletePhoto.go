package api

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) DeletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("user_name")
	var id ID
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		w.WriteHeader(http.StatusBadRequest)
	}

	photoId := ps.ByName("photo_id")

	err = rt.db.DeletePhoto(photoId, id.Id)

	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
	}
}
