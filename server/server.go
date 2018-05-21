package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func StartServer() {
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())

	r := mux.NewRouter()

	r.HandleFunc("/", welcomeHandler)

	n.UseHandler(r)

	err := http.ListenAndServe(fmt.Sprintf(":%d", 3000), n)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func welcomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Welcome to POC ctw")
}
