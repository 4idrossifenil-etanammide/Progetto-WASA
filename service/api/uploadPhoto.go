package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
)

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	username := ps.ByName("user_name")
	var id ID
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		w.WriteHeader(http.StatusBadRequest)
	}

	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		_ = os.Mkdir("images", os.ModePerm)
	}

	var photo Photo
	dbPhoto, err := rt.db.UploadPhoto(id.Id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Errore nell'inserimento nel database")
		w.WriteHeader(http.StatusBadRequest)
	}
	photo.FromDatabase(dbPhoto)
	file, _ := os.Create("./images/" + photo.PhotoID + ".jpg")
	_, err = io.Copy(file, r.Body)

	if err != nil {
		ctx.Logger.WithError(err).Error("Errore nella lettura dell'immagine")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
