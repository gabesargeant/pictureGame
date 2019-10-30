package main

import (
	"fmt"
	"net/http"
	"time"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	t := time.Now()
	t.Format(time.RFC822Z)
	fmt.Fprintf(writer, "The Time is....., %s!", t)
	fmt.Println("Fresh Server hit!")
}

func activeGame(writer http.ResponseWriter, request *http.Request) {

	fmt.Print(request.URL)

	http.Redirect(writer, request, request.URL.Path+"/"+newGameSession(), http.StatusSeeOther)

}

func newGameSession() string {
	return RandomSession()
}
