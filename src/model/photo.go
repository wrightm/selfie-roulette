package model

type Photo struct {
    Id       string `json:"id"`
    Name     string `json:"name"`
    Filename string `json:"filename"`
    Winner   bool   `json:"winner"`
}

type Photos []Photo
