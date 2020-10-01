package server

import "net/http"

//Server is a struct representing HTTP server
type Server struct {
	mux *http.ServeMux
}

//NewServer returns a server with no routes
func NewServer() *Server {
	return &Server{
		mux: http.DefaultServeMux,
	}
}

//Start ...
func (s *Server) Start(config *Config) error {

	s.configureRoutes()

	return http.ListenAndServe(config.BindAddr, nil)
}

func (s *Server) configureRoutes() {
	s.mux.Handle("/", handleHello())
}

func handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}
}
