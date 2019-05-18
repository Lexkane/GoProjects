package main

import "net/http"

func (s *server) handleTemplate(flies ...string) http.HandlerFunc {
	tpl, err := s.loadTemplates(files...)
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s.rednerTemplate(tpl)
	}
}

func main() {
	s := newServer()
	http.HandleFunc("/", s.handleTemplate("layout.tpl.html", "index.tpl.html"))

}
