package swagger

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mgo "examples/go-server-user_example/db"

	"gopkg.in/validator.v2"
)

/*
 * Create one User
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		mgo.Log.Printf("Error encoding %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + "Encoding error, " + err.Error() + `"}`))
		return
	}

	if err := validator.Validate(u); err != nil {
		mgo.Log.Printf("user validateion error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + "Invalid user data, " + err.Error() + `"}`))
		return
	}
	mgo.Log.Println("user validateion NO error")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := mgo.Collection.InsertOne(ctx, u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateUsersWithArrayInput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var users []User

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		mgo.Log.Printf("Error encoding %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + "Encoding error, " + err.Error() + `"}`))
		return
	}
	var ui []interface{}
	for _, u := range users {
		if err = validator.Validate(u); err != nil {
			mgo.Log.Printf("user validateion error: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message": "` + "Invalid user data, " + err.Error() + `"}`))
			return
		}
		ui = append(ui, u)
	}
	mgo.Log.Println("user validateion NO error")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := mgo.Collection.InsertMany(ctx, ui)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateUsersWithListInput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
 * Delete user by ID
 */
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := mgo.Collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func FindByAgeHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//var users []User
	users := []User{}
	hVals := r.Header
	filter := bson.M{}

	minAge, _ := strconv.Atoi(hVals.Get("X-Min-Age"))
	maxAge, _ := strconv.Atoi(hVals.Get("X-Max-Age"))
	if minAge > 0 {
		filter["$gte"] = minAge
	}
	if maxAge > 0 && maxAge > minAge {
		filter["$lte"] = maxAge
	}

	filter2 := bson.M{"age": filter}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := mgo.Collection.Find(ctx, filter2)
	defer cursor.Close(ctx)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "` + "Mongo connection error," + err.Error() + `"}`))
		return
	}
	for cursor.Next(ctx) {
		u := User{}
		cursor.Decode(&u)
		users = append(users, u)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "` + "Mongo Connection/query error, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func FindByAgeQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//var users []User
	users := []User{}
	qVals := r.URL.Query()
	filter := bson.M{}

	minAge, _ := strconv.Atoi(qVals["min"][0])
	maxAge, _ := strconv.Atoi(qVals["max"][0])
	if minAge > 0 {
		filter["$gte"] = minAge
	}
	if maxAge > 0 && maxAge > minAge {
		filter["$lte"] = maxAge
	}

	filter2 := bson.M{"age": filter}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := mgo.Collection.Find(ctx, filter2)
	defer cursor.Close(ctx)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "` + "Mongo connection error," + err.Error() + `"}`))
		return
	}
	for cursor.Next(ctx) {
		u := User{}
		cursor.Decode(&u)
		users = append(users, u)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "` + "Mongo Connection/query error, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

/*
 * Get user by ID by path param 'id'
 */
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + "Invalid User id, " + err.Error() + `"}`))
		return
	}
	var u User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mgo.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "` + "Mongo : user not found, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + "Invalid User id, " + err.Error() + `"}`))
		return
	}

	//update the existing user
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		mgo.Log.Printf("Input decode Error %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "` + "Input decode error, " + err.Error() + `"}`))
		return
	}
	tmpObj := bson.M{}
	//object to byte array
	tmpByteArray, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + "Data Marshall error, " + err.Error() + `"}`))
		return
	}

	//byte array to map
	err = json.Unmarshal(tmpByteArray, &tmpObj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + "Data UnMarshall error, " + err.Error() + `"}`))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": tmpObj}

	result, err := mgo.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "` + "Mongo connection error, " + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

/*
 * Get All users
 */
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//var users []User
	users := []User{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := mgo.Collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "` + "Mongo connection error," + err.Error() + `"}`))
		return
	}
	for cursor.Next(ctx) {
		u := User{}
		cursor.Decode(&u)
		users = append(users, u)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"message": "` + "Mongo Connection/query error, " + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
