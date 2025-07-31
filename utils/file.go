package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

var dataFiles = []string{
	"data/users.json",
	"data/rooms.json",
	"data/bookings.json",
	"data/tasks.json",
}

func EnsureDataFilesExist() {
	for _, file := range dataFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			f, err := os.Create(file)
			if err != nil {
				fmt.Println("Error creating", file, ":", err)
				continue
			}
			f.Write([]byte("[]")) // initialize as empty array
			f.Close()
		}
	}
}

func ReadJSON(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func WriteJSON(filePath string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}
