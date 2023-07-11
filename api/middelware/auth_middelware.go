package middelware
import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GithubAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		// Remove the "Bearer " prefix from the access token
		accessToken = accessToken[7:]

		// Call the GitHub API to validate the access token
		client := http.DefaultClient
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://api.github.com/user", nil)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
		resp, err := client.Do(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		respbody, _ := ioutil.ReadAll(resp.Body)

		// Convert byte slice to string and return
		var ghresp models.GithubUserData
		json.Unmarshal(respbody, &ghresp)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}

		// Set the authenticated user in the request context and call the next handler
		c.Set("userId", ghresp.ID)
		c.Next()
	}
}