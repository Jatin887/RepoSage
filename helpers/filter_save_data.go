package helpers
import(
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/initializers"
	log "github.com/sirupsen/logrus"
	"fmt"
)
func FilterAndSaveData(githubData string,userData models.GithubUserData,orgRepoData []models.Repo) {
	// Query the database for all existing GitHub data
	var existingData []models.Repo
	var uniqueData []models.Repo
	var userRepoData []models.UserRepo
	initializers.DB.Find(&existingData)

	// log the start of the function
	log.Println("Starting filterAndSaveData function")

	latestData := ConvertData(githubData,userData)
	for _, value := range orgRepoData {
		latestData = append(latestData, value)
	}
	if len(existingData) != 0 {
		existingMap := make(map[string]bool)
		// Iterate through the rows of data returned by the query
		for _, data := range existingData {
			// Add the existing data to the map
			existingMap[fmt.Sprintf("%d-%d", data.RepoID, data.OwnerID)] = true
		}
		for _, data := range latestData {
			// Check if the data already exists in the database
			if existingMap[fmt.Sprintf("%d-%d", data.RepoID, data.OwnerID)] {
				continue // Skip this data
			} else {
				result := models.UserRepo{
					UserID: data.OwnerID,
					RepoID: data.RepoID,
				}
				// Add the unique data to the slice
				uniqueData = append(uniqueData, data)
				userRepoData = append(userRepoData, result)
			}

		}

		if len(uniqueData) != 0 && len(userRepoData) != 0 {
			// log the number of unique data and user-repo relations being added
			log.Printf("Adding %d unique data and %d user-repo relations", len(uniqueData), len(userRepoData))

			initializers.DB.Create(&userRepoData)
			result := initializers.DB.Create(&uniqueData)
			if result.Error != nil {
				// log the error
				log.Fatal(result.Error)
			}
			// log the successful completion of the function
			log.Println("filterAndSaveData function completed successfully")
		} else {
			// log the exit of the function without adding any data
			log.Println("No unique data or user-repo relations to add")
			return
		}
	} else {
		for _, data := range latestData {
			result := models.UserRepo{
				UserID: data.OwnerID,
				RepoID: data.RepoID,
			}
			userRepoData = append(userRepoData, result)
			existingData = append(existingData, data)
		}
		initializers.DB.Create(&userRepoData)
		result := initializers.DB.Create(&existingData)
		if result.Error != nil {
			// log the error
			log.Fatal(result.Error)
		}
		// log the successful completion of the function
		log.Println("filterAndSaveData function completed successfully")
	}
}