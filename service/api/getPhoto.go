package api

import (
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) GetPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	http.ServeFile(w, r, filepath.Join("/tmp/images", ps.ByName("photo_id")+".png"))
}
