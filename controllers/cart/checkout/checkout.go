package checkout

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"unrolledtest/modules/template"
)

func Router(main *mux.Router) *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/cart/checkout", Checkout).Methods("GET")

	main.PathPrefix("/cart/checkout").Handler(negroni.New(
		negroni.Wrap(r),
	))

	return main
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	templateData := template.TemplateData(nil, w, r)
	template.Render.HTML(w, http.StatusOK, "cart/checkout", templateData)
}
