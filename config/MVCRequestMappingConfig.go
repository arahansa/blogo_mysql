package config

import (
	"../controller"
	"net/http"
)

func InitMVCRequestConfig(){
	var article = controller.Article{}
	// mux := http.NewServeMux()
	http.HandleFunc("/muxTest", article.Article)
	http.HandleFunc("/", article.Article)
	http.HandleFunc("/article", article.Article)
	http.ListenAndServe(":80", nil)

}
