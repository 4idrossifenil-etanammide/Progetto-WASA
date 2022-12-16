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
	rt.router.DELETE("/profiles/:user_name/banned/:banned_user", rt.wrap(rt.UnbanUser))

	rt.router.PUT("/profiles/:user_name/photos/:photo_id/likes/:user", rt.wrap(rt.LikePhoto))
	rt.router.DELETE("/profiles/:user_name/photos/:photo_id/likes/:user", rt.wrap(rt.UnlikePhoto))

	rt.router.POST("/profiles/:user_name/photos/:photo_id/comments", rt.wrap(rt.CommentPhoto))
	rt.router.DELETE("/profiles/:user_name/photos/:photo_id/comments/:user", rt.wrap(rt.DeleteComment))

	rt.router.PUT("/profiles/:user_name/followers/:user_following", rt.wrap(rt.Follow))
	rt.router.DELETE("/profiles/:user_name/followers/:user_following", rt.wrap(rt.Unfollow))

	rt.router.GET("/profiles/:user_name/photos", rt.wrap(rt.GetStream))
	rt.router.GET("/profiles/:user_name", rt.wrap(rt.GetProfile))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
