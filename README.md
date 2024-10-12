# Twitter API Assignment

## Introduction
This project demonstrates how to interact with the Twitter API using Go. It covers how to post a tweet and delete an existing tweet using OAuth 1.0a for authentication. The goal is to understand external API interactions, handle authentication, and manage API responses programmatically.

## Setup Instructions

### Step 1: Create a Twitter Developer Account and Generate API Keys
1. Visit the [Twitter Developer Platform](https://developer.twitter.com/).
2. Sign up for a developer account and create a new project.
3. Generate the following API keys from **Projects & Apps → Your App → Keys and Tokens**:
   - API Key
   - API Secret Key
   - Bearer Token
   - Access Token
   - Access Token Secret

### Step 2: Run the Go Program
1. Clone this repository.
2. Install dependencies:
   ```bash
   go get github.com/dghubble/oauth1
   ```
3. Set the environment variables:
   ```bash
    export API_KEY="oN8KBUX9mXKhn7bk33S7JVBv0"
    export API_SECRET_KEY="eOcYVBQR6XXcAwzSNEqNJbe44FJ4Dfl5xbmJKA2ATWy830y0A9"
    export ACCESS_TOKEN="1486025128232521729-L02vuLiVgm3ltx8bcRYmtjlRsu0OHQ"
    export ACCESS_TOKEN_SECRET="qLALe2sFc6LtynjFJzHVatQJA9Zsf6Y7M69l5DoKOGw04"
   ```
4. Run the program:
   ```bash
   go run main.go
   ```

## Program Details
## Posting a Tweet
- The program sends a `POST` request to the `https://api.twitter.com/1.1/statuses/update.json` endpoint.
- Example Tweet: "Hello from Twitter API with OAuth 1.0a!"
- To post a new tweet, modify the createTweet function with the text you want to tweet.

## Deleting a Tweet
- The program deletes a tweet using the `https://api.twitter.com/1.1/statuses/destroy/{tweet_id}.json `endpoint.
- To delete a tweet, provide the ID of the tweet in the `deleteTweet` function.
- Replace `your-tweet-id-here` with the actual ID of the tweet you want to delete.

## Error Handling
This program includes error handling for:

1. **Invalid API Credentials:** If the provided OAuth keys or tokens are incorrect, the program will output an error message.
2. **Invalid Tweet ID:** If the tweet ID you provide for deletion does not exist or is invalid, the program will inform you.
3. **Rate Limiting:** If you exceed the API rate limits, the program will notify you of the issue.

## Example Usage

### Example to Post a Tweet:
   ```bash
   curl -X POST "https://api.twitter.com/1.1/statuses/update.json?status=Hello from Twitter API!"
   ```

### Example to Delete a Tweet:
   ```bash
    curl -X POST "https://api.twitter.com/1.1/statuses/destroy/{tweet_id}.json"
   ```

## Conclusion
This project successfully demonstrates interacting with the Twitter API by posting and deleting tweets using Go. It also shows how to use OAuth 1.0a for secure authentication and handle various errors that might occur during API interactions.