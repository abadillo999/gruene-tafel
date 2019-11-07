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
	config  *config.ServerConfig
}

func NewServer (config *config.ServerConfig, processor *processor.Processor) *Server {
	_server := Server{}
	_server.config = config
	_server.router = mux.NewRouter()
	_server.handler = NewHandler(processor)
	return &_server
}


func (server *Server) Run() {
	server.setRouters()
	fmt.Println("Listen and serve on port" + server.config.Port)
	log.Fatal(http.ListenAndServe(server.config.Port, server.router))
}

//Private functions
func (server *Server) setRouters() {
	server.post("/task", server.handleRequest(server.handler.CreateTask))
	server.get("/task/{taskId}", server.handleRequest(server.handler.GetTask))
	server.post("/task/{taskId}", server.handleRequest(server.handler.UpdateTask))
	server.delete("/task/{taskId}", server.handleRequest(server.handler.DeleteTask))
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