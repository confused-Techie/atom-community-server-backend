package main

import (
  logger "github.com/confused-Techie/atom-community-server-backend/src/pkg/logger"
  webrequests "github.com/confused-Techie/atom-community-server-backend/src/pkg/webrequests"
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "os/signal"
  "syscall"
)

func main() {

  // listen to SIGINT calls
  captureExit := make(chan os.Signal)
  signal.Notify(captureExit, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
  go func() {
    <-captureExit
    // add logging, and ensuring to write any data in memory
    os.Exit(1)
  }()

  mux := mux.NewRouter()

  // ======== ORIGINAL API SERVER ENDPOINTS ==========

  // This is the major handler for all Packages
  // -- List Packages: GET /api/packages
  // -- Search Packages: GET /api/packages/search
  // -- Package Details: GET /api/packages/:package_name
  // -- Create Package: POST /api/packages
  // -- Delete Package: DELETE /api/packages/:package_name
  // -- Package Versions: GET /api/packages/:package_name/versions/:version_name
  // -- New Package Version: POST /api/packages/:package_name/versions
  // -- Delete Version: DELETE /api/packages/:package_name/versions/:version_name
  // -- Starring a Package: POST /api/packages/:name/star
  // -- Unstar a Package: DELETE /api/packages/:name/star
  // --List Package's Stargazers: GET /api/packages/:name/stargazers

  // List Packages: GET /api/packages
  // Create Package: POST /api/packages
  mux.HandleFunc("/api/packages", webrequests.PackageRootHandler)
  // Search Packages: GET /api/packages/search
  mux.HandleFunc("/api/packages/search", webrequests.PackageSearchHandler)
  // Package Details: GET /api/packages/:package_name
  // Delete Package: DELETE /api/packages/:package_name
  mux.HandleFunc("/api/packages/{package_name}", webrequests.PackageDetailHandler)  // ex. /api/packages/zelda-botw-ui
  // Starring a Package: POST /api/packages/:name/star
  // Unstar a Package: DELETE /api/packages/:name/star
  mux.HandleFunc("/api/packages/{package_name}/star", webrequests.PackageStarHandler)
  // List Package's Stargazers: GET /api/packages/:name/stargazers
  mux.HandleFunc("/api/packages/{package_name}/stargazers", webrequests.PackageStarGazersHandler)
  // New Package Version: POST /api/packages/:package_name/versions
  mux.HandleFunc("/api/packages/{package_name}/versions", webrequests.PackageUpdateHandler)
  // Package Versions: GET /api/packages/:package_name/versions/:version_name
  // Delete Version: DELETE /api/packages/:package_name/versions/:version_name
  mux.HandleFunc("/api/packages/{package_name}/versions/{version_name}", webrequests.PackageVersionHandler)
  //mux.Handle("/api/packages", http.HandlerFunc(webrequests.PackageHandler))

  // The User Endpoint
  // List a User's Starred Packages: GET /api/users/:login/stars
  mux.HandleFunc("/api/users/{login}/stars", webrequests.UserStarsHandler)
  //mux.Handle("/api/users", http.HandlerFunc(webrequests.UsersHandler))

  // Stars Endpoint
  // List Authenticated User's Starred Packages: GET /api/stars
  mux.HandleFunc("/api/stars", webrequests.StarsHandler)
  //mux.Handle("/api/stars", http.HandlerFunc(webrequests.StarsHandler))

  // Listing Atom Updates: GET /api/updates
  mux.HandleFunc("/api/updates", webrequests.AtomUpdateHandler)

  logger.InfoLogger.Println("Listening on 8080")
  logger.ErrorLogger.Fatal(http.ListenAndServe(":8080", mux))
}
