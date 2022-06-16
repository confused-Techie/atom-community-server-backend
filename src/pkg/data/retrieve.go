package data

import (
  models "github.com/confused-Techie/atom-community-server-backend/src/pkg/models"
  logger "github.com/confused-Techie/atom-community-server-backend/src/pkg/logger"
  "io/ioutil"
  "os"
  "encoding/json"
)

// ReadFile is intended to handle reading raw data from the file system, making the choices on how to do that as needed.
func ReadFile(loc string) ([]byte, error) {
  file, err := os.OpenFile(loc, os.O_RDWR|os.O_APPEND, 0666)

  if err != nil {
    return nil, err
  }

  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    return nil, err
  }
  return bytes, nil
}

// WriteFile is intended to handle writing raw data to the file system.
func WriteFile() {

}

// GetPackagePointer is intended to only retrieve the data of the PackagePointer file and return it.
func GetPackagePointer() []models.PackagePointer {
  b, err := ReadFile("./data/packages/packages_pointer.json")
  if err != nil {
    logger.ErrorLogger.Println(err)
  }
  var pps []models.PackagePointer
  json.Unmarshal(b, &pps)
  return pps
}

// Get Package is intended to only retrieve the data of a single package, and return it.
func GetPackage() {

}
