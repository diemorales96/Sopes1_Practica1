package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:amarillo1234@35.238.28.35:27017/sopesPractica1?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/suma", a.Suma).Methods("POST")
	r.HandleFunc("/Operaciones", a.ObtenerOperaciones).Methods("GET")
	r.HandleFunc("/resta", a.Resta).Methods("POST")
	r.HandleFunc("/multiplicacion", a.Multiplicacion).Methods("POST")
	r.HandleFunc("/division", a.Division).Methods("POST")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) Suma(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
		return
	}
	var x Operaciones
	json.Unmarshal(reqBody, &x)
	n1, _ := strconv.ParseFloat(x.Numero1, 64)
	n2, _ := strconv.ParseFloat(x.Numero2, 64)
	x.Resultado = fmt.Sprintf("%f", n1+n2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)
	collection := client.Database("sopesPractica1").Collection("Operacion")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, x)
	json.NewEncoder(w).Encode(result)
}

func (a *api) Resta(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
		return
	}
	var x Operaciones
	json.Unmarshal(reqBody, &x)
	n1, _ := strconv.ParseFloat(x.Numero1, 64)
	n2, _ := strconv.ParseFloat(x.Numero2, 64)
	x.Resultado = fmt.Sprintf("%f", n1-n2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)
	collection := client.Database("sopesPractica1").Collection("Operacion")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, x)
	json.NewEncoder(w).Encode(result)
}

func (a *api) Multiplicacion(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
		return
	}
	var x Operaciones
	json.Unmarshal(reqBody, &x)
	n1, _ := strconv.ParseFloat(x.Numero1, 64)
	n2, _ := strconv.ParseFloat(x.Numero2, 64)
	x.Resultado = fmt.Sprintf("%f", n1*n2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)
	collection := client.Database("sopesPractica1").Collection("Operacion")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, x)
	json.NewEncoder(w).Encode(result)
}

func (a *api) Division(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
		return
	}
	var x Operaciones
	json.Unmarshal(reqBody, &x)
	n1, _ := strconv.ParseFloat(x.Numero1, 64)
	n2, _ := strconv.ParseFloat(x.Numero2, 64)
	x.Resultado = fmt.Sprintf("%f", n1/n2)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(x)
	collection := client.Database("sopesPractica1").Collection("Operacion")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, x)
	json.NewEncoder(w).Encode(result)
}

func (a *api) ObtenerOperaciones(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var x []Operaciones
	collection := client.Database("sopesPractica1").Collection("Operacion")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var temp Operaciones
		cursor.Decode(&temp)
		x = append(x, temp)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(x)
}
