package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jeffutter/unrolledtest/controllers/cart"
	"github.com/jeffutter/unrolledtest/modules/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	utils.PanicIf(err)

	r := mux.NewRouter()
	r.StrictSlash(true)

	cart.Router(r)

	n := negroni.Classic()
	n.UseHandler(r)

	port := "3001"
	n.Run(fmt.Sprintf(":%s", port))

}
