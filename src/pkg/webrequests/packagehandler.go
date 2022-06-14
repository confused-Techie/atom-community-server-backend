package webrequests

import (
  "net/http"
  logger "github.com/confused-Techie/atom-community-server-backend/src/pkg/logger"
  "github.com/gorilla/mux"
)

func PackageRootHandler(w http.ResponseWriter, r *http.Request) {
  // We can expect GET and POST requests here to list or create packages
  if r.Method == "GET" {

  } else if r.Method == "POST" {

  }
  // ELSE
}

func PackageSearchHandler(w http.ResponseWriter, r *http.Request) {

}

func PackageDetailHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  packageName := vars["package_name"]

}

func PackageStarHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  packageName := vars["package_name"]

}

func PackageStarGazersHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  packageName := vars["package_name"]

}

func PackageUpdateHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  packageName := vars["package_name"]

}

func PackageVersionHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  packageName := vars["package_name"]
  versionName := vars["version_name"]

}
