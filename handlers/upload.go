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

func UploadFile(c *gin.Context) {
	minioClient, err := minio_client.GetClient(os.Getenv("MINIO_ENDPOINT"), os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"))
	if err != nil {
		panic(err)
	}

	if minioClient == nil {
		fmt.Println("Minio client is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Minio client"})
		return
	}

	// 1 MB limit
	if err := c.Request.ParseMultipartForm(1 << 20); err != nil {
		c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Failed to parse form or file is too large"))
		return
	}

	studentId := c.Request.FormValue("student_id")
	submission := c.Request.FormValue("submission")
	files := c.Request.MultipartForm.File["files"]
	fileNames := []string{}

	// studentId must be 7~9 character
	if len(studentId) < 7 || len(studentId) > 9 {
		c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Invalid student ID"))
		return
	}

	if studentId == "" || submission == "" {
		c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Student ID or submission is empty"))
		return
	}

	utils.RecordStudentIp(studentId, c.ClientIP())

	// Iterate through each uploaded file
	for _, fileHeader := range files {
		if !utils.CheckFileName(submission, fileHeader.Filename) {
			c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Invalid file name: "+fileHeader.Filename))
			return
		}

		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Failed to open file"))
			return
		}
		defer file.Close()

		file_path := studentId + "/" + submission + "/" + fileHeader.Filename

		_, err = minioClient.PutObject(context.TODO(), "final-exam", file_path, file, -1, minio.PutObjectOptions{})
		if err != nil {
			fmt.Println("Error uploading file:", err)
			c.Data(http.StatusAccepted, "text/html", utils.GetUploadErrorResponse("Failed to upload file to Minio"))
			return
		}

		fileNames = append(fileNames, fileHeader.Filename)
	}

	c.Data(http.StatusOK, "text/html", utils.GetUploadSuccessResponse(fileNames))
}

func WriteIpToFile(c *gin.Context) {
	utils.WriteIpToStudentIdToFile()
	c.JSON(http.StatusOK, gin.H{"message": "write ip to file successfully"})
}
