package repo

import (
	"errors"
	"os"
)

var REPO_USERNAME string
var REPO_PASSWORD string
var REPO_URL string

func GetEnv() error {
	REPO_URL, exist := os.LookupEnv("REPO_URL")
	if !exist {
		return errors.New("repo url not defined")
	}

	REPO_USERNAME, exist := os.LookupEnv("REPO_USERNAME")
	if !exist {
		return errors.New("repo url not defined")
	}

	REPO_PASSWORD, exist := os.LookupEnv("REPO_PASSWORD")
	if !exist {
		return errors.New("repo url not defined")
	}

	return nil
}

func LoginRepo(repoUrl string) error {
	if REPO_USERNAME == "" || REPO_PASSWORD == "" {
		return errors.New("repo credentials not defined")
	}

	return nil
}

func LatestVersion(imageName string) {
	err := GetEnv()
	if err != nil {
		panic(err)
	}
}
