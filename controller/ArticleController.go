package controller


import (
	"net/http"
	"log"
)



type Article struct{}

func (a *Article) Article(w http.ResponseWriter, r *http.Request){
	log.Println(" r :: ", r)
	log.Println(" Method ::"+r.Method)
	log.Println(" Path :: " , r.RequestURI)
	log.Println(" Form :: ", r.Form)
	log.Println(" URL  :: ", r.URL)
}
