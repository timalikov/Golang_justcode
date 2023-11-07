package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserProfile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var userProfiles = []UserProfile{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
}

func main() {
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("id")
		var profile UserProfile
		for _, p := range userProfiles {
			if p.ID == userID {
				profile = p
				break
			}
		}
		if profile.ID == "" {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(profile)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
