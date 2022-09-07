package main

import (
	"database/sql"
	"encoding/gob"
	"github.com/DeLuci/inventory-system/internal/config"
	"github.com/DeLuci/inventory-system/internal/driver"
	"github.com/DeLuci/inventory-system/internal/handlers"
	"github.com/DeLuci/inventory-system/internal/models"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {

		}
	}(db.SQL)

	srv := &http.Server{
		Addr:    app.ServerAddress,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what I am going to put in the session
	gob.Register(models.Product{})
	gob.Register(models.Size{})
	gob.Register(models.User{})
	gob.Register(models.SearchBoot{})

	// TODO: change this to true when in production
	app.InProduction = false
	app.OneTimeUse = true

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	app, err := config.LoadConfig(".")
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(app.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to the database!!")

	//tc, err := render.CreateTemplateCache()
	//if err != nil {
	//	log.Fatal("cannot create template cache")
	//	return nil, err
	//}
	//
	//app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	//render.NewRenderer(&app)
	//helpers.NewHelpers(&app)

	return db, nil
}
