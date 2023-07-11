package helpers
import(
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func ConvertData(githubData string,userData models.GithubUserData) []models.Repo {
	var repoData []models.Repository
	data := []byte(githubData)
	var details models.Repo
	var finalData []models.Repo
	err := json.Unmarshal(data, &repoData)
	if err != nil {
		log.Printf("Error unmarshalling data: %v", err)
		return finalData
	}
	for _, user := range repoData {
		if !user.Private {
			details = models.Repo{
				OwnerID:    user.OwnerU.ID,
				OwnerName:  userData.Name,
				OwnerEmail: userData.Email,
				RepoID:     user.ID,
				RepoName:   user.Name,
				Status:     models.PublicAccess,
				Stars:      user.StargazersCount,
			}
		} else {
			details = models.Repo{
				OwnerID:    user.OwnerU.ID,
				OwnerName:  userData.Name,
				OwnerEmail: userData.Email,
				RepoID:     user.ID,
				RepoName:   user.Name,
				Status:     models.PublicAccess,
				Stars:      user.StargazersCount,
			}
		}
		finalData = append(finalData, details)
	}
	log.Printf("Successfully converted data with %d entries.", len(finalData))
	return finalData
}