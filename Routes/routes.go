package Routes

import (
	"advertisement-api/Controller"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func HandleRequests() {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Route("/company", func(r chi.Router) {
			r.Get("/viewallcompany", Controller.ViewAllCompany)
			r.Post("/createcompany", Controller.CreateCompany)
			//r.Get("/", Controller.ViewAllSembako)
			//r.Put("/", Controller.EditSembako)
			//r.Delete("/", Controller.DeleteSembako)
		})
		r.Route("/ads", func(r chi.Router) {
			r.Get("/viewalladvertisement", Controller.ViewAllAdvertisement)
			r.Post("/placeadvertisement", Controller.PlaceAdvertisement)
			//r.Get("/", Controller.ViewAllSembako)
			//r.Put("/", Controller.EditSembako)
			//r.Delete("/", Controller.DeleteSembako)
		})
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("API Connection Error during attaching http handler")
		return
	}
	fmt.Println("Server Connected to localhost:3000")
}
