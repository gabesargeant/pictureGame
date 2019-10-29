package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// CheckPort checks that a port entered is correctly defined.
// ie it should be numbers 8080, :8080, not :80:80
func CheckPort(port string) (string, error) {

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
