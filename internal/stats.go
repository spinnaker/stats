// Package p contains an HTTP Cloud Function.
package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/logging"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/spinnaker/proto/stats"
)

const (
	ENV_SERVICE       = "K_SERVICE"
	ENV_REVISION      = "K_REVISION"
	ENV_CONFIGURATION = "K_CONFIGURATION"

	LOGGING_DELAY = 30 // seconds of delay before flushing any buffer.
)

const (
	DefaultLogID = "spinnaker-log-events-staging"
	ProdLogID    = "spinnaker-log-events-prod"

	LogIDEnvKey    = "LOG_ID_ENV"
	LogIdProdValue = "PROD"
)

var (
	projectID = os.Getenv("GCP_PROJECT")
	logID = DefaultLogID

	envVars   = []string{
		ENV_SERVICE,
		ENV_REVISION,
		ENV_CONFIGURATION,
	}

)

func init() {
	if (os.Getenv(LogIDEnvKey) == LogIdProdValue) {
		logID = ProdLogID
	}
}

func LogEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Received GET method for ", r.URL)
		handleGet(w, r)
		return
	case http.MethodPost:
		log.Println("Received POST method for ", r.URL)
		handlePost(w, r)
	default:
		http.Error(w, "405 - Method Not Allowed, punk!", http.StatusMethodNotAllowed)
	}
	fmt.Fprint(w, "Done.")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm healthy!\n")
	for _, key := range envVars {
		fmt.Fprintf(w, "%v: %v\n", key, os.Getenv(key))
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	event := &stats.Event{}
	um := &jsonpb.Unmarshaler{AllowUnknownFields: true}

	defer r.Body.Close()
	if err := um.Unmarshal(r.Body, event); err != nil {
		log.Printf("Error unmarshalling Event: %v", err)
		return
	}
	log.Printf("Unmarshaled:\n%+v\n", proto.MarshalTextString(event))

	client, err := logging.NewClient(r.Context(), projectID)
	if err != nil {
		log.Printf("could not create logging client: %v", err)
		return
	}

	logger := client.Logger(logID,
		logging.EntryCountThreshold(5),
		logging.DelayThreshold(time.Duration(LOGGING_DELAY)*time.Second))
	entry := logging.Entry{
		Payload:   event,
		Severity:  logging.Info,
		Timestamp: time.Now().UTC(),
	}
	logger.Log(entry)
}
