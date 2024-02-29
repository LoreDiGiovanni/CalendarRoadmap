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
    mux.HandleFunc("GET /component/empty", ErrorHandler(GetComponentEmpty))
    mux.HandleFunc("GET /component/eventadder", ErrorHandler(GetComponentEventAdderHandler)) 
    mux.HandleFunc("GET /component/dropdowncolors", ErrorHandler(GetComponentDropDownColors))
    mux.HandleFunc("GET /component/colorsbutton", ErrorHandler(GetComponentColorsButton))
    mux.HandleFunc("GET /404",ErrorHandler(GetFailHendler))
    log.Println("running on >> http://192.168.1.50:3080")
    http.ListenAndServe("192.168.1.50:3080",mux)
}
