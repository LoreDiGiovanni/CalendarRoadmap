package main

import (
	"log"
    "time"
	"net/http"
	"github.com/a-h/templ"
    "RoadmapCalendar/types"
)

type genericHandle func(http.ResponseWriter,*http.Request) error

func RenderView(w http.ResponseWriter, r *http.Request, view templ.Component, layoutPath string) error{
	if r.Header.Get("Hx-Request") == "true" { 
        return view.Render(r.Context(), w)
	}else{
        return Layout(layoutPath).Render(r.Context(), w)
    }
	 
}

func ErrorHandler(f genericHandle) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
        err := f(w,r)
        if err != nil{
            //tools.WriteJSON(w,http.StatusBadRequest,ApiError{Error: err.Error()})
            log.Println("Error:",err)
        }
    }
}

func GetHomeHendler(w http.ResponseWriter, r *http.Request)error{
    e := []types.Events{
        types.Events{Title: "Prova impegno 1",Notes: "Prova Nota ",Date: "20/06/2025",Time_start: "11:00",Time_end: "12:00"},
        types.Events{Title: "Prova impegno 2",Notes: "Prova Nota ",Date: "20/06/2025",Time_start: "11:00",Time_end: "12:00"},
        types.Events{Title: "Prova impegno 3",Notes: "Prova Nota",Date: "20/06/2025",Time_start: "11:00",Time_end: "12:00"}}
    currentTime := time.Now()
    page := Home(e, currentTime.Format("2006/01/02"))
    return RenderView(w,r,page,"/")
}
func GetComponentEventAdderHandler(w http.ResponseWriter, r *http.Request)error{
    return RenderView(w,r,InputEvent(),"/")
}
func GetComponentDropDownColors(w http.ResponseWriter, r *http.Request)error{
    colors := []types.Color{
    types.NewColor("Red", "#ff0000"),
    types.NewColor("Green", "#282c34"),}

    page := DropDownColors(colors)
    return RenderView(w,r,page,"/")
}

func GetFailHendler (w http.ResponseWriter, r *http.Request)error{
    page := Fail()
    return RenderView(w,r,page,"/404")
}
func GetComponentEmpty (w http.ResponseWriter, r *http.Request)error{
    page := Empty()
    return RenderView(w,r,page,"/404")
}
