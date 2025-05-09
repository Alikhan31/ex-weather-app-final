package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
	City        string  `json:"city"`
}

var redisClient *redis.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "6379")),
		Password: "",
		DB:       0,
	})
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getWeatherFromAPI(lat, lon string) (*WeatherResponse, error) {
	return &WeatherResponse{
		Temperature: 20.5,
		Description: "Sunny",
		City:        "Example City",
	}, nil
}

func getWeather(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	if lat == "" || lon == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing coordinates"})
		return
	}
	cacheKey := fmt.Sprintf("weather:%s:%s", lat, lon)
	cachedWeather, err := redisClient.Get(c, cacheKey).Result()
	if err == nil {
		var weather WeatherResponse
		if err := json.Unmarshal([]byte(cachedWeather), &weather); err == nil {
			c.JSON(http.StatusOK, weather)
			return
		}
	}
	weather, err := getWeatherFromAPI(lat, lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get weather data"})
		return
	}
	weatherJSON, _ := json.Marshal(weather)
	redisClient.Set(c, cacheKey, weatherJSON, time.Hour)

	c.JSON(http.StatusOK, weather)
}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/weather", getWeather)
	log.Fatal(r.Run(":8080"))
}
