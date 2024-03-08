package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type formatStruct struct {
	Data []string `json:"data"`
}

func UpdateData(time string, download float64, upload float64) {
	var timeDataStruct formatStruct
	timeBuffer, err := os.ReadFile("wsw-data/time.json")
	if err != nil {
		fmt.Print(err)
	}
	xDataStr := string(timeBuffer)
	json.Unmarshal([]byte(xDataStr), &timeDataStruct)
	timeDataStruct.Data = append(timeDataStruct.Data, time)
	timeJson, err := json.Marshal(timeDataStruct)
	if err != nil {
		fmt.Print(err)
	}
	err = os.WriteFile("wsw-data/time.json", timeJson, 0644)
	if err != nil {
		panic(err)
	}
	var downloadDataStruct formatStruct
	downloadBuffer, err := os.ReadFile("wsw-data/download.json")
	if err != nil {
		fmt.Print(err)
	}
	downloadStr := string(downloadBuffer)
	json.Unmarshal([]byte(downloadStr), &downloadDataStruct)
	downloadDataStruct.Data = append(downloadDataStruct.Data, fmt.Sprint(download))
	downloadJson, err := json.Marshal(downloadDataStruct)
	if err != nil {
		fmt.Print(err)
	}
	err = os.WriteFile("wsw-data/download.json", downloadJson, 0644)
	if err != nil {
		panic(err)
	}
	var uploadDataStruct formatStruct
	uploadBuffer, err := os.ReadFile("wsw-data/upload.json")
	if err != nil {
		fmt.Print(err)
	}
	uploadStr := string(uploadBuffer)
	json.Unmarshal([]byte(uploadStr), &uploadDataStruct)
	uploadDataStruct.Data = append(uploadDataStruct.Data, fmt.Sprint(upload))
	uploadJson, err := json.Marshal(uploadDataStruct)
	if err != nil {
		fmt.Print(err)
	}
	err = os.WriteFile("wsw-data/upload.json", uploadJson, 0644)
	if err != nil {
		panic(err)
	}
}
