package main

import (
  "net/http"
  "fmt"
  "time"
  "log"
  "github.com/justinas/alice"
)

//type handler struct {}

//func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
//  fmt.Fprintf(w, "Welcome!")
//}

func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "Welcome!")
}
func main(){
  commonHandlers :=alice.New(loggingHandler, recoverHandler)
  http.Handle("/", commonHandlers.ThenFunc(indexHandler))
  http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
  http.ListenAndServe(":8080", nil)
}
func loggingHandler(next http.Handler) http.Handler{
  fn :=func(w http.ResponseWriter, r *http.Request){
    t1 :=time.Now()
    next.ServeHTTP(w,r)
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
  }
  return http.HandlerFunc(fn)
}


func indexHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){

  fmt.Fprintf(w, "you are on the about page")

}

func recoverHandler(next http.Handler) http.Handler{
  fn := func(w http.ResponseWriter, r *http.Request){
    defer func() {
      if err := recover(); err !=nil{
        log.Printf("panic: %+v", err)
        http.Error(w, http.StatusText(500), 500)
      }
    }()
    next.ServeHTTP(w, r)
  }
  return http.HandlerFunc(fn)
}

func myStripPrefix(h http.Handler) http.Handler{
  return http.StripPrefix("/old", h)
}
