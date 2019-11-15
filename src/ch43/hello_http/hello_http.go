package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8090", router))
}

//func main() {
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintf(writer, "Hello World!")
//	})
//
//	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
//		t := time.Now()
//		timerStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
//		writer.Write([]byte(timerStr))
//	})
//
//	http.ListenAndServe(":8090", nil)
//}
