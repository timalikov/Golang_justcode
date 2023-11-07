package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type JobPosting struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var jobPostings = []JobPosting{
	{ID: "1", Title: "Software Engineer", Description: "Developing and maintaining web applications."},
	{ID: "2", Title: "Product Manager", Description: "Leading product development from ideation to launch."},
}

func main() {
	http.HandleFunc("/posting", func(w http.ResponseWriter, r *http.Request) {
		jobID := r.URL.Query().Get("id")
		var posting JobPosting
		for _, p := range jobPostings {
			if p.ID == jobID {
				posting = p
				break
			}
		}
		if posting.ID == "" {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(posting)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
