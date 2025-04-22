package handlers 

import(
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/AdamNgazzou/go_RESTFUL_APIs/internal/middleware"
)

func Handler(r *chi.Mux){ // it starts with capitalized letter = public function to be imported
	// Global middleware 
	r.Use(chimiddle.StripSlashes) // for example from  http:localhost:8000/account/coins/ => http:localhost:8000/account/coins

	r.Route("/account",func(router chi.Router){

		//Middleware for /account route 
		router.Use(middleware.Authorization)
		
		router.Get("/coins",GetCoinBalance)
	})

}