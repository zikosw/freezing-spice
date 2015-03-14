package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yosssi/ace"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("view/base", "view/inner", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love ")
}

func GetRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/ziko", handler2)
	return r
}
