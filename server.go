package main

import (
	"log"
	"net/http"

	"./routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	//"github.com/joho/godotenv"
)

func main() {
	router := CreateRoutes()

	walkF := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s, %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkF); err != nil {
		log.Fatalf("Logging Error: %s",err.Error())
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateRoutes() *chi.Mux{
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	router.Route("/", func(r chi.Router){
		r.Mount("/quad-key", routes.Routes())
	})
	return router
}