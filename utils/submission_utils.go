package utils

import (
	"fmt"
	"os"
	"strings"
)

func GetVaildSubmissionFileNames(submission string) []string {
	if submission == "lab12-2" {
		return []string{"main.c", "Makefile"}
	}

	if submission == "hw8-1" {
		return []string{
			"main.c",
			"browser.c",
			"browser.h",
			"webpage.c",
			"webpage.h",
			"webpage_stack.c",
			"webpage_stack.h",
			"Makefile",
		}
	}

    if submission == "final-makefile" {
        return []string {
            "Makefile",
        }
    }

	return []string{}
}

func CheckFileName(submission string, fileName string) bool {
	validFileNames := GetVaildSubmissionFileNames(submission)
	for _, validFileName := range validFileNames {
		if fileName == validFileName {
			return true
		}
	}
	return false
}

var studentIp = map[string][]string{}

func RecordStudentIp(studentId, ip string) {
	// If the studentId is not in the map, create a new list for it
	if _, ok := studentIp[studentId]; !ok {
		studentIp[studentId] = []string{}
	}

	// If the ip is not in the list, append it
	for _, v := range studentIp[studentId] {
		if v == ip {
			return
		}
	}

	studentIp[studentId] = append(studentIp[studentId], ip)
}

func WriteIpToStudentIdToFile() {
	file, err := os.OpenFile("static/ip_to_student_id.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("could not open ip_to_student_id.txt: %w", err)
	}
	defer file.Close()

	// Write the ip of each student, one student may have multiple ip
	// format: studentId, ip1, ip2, ip3
	for studentId, ips := range studentIp {
		_, err := file.WriteString(studentId + ",\"" + strings.Join(ips, ",") + "\"\n")
		if err != nil {
			fmt.Println("could not write ip_to_student_id.txt: %w", err)
		}
	}
}
