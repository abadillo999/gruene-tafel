package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"config"
	"processor"
)

type Server struct {
	router *mux.Router
	handler *Handler
}

func NewServer (config *config.Config, processor *processor.Processor) *Server {
	server := Server{}
	server.router = mux.NewRouter()
	return &server
}


func (server *Server) Run(port string) {
	server.setRouters()
	fmt.Println("Listen and serve...")
	log.Fatal(http.ListenAndServe(port, server.router))
}

//Private functions
func (server *Server) setRouters() {
	server.post("/order", server.handleRequest(server.handler.CreateOrder))
	server.get("/order/{orderId}", server.handleRequest(server.handler.GetOrder))
	server.post("/order/{orderId}", server.handleRequest(server.handler.UpdateOrder))
	server.delete("/order/{orderId}", server.handleRequest(server.handler.DeleteOrder))
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