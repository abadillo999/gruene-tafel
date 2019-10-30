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
	//_DB     *gorm.DB
}

func (server *Server) Initialize(config *config.Config, processor *processor.Processor) {
	server.router = mux.NewRouter()
	server.setRouters()
}


func (server *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, server.router))
}

//Private functions
func (server *Server) setRouters() {
	server.post("/order", a.handleRequest(handler.CreateOrder))
	server.get("/order/{orderId}", a.handleRequest(handler.GetOrder))
	server.post("/order/{orderId}", a.handleRequest(handler.UpdateOrder))
	server.delete("/order/{orderId}", a.handleRequest(handler.DeleteOrder))
}

func (server *Server) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("GET")
}

func (server *Server) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("POST")
}

func (server *Server) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("PUT")
}

func (server *Server) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (server *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}