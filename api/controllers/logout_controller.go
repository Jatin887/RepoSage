package controllers

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleSignout(c *gin.Context) {
	// Start by logging the start of the function execution
	log.Println("HandleSignout function execution started")

	http.Redirect(c.Writer, c.Request, "https://github.com/logout", http.StatusTemporaryRedirect)

	// Log the end of the function execution
	log.Println("HandleSignout function execution completed")
}