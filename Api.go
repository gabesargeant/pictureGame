package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	t := time.Now()
	t.Format(time.RFC822Z)
	fmt.Fprintf(writer, "The Time is....., %s!", t)
	fmt.Println("Fresh Server hit!")
}

// game matches on url /game/
func game(writer http.ResponseWriter, request *http.Request) {

	u, err := url.Parse(request.URL.Path)
	if err != nil {
		log.Print("error")
		log.Fatal(err)
	}
	s := u.Path
	//fmt.Print(u)

	isAlpha := regexp.MustCompile(`/game/[^[A-Za-z]+$]/`).MatchString

	if isAlpha(s) {
		log.Printf("Game running %s", s)

	}
	if !isAlpha(s) {
		redir := "/game/" + newGameSession()
		log.Printf("redirecting to %s", redir)
		http.Redirect(writer, request, redir, http.StatusSeeOther)
	}

}

func newGameSession() string {
	return RandomSession() + "/"
}
