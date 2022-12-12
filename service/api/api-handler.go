package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.DoLogin))
	rt.router.PUT("/profiles/:user_name/name", rt.wrap(rt.setMyUserName))
	rt.router.POST("/profiles/:user_name/photos", rt.wrap(rt.UploadPhoto))
	rt.router.DELETE("/profiles/:user_name/photos/:photo_id", rt.wrap(rt.DeletePhoto))
	rt.router.PUT("/profiles/:user_name/banned/:banned_user", rt.wrap(rt.BanUser))

	rt.router.PUT("/profiles/:user_name/photos/:photo_id/likes/:user", rt.wrap(rt.LikePhoto))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
