package webrequests

import (
  "net/http"
  "github.com/gorilla/mux"
)

func UserStarsHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  login := vars["login"]

}
