package main

import (
	"RoadmapCalendar/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)
type Server struct{
   address string 
   db  storage.Storage 
   mux *http.ServeMux
}

func NewServer(adress string, db storage.Storage) Server{
    mux := http.NewServeMux()
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("/static/*", http.StripPrefix("/static/", fs))
    return Server{address: adress,db: db,mux: mux} 
}
func (s Server)Run(){
    http.ListenAndServe(s.address,s.mux)
}

func main()  {
    
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    storage, err := storage.NewMongoStore()
    if err== nil{
        server := NewServer("localhost:3080",storage)
        server.mux.HandleFunc("/", ErrorHandler(server.GetHomeHendler)) 
        server.mux.HandleFunc("GET /component/empty", ErrorHandler(server.GetComponentColorsButton))
        server.mux.HandleFunc("GET /component/roadmap", ErrorHandler(server.GetRoadmapHendler))
        server.mux.HandleFunc("GET /component/eventadder", ErrorHandler(server.GetComponentEventAdderHandler)) 
        server.mux.HandleFunc("GET /component/dropdowncolors", ErrorHandler(server.GetComponentDropDownColors))
        server.mux.HandleFunc("GET /component/colorsbutton", ErrorHandler(server.GetComponentColorsButton))
        server.mux.HandleFunc("GET /component/inputeventplaceholder",ErrorHandler(server.GetInputEventPlaceholder))
        server.mux.HandleFunc("GET /component/login",ErrorHandler(server.GetComponentLogin))
        server.mux.HandleFunc("GET /component/sigin",ErrorHandler(server.GetComponentSigin))
        server.mux.HandleFunc("POST /event", ErrorHandler(server.PostEventHandler)) 
        server.mux.HandleFunc("POST /account", ErrorHandler(server.PostAccountHandler)) 
        server.mux.HandleFunc("POST /account/new", ErrorHandler(server.PostAccountNewHandler)) 
        server.mux.HandleFunc("GET /404",ErrorHandler(server.GetFailHendler))
        log.Println("running on >> http://"+server.address)
        server.Run()
    }else{
        log.Println("Error database connection")
    }
}
