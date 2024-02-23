package main

import(
    "net/http"
    "log"
)

//w.Write([]byte("<h1>Home</h1>"))
func main()  {
    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("/static/*", http.StripPrefix("/static/", fs))
    mux.HandleFunc("GET /a", func(w http.ResponseWriter ,r *http.Request){
        Layout("home").Render(r.Context(),w)
    }) 
    mux.HandleFunc("GET /home", func(w http.ResponseWriter ,r *http.Request){
        Home("ldg").Render(r.Context(),w)
    }) 
    log.Println("running on >> http://127.0.0.1:8080")
    http.ListenAndServe(":8080",mux)
}
