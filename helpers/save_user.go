package helpers

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/initializers"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
	log "github.com/sirupsen/logrus"
)

func SaveUserData(userData models.GithubUserData) {
	// log debug level message to indicate function entry
	log.WithFields(log.Fields{
		"function": "saveUserData",
	}).Debug("Entering saveUserData function")

	found := false
	var user []models.User
	initializers.DB.Find(&user)
	for _, n := range user {
		if n.ID == userData.ID {
			found = true
			break
		}
	}

	var filteredData []models.Repo
	initializers.DB.Find(&filteredData)

	if !found {
		// log info level message to indicate a new user entry is being created
		log.WithFields(log.Fields{
			"function": "saveUserData",
			"user_id":  userData.ID,
		}).Info("Creating new user entry")

		newEntry := models.User{
			ID:       userData.ID,
			Name:     userData.Name,
			Username: userData.Login,
			Email:    userData.Email,
		}
		initializers.DB.Create(&newEntry)

		// log debug level message to indicate successful user entry creation
		log.WithFields(log.Fields{
			"function": "saveUserData",
			"user_id":  userData.ID,
		}).Debug("User entry created successfully")
	}
	
	// log debug level message to indicate function exit
	log.WithFields(log.Fields{
		"function": "saveUserData",
	}).Debug("Exiting saveUserData function")
}