package main

import (
    "log"
    "net/http"
    "routing"
    "fmt"
    "config"
    "github.com/rs/cors"
)



func main() {


    fmt.Println(`

                    .---.
                    |[X]|
                    _.==._.""""".___n__
                    d __ ___.-''-. _____b
                    |[__]  /."""".\ _   |
                    |     // /""\ \\_)  |         Selfie Roulette Api... Cheese
                    |     \\ \__/ //    |
                    |kodak \'.__.'/     |
                    \======='-..-'======/
                     '-----------------'

    `)



    router := routing.NewRouter()
    handler := cors.New(cors.Options{
        AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
    }).Handler(router)
    log.Fatal(http.ListenAndServe(config.Config.ServerAddress, handler))

}
