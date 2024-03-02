package main

import (
	"RoadmapCalendar/types"
	"fmt"
    "encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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

func mdToHTML(md []byte) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock 

	p := parser.NewWithExtensions(extensions)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank 	
    opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.ToHTML(md, p, renderer))
}

func GetHomeHendler(w http.ResponseWriter, r *http.Request)error{
    test := []byte(`## markdown document $x^2$
- item $x^2$ 
- item 2`)
    tmp :=  mdToHTML(test)
    e := []types.Events{
        types.Events{Title: "Prova impegno 1",
                    Notes: "Prova Nota ",
                    Date: "20/06/2025",
                    Time_start: "11:00",
                    Time_end: "12:00",
                    Tags: []string{"#Tag1","#Tag2"},},
        types.Events{Title: "Prova impegno 2",Notes: "Prova Nota ",Date: "20/06/2025",Time_start: "11:00",Time_end: "12:00"},
        types.Events{Title: "Prova impegno 3",Notes: tmp,Date: "20/06/2025",Time_start: "11:00",Time_end: "12:00"}}
    currentTime := time.Now()
    page := Home(e, currentTime.Format("2006/01/02"))
    return RenderView(w,r,page,"/")
}
func PostEventHandler(w http.ResponseWriter, r *http.Request)error{
    var event types.Events
    json.NewDecoder(r.Body).Decode(&event)
    defer r.Body.Close()
    fmt.Println(event)
    page := Empty()
    return RenderView(w,r,page,"/404")
}

func GetComponentEventAdderHandler(w http.ResponseWriter, r *http.Request)error{
    return RenderView(w,r,InputEvent(),"/")
}
func GetComponentDropDownColors(w http.ResponseWriter, r *http.Request)error{
    colors := []types.Color{
    types.NewColor("c001", "#ff6c6b"),
    types.NewColor("c002", "#da8548"),
    types.NewColor("c003", "#98be65"),
    types.NewColor("c004", "#4db5bd"),
    types.NewColor("c005", "#ECBE7B"),
    types.NewColor("c006", "#51afef"),
    types.NewColor("dark-blue", "#2257A0"),
    types.NewColor("magenta", "#c678dd"),
    types.NewColor("violet", "#a9a1e1"),
    types.NewColor("cyan", "#46D9FF"),
    types.NewColor("dark-cyan", "#5699AF"),}

    page := DropDownColors(colors)
    return RenderView(w,r,page,"/")
}

func GetFailHendler (w http.ResponseWriter, r *http.Request)error{
    page := Fail()
    return RenderView(w,r,page,"/404")
}

func GetComponentColorsButton(w http.ResponseWriter, r *http.Request)error{
    page := ColorsButton()
    return RenderView(w,r,page,"/404")
}
func GetComponentEmpty (w http.ResponseWriter, r *http.Request)error{
    page := Empty()
    return RenderView(w,r,page,"/404")
}
