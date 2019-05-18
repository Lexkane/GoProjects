package main

func (s *server) handleTemplate(files ...string) http.HandlerFunc {
	var (
		init sync.Once
		tpl  *template.Template
		err  error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, err = s.loadTemplates(files...)
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServererror)
			return
		}
	}
	s.renderTemplate(tpl)

}

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !userIsAdmin(r) {
			http.Error(w, "Admin only", http.StatusUnauthorized)
			return
		}
		h(w, r)
	}
}
