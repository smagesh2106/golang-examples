package main

import (
	//"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"strconv"

	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"examples/ch10/mydriver"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname, omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname, omitempty"`
}

var client *mongo.Client
var collection *mongo.Collection
var logger *log.Logger

func CreatePersonEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	var person Person

	err := json.NewDecoder(request.Body).Decode(&person)
	fmt.Printf("firstname :%s, lastname :%s\n", person.Firstname, person.Lastname)
	if err != nil {
		fmt.Printf("Error encoding %v\n", err)
		//<FIXME>don't proceed after an error
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, person)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	var people []Person

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + "Mongo connection errro," + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + "Mongo Connection/query error, " + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func GetPersonEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	params := mux.Vars(request)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var person Person

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&person)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
			return
		} else {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
			return
		}
	}

	json.NewEncoder(response).Encode(person)
}

func DeletePersonEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	params := mux.Vars(request)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	//var person Person

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//result, err := collection.DeleteOne(ctx, Person{ID: id})
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func init() {
	logger = log.New(os.Stdout, "iam-policy-administration :", log.LstdFlags)
	logger.Println("mongo basics")

	tlsConfig := new(tls.Config)
	tlsConfig.InsecureSkipVerify = false

	key := "/home/magesh/02_Work/Openssl/client/cl.key"
	cacert := "/home/magesh/02_Work/Openssl/ca/ca.crt"
	cert := "/home/magesh/02_Work/Openssl/client/cl.crt"
	sub, err2 := mymongo.AddClientCertFromSeparateFiles(tlsConfig, key, cacert, cert, "")

	//sub, err2 := mymongo.AddClientCertFromConcatenatedFile(tlsConfig, keyCert, "")
	if err2 != nil {
		logger.Fatal(err2)
	} else {
		logger.Printf("certificates added...%v\n", sub)
	}
	logger.Printf("Verify Certificate : %s\n", strconv.FormatBool(!tlsConfig.InsecureSkipVerify))

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/?ssl=true").
		SetAuth(options.Credential{
			AuthSource: "magesh-mongo-test", Username: "user1", Password: "passw0rd",
		}).SetTLSConfig(tlsConfig) //.SetReplicaSet("rs0")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, _ = mongo.Connect(ctx, clientOptions)
	err1 := client.Ping(ctx, nil)
	if err1 != nil {
		logger.Printf(">> mongo connection error %v", err1)
	}

	collection = client.Database("magesh-mongo-test").Collection("people")
	logger.Println("done init ....")
}

func main() {
	router := mux.NewRouter()
	logger.Println("setting router....")
	router.HandleFunc("/person", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/person/{id}", DeletePersonEndPoint).Methods("DELETE")
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")

	logger.Println("Before start server ...")
	http.ListenAndServe(":8080", router)
	logger.Println("After start server ...")

}

//sudo mongod --tlsMode requireTLS --dbpath /var/lib/mongodb --auth --tlsCertificateKeyFile /home/magesh/02_Work/Openssl/server/server.pem --tlsCAFile /home/magesh/02_Work/Openssl/ca/ca.crt
//./mongobasic