package webrequests

import (
  "net/http"
  "encoding/json"
  //logger "github.com/confused-Techie/atom-community-server-backend/src/pkg/logger"
  models "github.com/confused-Techie/atom-community-server-backend/src/pkg/models"
)

func SiteWide404(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404)
  err := models.HTTPError{
    Message: "This is a standin for the proper site wide 404 page.",
  }

  json.NewEncoder(w).Encode(err)
}

func JSON404(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404)
  err := models.HTTPError{
    Message: "Not Found",
  }
  json.NewEncoder(w).Encode(err)
}

func JSONMissingAuth(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(401)
  err := models.HTTPError{
    Message: "Requires authentication. Please update your token if you haven't done so recently.",
  }
  json.NewEncoder(w).Encode(err)
}

func JSONServerError(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(500)
  err := models.HTTPError{
    Message: "Application error",
  }
  json.NewEncoder(w).Encode(err)
}

func JSONUnsupported(w http.ResponseWriter, r *http.Request) {
  // This error endpoint will ideally be removed once development is complete.
  w.WriteHeader(500)
  err := models.HTTPError{
    Message: "While under development this feature is not supported.",
  }
  json.NewEncoder(w).Encode(err)
}
