package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	ctx       = context.Background()
	moviesCol = getSession("movies")
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseMessage(w http.ResponseWriter, status int, results Message) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

/**
var movies = Movies{
	Movie{"Sin Limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 1999, "socrates"},
	Movie{"Rapidos", 2005, "Antonio"},
}*/

func getSession(collection string) *mongo.Collection {

	err := godotenv.Load("docker/.env")
	showError(err)

	//fmt.Printf("mongodb://%s:%s@%s:%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	showError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	showError(err)

	return client.Database(os.Getenv("DB_NAME")).Collection(collection)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Server")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Movie List \n")

	var results []Movie

	cursor, err := moviesCol.Find(ctx, bson.D{})

	showError(err)

	for cursor.Next(ctx) {

		var movie_data Movie
		err := cursor.Decode(&movie_data)
		showError(err)

		results = append(results, movie_data)
	}

	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	var oid primitive.ObjectID
	var err error
	var result Movie

	oid, err = primitive.ObjectIDFromHex(movie_id)
	showError(err)

	err = moviesCol.FindOne(ctx, bson.M{"_id": oid}).Decode(&result)

	log.Printf("Movie show %s", movie_id)
	responseMovie(w, 200, result)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)
	showError(err)

	defer r.Body.Close() //defer limpia lectura de body

	//log.Println(movie_data)
	_, err = moviesCol.InsertOne(ctx, movie_data)
	showErrorStatus(err, w)

	message := new(Message)
	message.setStatus("success")
	message.setMessage("The movie add: " + movie_data.Name)

	//movies = append(movies, movie_data)
	responseMessage(w, 200, *message)
}

func notFound(w http.ResponseWriter, r *http.Request) {

}

func showErrorStatus(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(500)
		panic(err.Error())
	}
}

func showError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
