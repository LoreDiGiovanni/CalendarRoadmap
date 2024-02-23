package main

import(
    "fmt"
    "net/http"
    //"github.com/a-h/templ"
)

type APIServer struct{
    address string
}

func (s* APIServer) Run() {
    mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServer(http.Dir("./static")))
    mux.HandleFunc("GET /login", func(w http.ResponseWriter ,r *http.Request){
        //w.Write([]byte("<h1>Home</h1>"))
        Hello().Render(r.Context())
    }) 
    fmt.Println("Server on 8080")
    http.ListenAndServe(":8080",mux)
}


