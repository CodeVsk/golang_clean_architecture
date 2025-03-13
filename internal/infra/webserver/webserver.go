package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router     chi.Router
	ServerPort string
	Handlers   map[string]map[string]http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:     chi.NewRouter(),
		ServerPort: serverPort,
		Handlers:   make(map[string]map[string]http.HandlerFunc),
	}
}

func (ws *WebServer) RegisterHandler(route string, method string, handler http.HandlerFunc) {
	if _, exists := ws.Handlers[route]; !exists {
		ws.Handlers[route] = make(map[string]http.HandlerFunc)
	}

	ws.Handlers[route][method] = handler
}

func (ws *WebServer) Run() {
	for route, methods := range ws.Handlers {
		for method, handler := range methods {
			ws.Router.Method(method, route, handler)
		}
	}

	if err := http.ListenAndServe(ws.ServerPort, ws.Router); err != nil {
		panic(err)
	}
}
