package main

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func indexHandler(context *appContext) http.Handler {
	return (http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			index(context, w, r)
		}))
}

func index(context *appContext, writer http.ResponseWriter, request *http.Request) {

	sess := newGameSession()
	cookie := newUserCookie()

	uSess := userInSessions{
		session: sess,
	}
	uSess.usersCookies = append(uSess.usersCookies, cookie)

	context.usersBySession = append(context.usersBySession, uSess)

	context.sessionIds = append(context.sessionIds, sess)

	redir := "/game/" + sess + "/"

	log.Printf("redirecting to %s", redir)
	http.Redirect(writer, request, redir, http.StatusSeeOther)

}

//gameandler is a wrapped version of the http.HandlerFunc
// which is extended to pass execution to the game function that
//does the heavy lifting for detecting an active game and making sure that
//each data stream is correctly routed.
// handy artical for the learner
//https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702
func gameHandler(context *appContext) http.Handler {
	return (http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			game(context, w, r)
		}))
}

// game matches on url /game/
func game(context *appContext, writer http.ResponseWriter, request *http.Request) {

	u, err := url.Parse(request.URL.Path)
	if err != nil {
		log.Print("error")
		log.Fatal(err)
	}
	s := u.Path
	log.Printf("path %s", u)

	validSession(isCorrect(context, s), writer, request)

}

func validSession(correct bool, writer http.ResponseWriter, request *http.Request) {

	if correct {
		log.Print("Game running...")
	} else {
		index := "/index"
		log.Printf("not  a valid session to %s", index)
		log.Printf("redirecting to %s", index)
		http.Redirect(writer, request, index, http.StatusSeeOther)
	}

}

func newUserCookie() http.Cookie {

	cookie := http.Cookie{
		Name:     "_gameCookie",
		Value:    GetUUID(),
		HttpOnly: true,
	}
	return cookie
}

func newGameSession() string {
	return RandomSession()
}

//isCorrect.
//Test is the session part of /game/blah is 16 char of A-z09
//And in the main session pool.
func isCorrect(context *appContext, path string) bool {

	//is a sequence of numbers and letters
	isAlphaNumericTest := regexp.MustCompile(`/game/([A-z0-9]+)/`).MatchString
	isAlphaNumeric := isAlphaNumericTest(path)

	bits := strings.Split(path, "/")
	sess := bits[2]
	log.Printf("bits %s", bits)
	log.Printf("session %s", sess)

	correctLength := len(sess) == 16

	realSession := false
	if len(context.sessionIds) > 0 {
		log.Print("application has sessions")

		for i := 0; i < len(context.sessionIds); i++ {
			if sess == context.sessionIds[i] {
				realSession = true
			}
		}
	}

	return isAlphaNumeric && realSession && correctLength
}
