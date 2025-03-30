package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/server/config"
	"www.github.com/NirajSalunke/server/helpers"
	"www.github.com/NirajSalunke/server/routes"
)

func init() {

	config.LoadEnv()
	config.Oauth()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Working EK NUMBERRRRR"})
	})

	routes.LoadRoutes(router)

	go hitter()

	router.Run(":" + os.Getenv("PORT"))
}

func hitter() {
	url := os.Getenv("BACKEND_URL") + "mail/"
	// fmt.Println(url)
	client := &http.Client{}
	helpers.PrintGreen("Mail Retrieving Automated")
	for {
		fmt.Println("Waiting for response...")
		startTime := time.Now()

		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error hitting the URL:", err)
		} else {
			if resp.StatusCode == http.StatusOK {
				fmt.Println("Got the response! Status Code:", resp.StatusCode)
			} else {

				body, _ := io.ReadAll(resp.Body)
				fmt.Printf("Error Response (Status: %d): %s\n", resp.StatusCode, string(body))
			}
			resp.Body.Close()
		}

		elapsed := time.Since(startTime)
		fmt.Printf("Request took: %v\n", elapsed)

		sleepTime := os.Getenv("SLEEP_TIME")
		fmt.Println("Wait period:", sleepTime, "seconds...")

		num, errInt := strconv.ParseInt(sleepTime, 10, 64)
		if errInt != nil || num < 10 {
			fmt.Println("Error converting env variable:", errInt)
			num = 50
		}

		time.Sleep(time.Duration(num) * time.Second)
	}
}
