package routes

import (
	"github.com/go-chi/chi"

	"../controllers"
//	"../middleware"
)

func Routes() *chi.Mux{
	router:= chi.NewRouter()
	router.Get("/{coord}", controllers.GetTile)
	return router
}
