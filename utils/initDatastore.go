package utils

import (
	"fmt"
	"os"
)

func InitDataStore() {
	fmt.Println("Initializing Datastore.")
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	timePath := currentDirectory + "/wsw-data/time.json"
	downloadPath := currentDirectory + "/wsw-data/download.json"
	uploadPath := currentDirectory + "/wsw-data/upload.json"

	if _, err := os.Stat(timePath); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir("wsw-data", os.ModePerm); err != nil {
				fmt.Println(err)
				return
			}
			f, err := os.Create("wsw-data/time.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			errr := os.WriteFile(timePath, []byte(`{"data":[]}`), 0644)
			if errr != nil {
				fmt.Println(errr)
				return
			}
		} else {
			fmt.Println(err)
		}
	}

	if _, err := os.Stat(downloadPath); err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create("wsw-data/download.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			errr := os.WriteFile(downloadPath, []byte(`{"data":[]}`), 0644)
			if errr != nil {
				fmt.Println(errr)
				return
			}
		} else {
			fmt.Println(err)
		}
	}

	if _, err := os.Stat(uploadPath); err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create("wsw-data/upload.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			errr := os.WriteFile(uploadPath, []byte(`{"data":[]}`), 0644)
			if errr != nil {
				fmt.Println(errr)
				return
			}
		} else {
			fmt.Println(err)
		}
	}
}
