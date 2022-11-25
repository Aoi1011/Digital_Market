package main

import (
	"fmt"
	"net/http"

	"github.com/Aoi1011/lenslocked/controllers"
	"github.com/Aoi1011/lenslocked/models"
	"github.com/Aoi1011/lenslocked/templates"
	"github.com/Aoi1011/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	userC := controllers.Users{
		UserService: &userService,
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(
		views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	userC.Templates.New = views.Must(
		views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	userC.Templates.SignIn = views.Must(
		views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", userC.New)
	r.Post("/signup", userC.Create)
	r.Get("/signin", userC.SignIn)
	r.Post("/signin", userC.ProcessSignIn)
	r.Get("/users/me", userC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false),
	)
	http.ListenAndServe(":3000", csrfMw(r))
}
