package handlers

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "model"
    "net/http"
    "persistence"
    "responseErrors"
)

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
    var photo model.Photo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &photo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    photo, error := persistence.RepoUpdatePhoto(photo)

    if error == nil {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        if err := json.NewEncoder(w).Encode(photo); err != nil {
            panic(err)
        }
    } else {
        w.WriteHeader(http.StatusNotModified)
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        if err := json.NewEncoder(w).Encode(responseErrors.JSONErr{Code: http.StatusNotModified, Text: "Photo not modified"}); err != nil {
            panic(err)
        }
    }


}

