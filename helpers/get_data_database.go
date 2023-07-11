package helpers

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/initializers"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"

	log "github.com/sirupsen/logrus"
)

func GetReposForUser(userID int) ([]models.Repo, error) {
	var repos []models.Repo
	log.WithFields(log.Fields{"userID": userID}).Info("Retrieving repositories for user")
	err := initializers.DB.Joins("JOIN user_repos ON repos.repo_id = user_repos.repo_id").
		Where("user_repos.user_id = ?", userID).
		Find(&repos).Error
	if err != nil {
		log.WithError(err).WithFields(log.Fields{"userID": userID}).Error("Error retrieving repositories for user")
		return nil, err
	}
	log.WithFields(log.Fields{"userID": userID, "numRepos": len(repos)}).Info("Successfully retrieved repositories for user")
	return repos, nil
}