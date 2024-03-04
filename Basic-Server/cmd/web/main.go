package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/Heinirich/basic-server/pkg/config"
	"github.com/Heinirich/basic-server/pkg/handlers"
	"github.com/Heinirich/basic-server/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const SERVERPORT = ":8080"

var session *scs.SessionManager

var app config.AppConfig

func main() {

	// change to true when in production
	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * 7 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		fmt.Println(err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	

	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
	// 	http.NotFound(w, r)
	// })
	// http.HandleFunc("/divide", Divide)

	fmt.Println("Server is running on port", SERVERPORT)
	// err = http.ListenAndServe(SERVERPORT, nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }
	srv := &http.Server{
		Addr:    SERVERPORT,
		Handler: routes(&app),
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
