package k8s

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/nemcikjan/app-balancer/utils"
)

type PodData struct {
	Cpu      int    `json:"cpu"`
	Memory   int    `json:"memory"`
	Priority int    `json:"priority"`
	Color    string `json:"color"`
	ExecTime int    `json:"execTime"`
}

type PodRequest struct {
	PodData
	Name string `json:"name"`
}

var eaodaUrl string = os.Getenv("EAODA_URL")

func CreatePodRequest(id string) []byte {
	podData := prepareData()
	request := PodRequest{
		Name:    id,
		PodData: *podData,
	}

	// Marshal the data into a JSON byte slice
	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", eaodaUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and log the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	return responseBody
}

func prepareData() *PodData {
	execTime := utils.RandomNumber(2, 100)
	var color string
	if execTime <= 10 {
		color = "red"
	} else if execTime <= 50 {
		color = "green"
	} else {
		color = "blue"
	}
	cpu := utils.RandomNumber(5, 200)
	memory := utils.RandomNumber(10, 400)
	prio := utils.RandomNumber(1, 5)
	return &PodData{Cpu: cpu, Memory: memory, Priority: prio, Color: color, ExecTime: execTime}
}
