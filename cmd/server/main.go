package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spinnaker/internal"
)

var cliPort = flag.Int("port", 8080, "Ignored if PORT is set as environmental variable")

func main() {
	flag.Parse()

	// TODO(ttomsu): Redirect to landing page with pretty graphs of all the data we get from this.
	// TODO(ttomsu): Make HandleGet private again after that.
	http.HandleFunc("/", internal.HandleGet)
	http.HandleFunc("/log", internal.LogEvent)

	port := 0
	envPort := os.Getenv("PORT")
	if envPort == "" {
		port = *cliPort
	} else {
		var err error
		if port, err = strconv.Atoi(envPort); err != nil {
			log.Printf("Cannot parse port from environment variable (value %v): %v", envPort, err)
			port = *cliPort
		}
	}

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
