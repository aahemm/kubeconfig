package actions

import "os"

func areFilesEqual(fileOnePath, fileTwoPath string) bool {
	fileOneData, _ := os.ReadFile(fileOnePath)
	fileTwoData, _ := os.ReadFile(fileTwoPath)
	return string(fileOneData) == string(fileTwoData)
}