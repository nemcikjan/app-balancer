package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nemcikjan/app-balancer/k8s"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

type PostData struct {
	Id string `json:"id"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {

	// Ensure it's a POST request
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	var postData PostData
	// Read the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &postData); err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	// Convert the body to a string and print it
	fmt.Printf("Received: %s\n", string(body))

	res := k8s.CreatePodRequest(postData.Id)

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func main() {
	flag.Parse()
	http.HandleFunc("/create", handlePost)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal(err)
	}
}
