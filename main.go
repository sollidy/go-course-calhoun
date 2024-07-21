package main

import (
	"fmt"
	"main/controllers"
	"main/templates"
	"main/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	t := views.Must(views.ParseFs(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(t))

	t = views.Must(views.ParseFs(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(t))

	t = views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(t))

	t = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", controllers.FAQ(t))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3001 ...")
	http.ListenAndServe(":3001", r)
}
