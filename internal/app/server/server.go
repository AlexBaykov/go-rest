package server

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

//Server is a struct representing HTTP server
type Server struct {
	mux    *http.ServeMux
	logger *logrus.Logger
}

//NewServer returns a server with no routes
func NewServer() *Server {
	return &Server{
		mux:    http.NewServeMux(),
		logger: logrus.New(),
	}
}

//Start ...
func (s *Server) Start(config *Config) error {

	s.configureRoutes()

	s.configureLogger(config.LogLevel)

	return http.ListenAndServe(config.BindAddr, s.mux)
}

func (s *Server) configureRoutes() {

	//This handler is for static part of the website: homepage and blog.
	staticHandler := http.FileServer(http.Dir("./web/static/"))

	s.mux.Handle("/", s.logRequest(staticHandler))
}

func (s *Server) configureLogger(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		log.Fatal(err)
	}
	s.logger.SetLevel(level)
}
