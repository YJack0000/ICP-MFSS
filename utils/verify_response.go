package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetVerifyResponse(uploadedFileNames, notUploadedFileNames []string) []byte {
	content, err := os.ReadFile("static/components/verify_response.html")
	if err != nil {
		fmt.Println("could not read verify_response.html: %w", err)
	}

	successResponse := strings.ReplaceAll(string(content), "{.uploaded}", strings.Join(uploadedFileNames, "<br />"))

	successResponse = strings.ReplaceAll(successResponse, "{.not_uploaded}", strings.Join(notUploadedFileNames, "<br />"))

	return []byte(successResponse)
}
