package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"wasaphoto.uniroma1.it/wasaphoto/service/api/reqcontext"
	"wasaphoto.uniroma1.it/wasaphoto/service/database"
)

func (rt *_router) UploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Initialize the variable
	var id ID
	var photo Photo

	// Collect info from the parameters
	username := ps.ByName("user_name")

	// Check Authorization
	id.Id = strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := rt.db.CheckToken(id.Id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Authorization failed!")
		if errors.Is(err, database.ErrAuthentication) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// If the folder where we want to insert the photos does not exist, we create it
	if _, err := os.Stat("/tmp/images"); os.IsNotExist(err) {
		err = os.Mkdir("/tmp/images", os.ModePerm)
		if err != nil {
			ctx.Logger.WithError(err).Error("Unable to create the folder /images/")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Unable to create the folder /images/")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the info into the database, using the id of the user and the path to the folder where we store the images
	dbPhoto, err := rt.db.UploadPhoto(id.Id, "/tmp/images/", username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when inserting on database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the photo information from the database struct
	photo.FromDatabase(dbPhoto)

	// Put the image into the folder
	file, err := os.Create("/tmp/images/" + photo.PhotoID + ".png")
	if err != nil {
		ctx.Logger.WithError(err).Error("Errore nella lettura dell'immagine")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(file, r.Body)
	if err != nil {
		ctx.Logger.WithError(err).Error("Errore nella lettura dell'immagine")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if everything is ok, we send the result
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photo)
}
