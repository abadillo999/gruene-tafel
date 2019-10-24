package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"server/handler"
	"config"
)

type Server struct {
	router *mux.Router
	_DB     *gorm.DB
}

func (server *Server) initialize(config *config.Config, processor *processor.Processor) {
	server.router = mux.NewRouter()
	server.setRouters()
}

// setRouters sets the all required routers
func (server *Server) setRouters() {
	// Routing for handlers
	server.post("/order", a.handleRequest(handler.CreateOrder))
	server.get("/order/{orderId}", a.handleRequest(handler.GetOrder))
	server.post("/order/{orderId}", a.handleRequest(handler.UpdateOrder))
	server.delete("/order/{orderId}", a.handleRequest(handler.DeleteOrder))
}

// Get wraps the router for GET method
func (server *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (server *Server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (server *Server) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (server *Server) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("DELETE")
}

// Run the server on it's router
func (server *Server) run(port string) {
	log.Fatal(http.ListenAndServe(port, server.router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (server *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}