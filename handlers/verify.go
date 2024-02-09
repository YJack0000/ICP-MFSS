package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/minio/minio-go/v7"

	"github.com/YJack0000/ICP-MFSS/pkg/minio_client"
	"github.com/YJack0000/ICP-MFSS/utils"
)

func VerifyFile(c *gin.Context) {
	minioClient, err := minio_client.GetClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"))
	if err != nil {
		panic(err)
	}

	if minioClient == nil {
		fmt.Println("Minio client is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Minio client"})
		return
	}

	studentId := c.Request.FormValue("student_id")
	submission := c.Request.FormValue("submission")

	// studentId must be 9 character
	if len(studentId) != 9 {
		c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Invalid student ID"))
		return
	}

	if studentId == "" || submission == "" {
		c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Student ID or submission is empty"))
		return
	}

	// Check the files in the submission folder
	requiredFiles := utils.GetVaildSubmissionFileNames(submission)
	file_path := studentId + "/" + submission + "/"
	uploadedFiles := []string{}
	notUploadFiles := []string{}
	for _, requiredFile := range requiredFiles {
		// Iterate through each uploaded file in minio
		found := false
		for object := range minioClient.ListObjects(context.Background(), "final-exam", minio.ListObjectsOptions{Prefix: file_path, Recursive: true}) {
			file := object.Key
			if file == file_path+requiredFile {
				found = true
				break
			}
		}
		if found {
			uploadedFiles = append(uploadedFiles, requiredFile)
		} else if !found {
			notUploadFiles = append(notUploadFiles, requiredFile)
		}
	}

	c.Data(http.StatusOK, "text/html", utils.GetVerifyResponse(uploadedFiles, notUploadFiles))
}
