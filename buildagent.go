package buildagent

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Run() error {

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var json GitEvent
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		go HandlePushEvent(&json)
		c.Done()
	})

	r.POST("/loginJSON", func(c *gin.Context) {
	})

	return r.Run()
}

func HandlePushEvent(event *GitEvent) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("No Home Directory: %s \n", err.Error())
		return
	}

	fmt.Printf("Temp Dir: %s\n", homeDir)
	_, repoName := filepath.Split(event.Repository.Url)
	fmt.Printf("File Name: %s\n", repoName)

	cloneDir := homeDir + "/tmp/" + repoName + "-" + uuid.NewString()
	fmt.Printf("Clone Dir: %s \n", cloneDir)

	CloneRepo(cloneDir, event.Repository.Url)
	files := FindDockerfiles(cloneDir)
	if len(files) == 0 {
		fmt.Println("No Dockerfiles found")
		return
	}

	for _, dockerFile := range files {
		go BuildImage(repoName, dockerFile)
	}

}

func CloneRepo(cloneDirPath string, repoUrl string) {
	cmd := exec.Command("git", "clone", repoUrl, cloneDirPath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error in Executing Git Clone: %s \n", err.Error())
	}

	cmd.Wait()
	fmt.Println("Git clone completed")
}

func FindDockerfiles(cloneDir string) []string {
	files, _ := filepath.Glob(cloneDir + "/*/Dockerfile")

	fmt.Printf("Dockerfiles : %s\n", files)

	return files
}

var langs []string = []string{"golang", "dotnet", "java", "nodejs"}

func GetImagePrefix(repoName string, filePath string) string {
	_, fileName := filepath.Split(filePath)

	for _, lang := range langs {
		if strings.Contains(filePath, lang) {
			fileName = repoName + "-" + lang
			return fileName
		}
	}

	return fileName
}

func BuildImage(repoName string, filePath string) {
	imageName := GetImagePrefix(repoName, filePath)
	fmt.Printf("Image Name: %s\n", imageName)

	config := RepoConfig{}
	config.Get()

	repo := &DockerRepo{}
	err := repo.Login(config)
	if err != nil {
		fmt.Printf("Error in Repo Login: %s", err.Error())
		return
	}

	// cmd := exec.Command("docker", "build", "-t", "")
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Printf("Error in Executing Git Clone: %s \n", err.Error())
	// }

	// cmd.Wait()
	// fmt.Println("Git clone completed")

}

type GitRepository struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type GitEvent struct {
	Repository GitRepository `json:"repository"`
}
