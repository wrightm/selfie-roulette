package handlers

import (
    "encoding/json"
    "net/http"
    "persistence"
    "responseErrors"
    "github.com/gorilla/mux"
)

func GetPhoto(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var photoId string
    var ok bool
    if photoId, ok = vars["photoId"]; !ok {
        panic("cannot find {photoId}")
    }
    photo, err := persistence.RepoFindPhoto(photoId)
    if err == nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(photo); err != nil {
            panic(err)
        }
        return
    }

    // If we didn't find it, 404
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)
    if err := json.NewEncoder(w).Encode(responseErrors.JSONErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
        panic(err)
    }

}
