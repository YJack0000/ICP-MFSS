package utils

import (
    "fmt"
    "os"
    "strings"
)

func GetUploadSuccessResponse(fileNames []string) []byte {
   content, err := os.ReadFile("static/components/upload_success.html")
	if err != nil {
		fmt.Println("could not read upload_success.html: %w", err)
    }

    successResponse := strings.ReplaceAll(string(content), "{.files}", strings.Join(fileNames, " "))

    return []byte(successResponse)
}

func GetUploadErrorResponse(error string) []byte {
    content, err := os.ReadFile("static/components/upload_error.html")
    if err != nil {
        fmt.Println("could not read upload_error.html: %w", err)
    }

    errorResponse := strings.ReplaceAll(string(content), "{.error}", error)
    return []byte(errorResponse)
}
