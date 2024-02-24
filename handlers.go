package main

import (
	"log"
	"net/http"
	"github.com/a-h/templ"
)

type genericHandle func(http.ResponseWriter,*http.Request) error

func RenderView(w http.ResponseWriter, r *http.Request, view templ.Component, layoutPath string) error{
	if r.Header.Get("Hx-Request") == "true" {
		err := view.Render(r.Context(), w)
		return err
	}else{
        err := Layout(layoutPath).Render(r.Context(), w)
        return err   
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
    username:= "User"
    page := Home(username)
    err := RenderView(w,r,page,"/")
    return err 
}
