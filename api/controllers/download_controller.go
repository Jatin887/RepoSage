package controllers
import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/helpers"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
	"bytes"
	"encoding/csv"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func MarshalCSV(m models.Repo) ([]string, error) {
	record := []string{strconv.Itoa(m.OwnerID), m.OwnerName, m.OwnerEmail, strconv.Itoa(m.RepoID), m.RepoName, string(m.Status), strconv.Itoa(m.Stars)}
	for i, v := range record {
		record[i] = strconv.Quote(v)
	}
	return record, nil
}

func CsvHandler(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	// Remove the "Bearer " prefix from the access token
	accessToken = accessToken[7:]
	helpers.GetGithubData(accessToken)
	// Retrieve data from the database
	user_id := c.MustGet("userId").(int)
	log.Printf("USERID %s",user_id)
	var data []models.Repo
	// data, err := GetReposForUser(userID)
	data, err := helpers.GetReposForUser(user_id)
	if err != nil {
		log.Printf("Error retrieving data from the database: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Write data to CSV
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write header row
	header := []string{"OwnerID", "OwnerName", "OwnerEmail", "RepoID", "RepoName", "Status", "Stars"}
	err = writer.Write(header)
	if err != nil {
		log.Printf("Error writing header row to CSV: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Write data rows
	for _, d := range data {
		record, err := MarshalCSV(d)
		if err != nil {
			log.Printf("Error marshaling data to CSV: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = writer.Write(record)
		if err != nil {
			log.Printf("Error writing data row to CSV: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	writer.Flush()

	// Set HTTP headers and write response
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=data.csv")

	c.Writer.Write(buf.Bytes())
}