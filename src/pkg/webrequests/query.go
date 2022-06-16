package webrequests

import(
  "net/http"
)

// ============ Specials to List ==================================

func GetPage(r *http.Request) (string) {
  // This is intended for use within the /api/package call for the query parameter page.
  // TODO: Ensure the resulting key is a numerical value.
  defaultKey := "1"

  keys, ok := r.URL.Query()["page"]

  if !ok || len(keys[0]) < 1 {
    // URL param page is missing. Default to 1
    return defaultKey
  }

  key := keys[0]
  return key
}

func GetPackageSort(r *http.Request) (string) {
  // /api/package call query sort
  // Valid Queries: downloads, created_at, updated_at, stars
  // Defaults to: downloads
  // Will use default variable if provided variable is invalid
  validKeys := []string{"downloads", "created_at", "updated_at", "stars"}
  defaultKey := "downloads"

  keys, ok := r.URL.Query()["sort"]

  if !ok || len(keys[0]) < 1 {
    // URL param sort is missing.Default to downloads
    return defaultKey
  }

  key := keys[0]

  if contains(validKeys, key) {
    // the key is within the slice of valid keys for this sorting method
    return key
  }
  // otherwise it seems the key is not within the valid list of sort options.
  // return default
  return defaultKey
}

func GetPackageDirection(r *http.Request) (string) {
  // /api/package call query direction
  // Valid Queries: asc, desc
  // Defaults to desc,
  // If sorting by stars it can only be desc, originally.
  validKeys := []string{"asc", "desc"}
  defaultKey := "desc"

  keys, ok := r.URL.Query()["direction"]

  if !ok || len(keys[0]) < 1 {
    // url param direction is missing. Default to desc
    return defaultKey
  }

  key := keys[0]

  if contains(validKeys, key) {
    // key is within valid list,
    return key
  }
  // key is not within valid list. Return default
  return defaultKey
}

// ================== Specials for Search ==========================

func GetSearchSort(r *http.Request) (string) {
  // GetSearchSort is identical to GetPackageSort EXCEPT it defaults to relevance.
  validKeys := []string{"downloads", "created_at", "updated_at", "stars", "relevance"}
  defaultKey := "relevance"

  keys, ok := r.URL.Query()["sort"]

  if !ok || len(keys[0]) < 1 {
    // url param sort is missing default
    return defaultKey
  }

  key := keys[0]

  if contains(validKeys, key) {
    return key
  }
  return defaultKey
}

func GetSearchQuery(r *http.Request) (string) {
  // this has no defaults, and will return empty if there is no valid key, otherwise will return whatever the user passed.
  // TODO prevent types of vulnerabilities that may be related to passing this to the search.
  keys, ok := r.URL.Query()["q"]

  if !ok || len(keys[0]) < 1 {
    return ""
  }
  return keys[0]
}

// =========== Specials for PackageDetails =============

func GetEngine(r *http.Request) (string) {
  // TODO confirm that engine provided is valid SemVer version. Possibly even Atom Version.
  keys, ok := r.URL.Query()["engine"]

  if !ok || len(keys[0]) < 1 {
    // nothing provided. Should default to none, not pruning results by engine.
    return ""
  }

  // otherwise an engine was provided, return it
  return keys[0]
}

// contains checks if a string is present in a slice of strings
func contains(s []string, str string) bool {
  for _, v := range s {
    if v == str {
      return true
    }
  }
  return false
}
