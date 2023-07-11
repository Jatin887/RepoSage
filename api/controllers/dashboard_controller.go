package controllers
import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)
func DashboardHandler(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	// Remove the "Bearer " prefix from the access token
	accessToken = accessToken[7:]
	helpers.GetGithubData(accessToken)
	user_id := c.MustGet("userId").(int)
	log.Info("Data filtered and saved")

	// Get repositories data for the user
	repoData, err := helpers.GetReposForUser(user_id)
	if err != nil {
		// Log the error
		log.Error("Error getting repositories for user")
		// Return an internal server error message
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Info("Repositories data retrieved for user")

	// Convert the data to JSON
	jsonString, err := json.Marshal(repoData)
	if err != nil {
		// Log the error
		log.Error("Error marshalling data to JSON")
		// Return an internal server error message
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Info("Data marshalled to JSON")

	data := string(jsonString)
	if data == "" {
		// Unauthorized users get an unauthorized message
		http.Error(c.Writer, "UNAUTHORIZED!", http.StatusUnauthorized)
		return
	}

	// Set return type JSON
	c.Writer.Header().Set("Content-type", "application/json")

	// Prettify the JSON
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(data), "", "\t")
	if err != nil {
		// Log the error
		log.Error("Error prettifying JSON")
		// Return an internal server error message
		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Info("JSON prettified")

	// Return the prettified JSON as a string
	fmt.Fprintf(c.Writer, string(prettyJSON.Bytes()))
}