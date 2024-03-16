package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	r.GET("/status", func(c *gin.Context) {
		status := getStatus()
		c.JSON(http.StatusOK, status)
	})

	go updateJSON()

	r.Run(":8080")
}

func getStatus() Status {
	file, err := os.Open("status.json")
	if err != nil {
		fmt.Println("Error:", err)
		return Status{Water: 0, Wind: 0, WaterStatus: "unknown", WindStatus: "unknown"}
	}
	defer file.Close()

	var status Status
	if err := json.NewDecoder(file).Decode(&status); err != nil {
		fmt.Println("Error:", err)
		return Status{Water: 0, Wind: 0, WaterStatus: "unknown", WindStatus: "unknown"}
	}

	status.WaterStatus = getWaterStatus(status.Water)
	status.WindStatus = getWindStatus(status.Wind)

	return status
}

func getWaterStatus(water int) string {
	switch {
	case water < 5:
		return "aman"
	case water >= 6 && water <= 8:
		return "siaga"
	default:
		return "bahaya"
	}
}

func getWindStatus(wind int) string {
	switch {
	case wind < 6:
		return "aman"
	case wind >= 7 && wind <= 15:
		return "siaga"
	default:
		return "bahaya"
	}
}

func updateJSON() {
	for {
		water := rand.Intn(10) + 1
		wind := rand.Intn(10) + 1

		status := Status{Water: water, Wind: wind}
		writeJSON(status)

		time.Sleep(15 * time.Second)
	}
}

func writeJSON(status Status) {
	file, err := os.Create("status.json")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(status); err != nil {
		fmt.Println("Error:", err)
	}
}
