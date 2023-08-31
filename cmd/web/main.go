package main

import (
	"encoding/gob"
	"fmt"
	"github.com/RazvanING/bookings/internal/config"
	"github.com/RazvanING/bookings/internal/handlers"
	"github.com/RazvanING/bookings/internal/models"
	"github.com/RazvanING/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8082"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function

func main() {
	// what am i going to put in the session

	gob.Register(models.Reservation{})

	//change this to true when in production

	app.InProduction = false

	session = scs.New()
	//how long a session should last, shorter for more secure applications
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
