package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/database"
)

func (rt *_router) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var id ID

	userComment := ps.ByName("user")
	photoId := ps.ByName("photo_id")
	commentId := ps.ByName("comment_id")

	// Check authorization, in this case we have to check the authorization for the
	// user who commented
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, userComment)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		if errors.Is(err, database.ErrAuthentication) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = rt.db.DeleteComment(photoId, commentId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
