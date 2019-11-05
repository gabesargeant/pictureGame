package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"regexp"
	"strings"
	"time"
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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

//RandomSession creates a random 16 letter and number string for
//representing a sharable game string
func RandomSession() string {
	//16 byte array
	b := make([]byte, 16)

	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)

}

//GetUUID returns a UUID
func GetUUID() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(uuid)
}
