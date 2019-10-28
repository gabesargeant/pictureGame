package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
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
	args.port, err = checkPort(args.port)

	if err != nil {
		abortStartUp(err)
	}

	setHandlers()

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

func setHandlers() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index/", indexHandler)
}

func startHTTPServer(args serverArgs) *http.Server {

	srv := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         args.port}

	return srv
}

func indexHandler(writer http.ResponseWriter, r *http.Request) {
	t := time.Now()
	t.Format(time.RFC822Z)
	fmt.Fprintf(writer, "The Time is....., %s!", t)
	fmt.Println("Fresh Server hit!")
}

/*
* CheckPort checks that a port entered is correctly defined.
* ie it should be numbers 8080, :8080, not :80:80
 */
func checkPort(port string) (string, error) {

	//cleanup whitespace
	port = strings.TrimSpace(port)

	re := regexp.MustCompile(":")
	matches := re.FindAllString(port, -1)

	//port entered without : adding :.
	if len(matches) == 0 {
		port = ":" + port
		fmt.Printf("adding : to port numbder %s \n", port)
		return port, nil
	}

	//port entered with 1 : correctly.
	if len(matches) == 1 {
		//confirming it's placement at 0 index
		if strings.IndexAny(port, ":") == 0 {
			return port, nil
		}
		if strings.IndexAny(port, ":") == 0 {
			return port, errors.New("Port Malformed, single : not at start")
		}
	}

	if len(matches) > 1 {
		return port, errors.New("Port Malformed ")
	}

	return port, nil

}
