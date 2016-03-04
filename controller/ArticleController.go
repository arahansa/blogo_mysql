package controller


import (
	"net/http"
	"log"
	"regexp"
	"strings"
)

var rNum = regexp.MustCompile(`\d`)  // Has digit(s)
var rAbc = regexp.MustCompile(`abc`) // Contains "abc"

type Article struct{}

func (a *Article) Article(w http.ResponseWriter, r *http.Request){
	values := strings.Split(r.URL.Path, "/")
	log.Println("경로들 : ", values, len(values))
	
	switch  {
	case rNum.MatchString(r.URL.Path):
		log.Println("숫자만 있네요?" , r.URL.Path)
	case rAbc.MatchString(r.URL.Path) :
		log.Println("abc를 포함하네요 ? ", r.URL.Path)
	}
	log.Println(" r :: ", r)
	log.Println(" Method ::"+r.Method)
	log.Println(" Path :: " , r.RequestURI)
	log.Println(" Form :: ", r.Form)
	log.Println(" URL  :: ", r.URL)
}
