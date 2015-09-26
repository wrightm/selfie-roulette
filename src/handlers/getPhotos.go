package handlers

import (
    "encoding/json"
    "net/http"
    "persistence"
)

func GetPhotos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(persistence.RepoPhotos()); err != nil {
        panic(err)
    }
}
