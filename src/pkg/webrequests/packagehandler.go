package webrequests

import (
  "net/http"
  "encoding/json"
  logger "github.com/confused-Techie/atom-community-server-backend/src/pkg/logger"
  data "github.com/confused-Techie/atom-community-server-backend/src/pkg/data"
  //"github.com/gorilla/mux"
)

func PackageRootHandler(w http.ResponseWriter, r *http.Request) {
  // We can expect GET and POST requests here to list or create packages
  if r.Method == "GET" {
    pageKey := GetPage(r)
    sortKey := GetPackageSort(r)
    directionKey := GetPackageDirection(r)

    logger.InfoLogger.Println(pageKey)
    logger.InfoLogger.Println(sortKey)
    logger.InfoLogger.Println(directionKey)

    json.NewEncoder(w).Encode(data.GetPackagePointer())

  } else if r.Method == "POST" {
    // TODO
    JSONUnsupported(w, r)
  } else {
    // Otherwise this an unsupported method and should return the sitewide 404
    SiteWide404(w, r)
  }
}

func PackageSearchHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    //pageKey := GetPage(r)
    //queryKey := GetSearchQuery(r)
    //sortKey := GetSeachSort(r)
    //directionKey := GetPackageDirection(r)

  } else {
    SiteWide404(w, r)
  }
}

func PackageDetailHandler(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //packageName := vars["package_name"]
  if r.Method == "GET" {
    //engineKey := GetEngine(r)

  } else if r.Method == "DELETE" {

  } else {
    SiteWide404(w, r)
  }
}

func PackageStarHandler(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //packageName := vars["package_name"]
  if r.Method == "POST" {

  } else if r.Method == "DELETE" {

  } else {
    SiteWide404(w, r)
  }
}

func PackageStarGazersHandler(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //packageName := vars["package_name"]
  if r.Method == "GET" {

  } else {
    SiteWide404(w, r)
  }
}

func PackageUpdateHandler(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //packageName := vars["package_name"]
  if r.Method == "POST" {

  } else {
    SiteWide404(w, r)
  }
}

func PackageVersionHandler(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
  //packageName := vars["package_name"]
  //versionName := vars["version_name"]
  if r.Method == "GET" {

  } else if r.Method == "DELETE" {

  } else {
    SiteWide404(w, r)
  }
}
