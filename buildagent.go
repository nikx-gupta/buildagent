package buildagent

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

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
		// var json Login
		// if err := c.ShouldBindJSON(&json); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }

		// if json.User != "manu" || json.Password != "123" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		// 	return
		// }

		// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	return r.Run()
}

func HandlePushEvent(event *GitEvent) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("No Home Directory: %s \n", err.Error())
		return
	}

	fmt.Printf("Temp Dir: %s \n", homeDir)
	cloneDir := homeDir + "/tmp/" + event.Repository.Name + "-" + uuid.NewString()
	fmt.Printf("Clone Dir: %s \n", cloneDir)

	CloneRepo(cloneDir, event.Repository.Url)
	files := FindDockerfiles(cloneDir)
	if len(files) == 0 {
		fmt.Printf("No Dockerfiles found")
		return
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

	fmt.Printf("Dockerfiles : %s", files)

	return files
}

func BuildImage(filePath string) {
	cmd := exec.Command("docker", "build", "-t", "")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error in Executing Git Clone: %s \n", err.Error())
	}

	cmd.Wait()
	fmt.Println("Git clone completed")

}

type GitRepository struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type GitEvent struct {
	Repository GitRepository `json:"repository"`
}
