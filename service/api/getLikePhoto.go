package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) GetLikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Collect information from the parameters
	photoId := ps.ByName("photo_id")
	userLike := ps.ByName("user")

	isAlreadyLiked, err := rt.db.GetLikePhoto(photoId, userLike)
	if err != nil {
		ctx.Logger.WithError(err).Error("Operation failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var like Like
	if isAlreadyLiked {
		like.Liked = "true"
	} else {
		like.Liked = "false"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(like)

}

type Like struct {
	Liked string `json:"liked"`
}
