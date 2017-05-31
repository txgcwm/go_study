package main


import (
    // "fmt"
    "log"
    "net/http"
    "github.com/issue9/mux"
)

func main() {
    router := mux.New(true, true, nil, nil)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}


// https://git.oschina.net/caixw/mux
// https://my.oschina.net/zijingshanke/blog/907955