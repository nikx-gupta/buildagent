package buildagent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var token string

func LoginRepo(config RepoConfig) error {
	body, _ := json.Marshal(map[string]string{
		"username": config.REPO_USERNAME,
		"password": config.REPO_PASSWORD,
	})

	res, err := http.Post(config.REPO_URL+"/v2/users/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	var resBody map[string]string
	if res.StatusCode == 200 {
		json.NewDecoder(res.Body).Decode(&resBody)
		fmt.Printf("Authorization Response: %s", resBody)
		token = resBody["token"]
	}

	return nil
}

func LatestVersion(imageName string) {

}
