package main

import (
	"codegen-gotots/backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var inventory []models.Item

func loadJson() {
	data, err := os.ReadFile("data/inventory.json")
	if err != nil {
		log.Fatalf("cannot read json file: %v", err)
	}
	if err := json.Unmarshal(data, &inventory); err != nil {
		log.Fatalf("invalid json: %v", err)
	}
}

func main() {
	loadJson()

	http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(inventory)
	})
	// Handles Pagination by setting Limit as query param from Request
	//http://localhost:8080/getAIServers?limit=20
	http.HandleFunc("/getAIServers", func(w http.ResponseWriter, r *http.Request) {
		var limitQuery string = r.URL.Query().Get("limit")
		num, err := strconv.Atoi(limitQuery)
		if err != nil {
			fmt.Print(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		var gpuServers []models.Item
		for _, it := range inventory {
			if len(it.Gpus) > 0 {
				gpuServers = append(gpuServers, it)
			}
		}
		filteredResults := gpuServers[:num]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredResults)
	})
	//Gets server by GPU model
	http.HandleFunc("/getServerByGPU", func(w http.ResponseWriter, r *http.Request) {
		var gpuQuery string = r.URL.Query().Get("gpu")
		w.Header().Set("Content-Type", "application/json")
		var gpuServers []models.Item
		for _, it := range inventory {
			if len(it.Gpus) > 0 {
				for _, gpuModel := range it.Gpus {
					if strings.Contains(gpuModel, gpuQuery) {
						gpuServers = append(gpuServers, it)
					}
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gpuServers)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
