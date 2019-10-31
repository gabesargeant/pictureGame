// main package for pictureGame
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type serverArgs struct {
	port string
}

type appContext struct {
	active     int
	sessionIds []string
}

func getServerArgs() serverArgs {
	port := flag.String("p", ":8080", "Port for the server, must include : prefix, ie :8080")
	flag.Parse()
	return serverArgs{port: *port}
}

func abortStartUp(err error) {
	fmt.Printf("There was an error with the value entered for a port: %s", err)
	panic(err)
}

func main() {

	args := getServerArgs()
	var err error
	args.port, err = CheckPort(args.port)

	if err != nil {
		abortStartUp(err)
	}

	cntxt := &appContext{}

	srv := startHTTPServer(args, cntxt)
	fmt.Printf("Starting on port %s \n", args.port)
	fmt.Println("Starting Picture Game Server")

	log.Fatal(srv.ListenAndServe())

	if err := srv.Shutdown(context.TODO()); err != nil {
		println(err)
		panic(err)
	}

	fmt.Println("Server off")
}

func setHandlers(mux *http.ServeMux, context *appContext) {

	//mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/index/", indexHandler)

	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/", http.StripPrefix("/static/", files))

	game := gameHandler(context)

	mux.Handle("/game/", game)

}

func startHTTPServer(args serverArgs, context *appContext) *http.Server {

	mux := http.NewServeMux()
	setHandlers(mux, context)

	srv := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         args.port,
		Handler:      mux}

	return srv
}
