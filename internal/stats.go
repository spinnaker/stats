// Package p contains an HTTP Cloud Function.
package internal

import (
	"context"
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
	StdOutLogID  = "spinnaker-stats-stdout"
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

	infoLogger *log.Logger
	debugLogger *log.Logger

	// The eventLogger is the logger that actually writes out the structured
	// data from the request payload.
	eventLogger *logging.Logger
)

// TODO(ttomsu): Consider wrapping this into a proper Server object and stop abusing package initialization.
// This would also allow a flag to dump stdout to a terminal during local development.
func init() {
	var loggingClient *logging.Client
	var err error
	if loggingClient, err = logging.NewClient(context.Background(), projectID); err != nil {
		log.Fatalf("could not create logging client: %v", err)
		return
	}

	stdOutLogger := loggingClient.Logger(StdOutLogID)
	infoLogger = stdOutLogger.StandardLogger(logging.Info)
	debugLogger = stdOutLogger.StandardLogger(logging.Debug)

	if os.Getenv(LogIDEnvKey) == LogIdProdValue {
		logID = ProdLogID
	}
	infoLogger.Printf("Using event log id: %v", logID)

	eventLogger = loggingClient.Logger(logID,
		logging.EntryCountThreshold(5),
		logging.DelayThreshold(time.Duration(LOGGING_DELAY)*time.Second))
}

func LogEvent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		HandleGet(w, r)
		return
	case http.MethodPost:
		debugLogger.Println("Received POST method for ", r.URL)
		handlePost(w, r)
	default:
		http.Error(w, "405 - Method Not Allowed, punk!", http.StatusMethodNotAllowed)
	}
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
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
		debugLogger.Printf("Error unmarshalling Event: %v", err)
		http.Error(w, "Bad input", http.StatusUnprocessableEntity)
		return
	}
	debugLogger.Printf("Unmarshaled:\n%+v\n", proto.MarshalTextString(event))

	entry := logging.Entry{
		Payload:   event,
		Severity:  logging.Info,
		Timestamp: time.Now().UTC(),
	}
	eventLogger.Log(entry)
	fmt.Fprint(w, "Done.")
}
