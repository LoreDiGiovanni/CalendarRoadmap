package main

import(
    "net/http"
    "log"
)

func main()  {
    ip:= "localhost"
    port := "3080"
    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("/static/*", http.StripPrefix("/static/", fs))
    mux.HandleFunc("/", ErrorHandler(GetHomeHendler)) 
    mux.HandleFunc("GET /component/empty", ErrorHandler(GetComponentEmpty))
    mux.HandleFunc("GET /component/eventadder", ErrorHandler(GetComponentEventAdderHandler)) 
    mux.HandleFunc("GET /component/dropdowncolors", ErrorHandler(GetComponentDropDownColors))
    mux.HandleFunc("GET /component/colorsbutton", ErrorHandler(GetComponentColorsButton))
    mux.HandleFunc("POST /events", ErrorHandler(PostEventHandler)) 
    mux.HandleFunc("GET /404",ErrorHandler(GetFailHendler))
    log.Println("running on >> http://"+ip+":"+port)
    http.ListenAndServe(ip+":"+port,mux)
}
