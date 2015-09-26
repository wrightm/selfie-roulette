package routing

import (
    "time"
    "log"
    "net/http"
    "handlers"
)

type Router struct {
    Router *mux.Router
}

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var router *Router

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "GetPhotos",
        "GET",
        "/photos",
        handlers.GetPhotos,
    },
    Route{
        "UpdatePhoto",
        "PUT",
        "/photos/{photoId}",
        handlers.UpdatePhoto,
    },
    Route{
        "NewPhoto",
        "POST",
        "/photos",
        handlers.NewPhoto,
    },
    Route{
        "GetPhoto",
        "GET",
        "/photos/{photoId}",
        handlers.GetPhoto,
    },
    Route{
        "DeletePhoto",
        "DELETE",
        "/photos/{photoId}",
        handlers.DeletePhoto,
    },
}

func NewRouter() *mux.Router {

    router = &Router{Router: mux.NewRouter()}
    router.Router.StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Log(handler, route.Name)
        router.Router.
        Methods(route.Method).
        Path(route.Pattern).
        Name(route.Name).
        Handler(handler)

    }

    return router.Router
}

func Log(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
