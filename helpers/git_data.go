package helpers

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func GetGithubData(accessToken string) {
	const maxRetries = 3
	const retryInterval = 200
	log.Trace("Getting Github data")

	// Get request to a set URL
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.WithError(reqerr).Error("API Request creation failed")
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request with retries
	resp, err := SendRequestWithRetries(req)
	if err != nil {
		log.WithError(err).Error("Request failed")
	}

	// Read the response as a byte slice
	respbody, _ := ioutil.ReadAll(resp.Body)

	// Convert byte slice to string and return
	var ghresp models.GithubUserData
	json.Unmarshal(respbody, &ghresp)

	url := ghresp.ReposURL

	log.Trace("Getting Github user repos")

	req_n, reqerr_n := http.NewRequest(
		"GET",
		url,
		nil,
	)

	if reqerr_n != nil {
		log.WithError(reqerr_n).Error("API Request creation failed")
	}

	// Make the request with retries
	resp_n, err := SendRequestWithRetries(req_n)
	if err != nil {
		log.WithError(err).Error("Request failed")
	}

	respbody_n, _ := ioutil.ReadAll(resp_n.Body)

	org_url := fmt.Sprintf("https://api.github.com/users/%s/orgs", ghresp.Login)

	log.Trace("Getting Github user orgs")

	req_o, reqerr_o := http.NewRequest(
		"GET",
		org_url,
		nil,
	)

	var temp []models.Repo

	if reqerr_o != nil {
		log.WithError(reqerr_o).Error("API Request creation failed")
	}

	req_o.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	// Make the request with retries
	resp_o, err := SendRequestWithRetries(req_o)
	if err != nil {
		log.WithError(err).Error("Request failed")

	}

	respbody_o, _ := ioutil.ReadAll(resp_o.Body)

	var orgRes []models.Org
	json.Unmarshal(respbody_o, &orgRes)

	for _, orgRepo := range orgRes {
		repoURL := orgRepo.ReposURL
		req_o, reqerr_o := http.NewRequest(
			"GET",
			repoURL,
			nil,
		)

		if reqerr_o != nil {
			log.WithError(reqerr_o).Error("API Request creation failed")
		}

		req_o.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		// Make the request with retries
		repo_r, err := SendRequestWithRetries(req_o)
		if err != nil {
			log.WithError(err).Error("Request failed")
			// Retry if there's an error
			for retryCount := 1; retryCount <= maxRetries; retryCount++ {
				log.WithField("retryCount", retryCount).Warn("Retrying request...")
				repo_r, err = SendRequestWithRetries(req_o)
				if err == nil {
					break
				}
				log.WithError(err).Error("Request failed")
				time.Sleep(retryInterval)
			}
			if err != nil {
				log.WithError(err).Error("Max retries exhausted. Request failed")
			}
		}
		respbody_r, _ := ioutil.ReadAll(repo_r.Body)

		var orgRepoData []models.OrgRepo
		json.Unmarshal(respbody_r, &orgRepoData)
		for _, entry := range orgRepoData {
			if entry.Private == false {
				result := models.Repo{
					OwnerID:    ghresp.ID,
					OwnerName:  ghresp.Name,
					OwnerEmail: ghresp.Email,
					RepoID:     int(entry.ID),
					RepoName:   entry.FullName,
					Status:     models.PublicAccess,
					Stars:      entry.StargazersCount,
				}
				temp = append(temp, result)
			} else {
				result := models.Repo{
					OwnerID:    ghresp.ID,
					OwnerName:  ghresp.Name,
					OwnerEmail: ghresp.Email,
					RepoID:     int(entry.ID),
					RepoName:   entry.FullName,
					Status:     models.PrivateAccess,
					Stars:      entry.StargazersCount,
				}
				temp = append(temp, result)
			}
		}
	}
	SaveUserData(ghresp)
	FilterAndSaveData(string(respbody_n), ghresp, temp)
}