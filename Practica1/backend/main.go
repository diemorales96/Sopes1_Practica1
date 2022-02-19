package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"practica1/server"

	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

var client *mongo.Client

func main() {

	s := server.New()
	Cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	Handler := Cors.Handler(s.Router())
	fmt.Println("Server UP on port: 4000")

	log.Fatal(http.ListenAndServe(":4000", Handler))

}
