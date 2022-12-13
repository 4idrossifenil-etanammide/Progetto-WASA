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
		if errors.Is(err, database.AuthenticationError) {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// If the folder where we want to insert the photos does not exist, we create it
	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		_ = os.Mkdir("images", os.ModePerm)
	} else if err != nil {
		ctx.Logger.WithError(err).Error("Unable to create the folder /images/")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Insert the info into the database, using the id of the user and the path to the folder where we store the images
	dbPhoto, err := rt.db.UploadPhoto(id.Id, "./images/")
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when inserting on database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Convert the photo information from the database struct
	photo.FromDatabase(dbPhoto)

	// Put the image into the folder
	file, _ := os.Create("./images/" + photo.PhotoID + ".jpg")
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
