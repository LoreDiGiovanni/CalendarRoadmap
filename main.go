package main

import(
    "net/http"
    "log"
)

func main()  {
    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("/static/*", http.StripPrefix("/static/", fs))
    mux.HandleFunc("/", ErrorHandler(GetHomeHendler)) 
    log.Println("running on >> http://127.0.0.1:8080")
    http.ListenAndServe(":8080",mux)
}
