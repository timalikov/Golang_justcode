package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var mySigningKey = []byte("secret")

// In-memory "database" for demonstration purposes
var usersDB = map[string]string{} // map username to password
var entitiesDB = map[string]string{
	"1": "Entity1",
	"2": "Entity2",
	"3": "Entity3",
}

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWTToken struct
type JWTToken struct {
	Token string `json:"token"`
}

// GenerateJWT generates a JWT token for a User
func GenerateJWT(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", fmt.Errorf("Something Went Wrong: %s", err.Error())
	}

	return tokenString, nil
}

// IsAuthorized is a middleware for checking the JWT token
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, "Error: %s", err.Error())
				return
			}

			if token.Valid {
				context := r.Context()
				context = context.WithValue(context, "user", token.Claims.(jwt.MapClaims)["user"])
				r = r.WithContext(context)
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// CreateTokenEndpoint creates a token for authenticated users
func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	// Create user in the "database"
	usersDB[user.Username] = user.Password

	tokenString, err := GenerateJWT(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	json.NewEncoder(w).Encode(JWTToken{Token: tokenString})
}

// GetEntitiesEndpoint returns a list of entities
func GetEntitiesEndpoint(w http.ResponseWriter, r *http.Request) {
	entities := make([]string, 0, len(entitiesDB))
	for _, entity := range entitiesDB {
		entities = append(entities, entity)
	}
	json.NewEncoder(w).Encode(entities)
}

// GetEntityEndpoint returns an entity by ID
func GetEntityEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entity, exists := entitiesDB[params["id"]]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Entity not found")
		return
	}
	fmt.Fprintf(w, "Entity with ID: %s is %s", params["id"], entity)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/auth", CreateTokenEndpoint).Methods("POST")
	router.Handle("/entities", IsAuthorized(GetEntitiesEndpoint)).Methods("GET")
	router.Handle("/entity/{id}", IsAuthorized(GetEntityEndpoint)).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))
}
