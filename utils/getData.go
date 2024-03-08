package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetData() ([]string, []string, []string) {
	var timeDataStruct formatStruct
	var downloadDataStruct formatStruct
	var uploadDataStruct formatStruct
	timeBuffer, err := os.ReadFile("wsw-data/time.json")
	if err != nil {
		fmt.Print(err)
	}
	downloadBuffer, err := os.ReadFile("wsw-data/download.json")
	if err != nil {
		fmt.Print(err)
	}
	uploadBuffer, err := os.ReadFile("wsw-data/upload.json")
	if err != nil {
		fmt.Print(err)
	}
	timeStr := string(timeBuffer)
	downloadStr := string(downloadBuffer)
	uploadStr := string(uploadBuffer)
	json.Unmarshal([]byte(timeStr), &timeDataStruct)
	json.Unmarshal([]byte(downloadStr), &downloadDataStruct)
	json.Unmarshal([]byte(uploadStr), &uploadDataStruct)
	return timeDataStruct.Data, downloadDataStruct.Data, uploadDataStruct.Data
}
