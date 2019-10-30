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

	srv := startHTTPServer(args)
	fmt.Printf("Starting on port %s \n", args.port)
	fmt.Println("Starting Cache Test Server")

	log.Fatal(srv.ListenAndServe())

	if err := srv.Shutdown(context.TODO()); err != nil {
		println(err)
		panic(err)
	}

	fmt.Println("Server off")
}

func setHandlers(mux *http.ServeMux) {

	//mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/index/", indexHandler)

	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/game/", game)
}

func startHTTPServer(args serverArgs) *http.Server {

	mux := http.NewServeMux()
	setHandlers(mux)

	srv := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         args.port,
		Handler:      mux}

	return srv
}
