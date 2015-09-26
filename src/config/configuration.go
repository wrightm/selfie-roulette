package config


import (
    "encoding/json"
    "os"
    "log"
)

type Configuration struct {
    BaseDir       string `json:"baseDir"`
    ImagesDir     string `json:"imagesDir"`
    WebAppDir     string `json:"webAppDir"`
    RepoDir       string `json:"repoDir"`
    ServerAddress string `json:"serverAddress"`
    ImageFileType string `json:"imageFileType"`
}

var Config Configuration

func init() {

    file, _ := os.Open("src/config/properties.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Fatal("Cannot open configuration file:", err)
    }
    defer file.Close()
    Config = configuration
}
