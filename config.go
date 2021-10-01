package buildagent

import (
	"errors"
	"os"
)

type RepoConfig struct {
	REPO_USERNAME string
	REPO_PASSWORD string
	REPO_URL      string
}

func GetRepoConfig() RepoConfig {
	config := new(RepoConfig)

	err := GetEnv(config)
	if err != nil {
		panic("Error in Repo config:" + err.Error())
	}

	return *config
}

func GetEnv(config *RepoConfig) error {
	url, exist := os.LookupEnv("REPO_URL")
	if !exist {
		return errors.New("repo url not defined")
	}

	user, exist := os.LookupEnv("REPO_USERNAME")
	if !exist {
		return errors.New("repo username not defined")
	}

	pass, exist := os.LookupEnv("REPO_PASSWORD")
	if !exist {
		return errors.New("repo password not defined")
	}

	config.REPO_URL = url
	config.REPO_USERNAME = user
	config.REPO_PASSWORD = pass

	return nil
}
