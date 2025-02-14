package main

import (
	"log"
	"net/http"
)

// token based bearer auth jwt
func authUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // implicit handler conversion
		log.Printf("authyy")
		// open up a connection to the user db
		// if bearer doesnt match user creds, die
		next.ServeHTTP(w, r)
	})
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
}

func listImages(w http.ResponseWriter, r *http.Request)    {}
func activateImage(w http.ResponseWriter, r *http.Request) {}
func delImage(w http.ResponseWriter, r *http.Request)      {}
func getImage(w http.ResponseWriter, r *http.Request) {
	log.Printf("get image from %s", r.PathValue("username"))
}

func main() {
	addr := "localhost:8080"

	//user
	http.HandleFunc("POST /users/{username}/images", uploadImage)
	http.HandleFunc("GET /users/{username}/images", listImages)
	http.HandleFunc("PUT /users/{username}/images/{image_id}/activate", activateImage)
	http.HandleFunc("DELETE /users/{username}/images/{image_id}", delImage)
	// func: del user

	//public
	http.HandleFunc("GET /{username}", getImage)
	// func: create user

	// middlware
	handler := authUser(http.DefaultServeMux)

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
