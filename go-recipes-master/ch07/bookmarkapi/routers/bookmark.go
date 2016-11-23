package routers

import (
	"github.com/gorilla/mux"

	"github.com/shijuvar/go-recipes/ch07/bookmarkapi/controllers"
)

// SetBookmarkRoutes registers routes for bookmark entity.
func SetBookmarkRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/bookmarks", controllers.CreateBookmark).Methods("POST")
	router.HandleFunc("/bookmarks/{id}", controllers.UpdateBookmark).Methods("PUT")
	router.HandleFunc("/bookmarks", controllers.GetBookmarks).Methods("GET")
	router.HandleFunc("/bookmarks/{id}", controllers.GetBookmarkByID).Methods("GET")
	router.HandleFunc("/bookmarks/users/{id}", controllers.GetBookmarksByUser).Methods("GET")
	router.HandleFunc("/bookmarks/{id}", controllers.DeleteBookmark).Methods("DELETE")
	return router
}
