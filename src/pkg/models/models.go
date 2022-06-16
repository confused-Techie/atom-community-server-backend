package models

import (

)

type HTTPError struct {
  Message string `json:"message"`
}

type PackagePointerSlice struct {
  Pointers []*PackagePointer
}

type PackagePointer struct {
  Name string `json:"name"`
  Names []string `json:"names"`
  Pointer string `json:"pointer"`
  Downloads int `json:"downloads"`
  StargazersCount int `json:"stargazers_count"`
  Stargazers []string `json:"stargazers"`
  Created string `json:"created_at"`
  Updated string `json:"updated_at"`
}

type Package struct {
  Name string `json:"name"`
  Repository *PackageRepository `json:"repository"`
  Releases []*PackageReleases `json:"releases"`
  ReadMe string `json:"readme"`
  MetaData *PackageMetaData `json:"metadata"`
}

type PackageRepository struct {
  Type string `json:"type"`
  URL string `json:"url"`
}

type PackageReleases struct {
  Release map[string]string
}

type PackageMetaData struct {
  // Here will likely need to be fully compatible with the standard package.json file.
  // additionally here is where the many checks for non-existant values should be.
  Name string `json:"name"`
  Main string `json:"main,omitempty"`
  Version string `json:"version"`
  Description string `json:"description,omitempty"`
  ActivationCommands *PackageActivationCommands `json:"activationCommands,omitempty"`
  Repository string `json:"repository"`
  License string `json:"license,omitempty"`
  Engines *PackageEngines `json:"engines,omitempty"`
  Dependencies *PackageDependencies `json:"dependencies"`
}

type PackageActivationCommands struct {
  // todo
}

type PackageEngines struct {
  // todo
}

type PackageDependencies struct {

}
