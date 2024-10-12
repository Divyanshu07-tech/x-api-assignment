package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

// Function to create a tweet
func createTweet(client *http.Client, tweet string) {
	apiURL := "https://api.twitter.com/1.1/statuses/update.json" // Twitter API v1.1 endpoint

	// Prepare the tweet data
	formData := url.Values{}
	formData.Set("status", tweet)

	// Create a new POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Read and handle the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error sending request: %s - Response: %s", resp.Status, string(body))
	}

	fmt.Println("Tweet posted:", string(body))
}

// Function to delete a tweet
func deleteTweet(client *http.Client, tweetID string) {
	apiURL := fmt.Sprintf("https://api.twitter.com/1.1/statuses/destroy/%s.json", tweetID) // API v1.1 delete endpoint

	// Create a new POST request for deletion
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response:", err)
		}
		log.Fatalf("Error deleting tweet: %s - Response: %s", resp.Status, string(body))
	}

	fmt.Println("Tweet deleted:", tweetID)
}

func main() {
	// API credentials
	apiKey := "oN8KBUX9mXKhn7bk33S7JVBv0"
	apiSecret := "eOcYVBQR6XXcAwzSNEqNJbe44FJ4Dfl5xbmJKA2ATWy830y0A9"
	accessToken := "1486025128232521729-L02vuLiVgm3ltx8bcRYmtjlRsu0OHQ"
	accessTokenSecret := "qLALe2sFc6LtynjFJzHVatQJA9Zsf6Y7M69l5DoKOGw04"

	// Create OAuth config and token
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	// Create a new HTTP client
	httpClient := config.Client(oauth1.NoContext, token)

	// Posting a tweet
	createTweet(httpClient, "Hello from Twitter API!")

	// Deleting a tweet (replace with actual tweet ID you want to delete)
	deleteTweet(httpClient, "your-tweet-id-here")
}
