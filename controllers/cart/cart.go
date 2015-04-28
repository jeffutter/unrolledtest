package cart

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jeffutter/unrolledtest/controllers/cart/checkout"
	"github.com/jeffutter/unrolledtest/modules/template"
	"net/http"
)

func Router(main *mux.Router) *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	checkout.Router(r)

	r.HandleFunc("/cart", makeHandler(Cart)).Methods("GET")

	main.PathPrefix("/cart").Handler(negroni.New(
		negroni.Wrap(r),
	))

	return main
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func Cart(w http.ResponseWriter, r *http.Request) {
	templateData := template.TemplateData(nil, w, r)
	template.Render.HTML(w, http.StatusOK, "cart", templateData)
}
