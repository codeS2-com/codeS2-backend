package app

import (
	"github.com/codeS2-com/codeS2-backend/app/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/code", a.getAllCode)

	//a.Post("/projects/{title}/tasks", a.CreateTask)
	//a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.GetTask)
	//a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.UpdateTask)
	//a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.DeleteTask)
}

/*
** Projects Handlers
 */
func (a *App) GetAllCode(w http.ResponseWriter, r *http.Request) {
	handler.GetAllCode(w)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
