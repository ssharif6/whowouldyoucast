package main

import (
	"os"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/ssharif6/whowouldyoucast/servers/gateway/models"
	"github.com/ssharif6/whowouldyoucast/servers/gateway/handlers"
	"net/http"
)

// TODO: Https
func main() {
	portAddress := os.Getenv("ADDR")
	if len(portAddress) == 0 {
		portAddress = "0.0.0.0:4000"
	}

	mongoAddr := os.Getenv("MONGO_ADDR")
	if len(mongoAddr) == 0 {
		log.Fatalf("MONGO_ADDR NOT SET")
	}

	mongoSess, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}

	reviewStore, err := models.NewMongoStore(mongoSess, "whowouldyoucast", "reviews")
	if err != nil {
		log.Fatalf(err.Error())
	}

	salt := os.Getenv("SALT")
	if len(salt) == 0 {
		log.Fatalf("Salt not set")
	}

	handlerCtx := handlers.NewHandlerCtx(salt, reviewStore)

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/reviews/review", handlerCtx.ReviewHandler)

	log.Printf("Listening on port %s\n", portAddress)
	log.Fatal(http.ListenAndServe(portAddress, mux))
}
