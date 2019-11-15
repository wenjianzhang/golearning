package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World!")
	})

	http.HandleFunc("/time", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timerStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		writer.Write([]byte(timerStr))
	})

	http.ListenAndServe(":8090", nil)
}
