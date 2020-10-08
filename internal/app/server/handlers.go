package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"text/template"
)

var (
	tpl                 = template.Must(template.ParseGlob("web/static/*.*html"))
	errMethodNotAllowed = errors.New("This method is not allowed")
)

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s.logger.Infof("Started request %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		s.logger.Infof("Request finished.")

	})
}

func (s *Server) handleSimple(page string) http.Handler {

	s.logger.Infof("Serving: %s", page)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			s.error(w, r, http.StatusMethodNotAllowed, errMethodNotAllowed)
			return
		}

		if err := tpl.ExecuteTemplate(w, page, nil); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
	})
}
