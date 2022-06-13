package candystore

import (
	"encoding/json"
	"io/ioutil"
)

// ReportUsers creates a JSON file with the users and snacks
// Returns file content is bytes an error if the file cannot be created
func ReportUsers(users []*Users) ([]byte, error) {
	data, _ := json.Marshal(users)

	u := []Users{}
	if err := json.Unmarshal(data, &u); err != nil {
		return nil, err // handle error
	}

	// Write JSON file from RAW
	file, _ := json.MarshalIndent(u, "", "  ")
	if err := ioutil.WriteFile(ReportUsersFile, file, FileMode); err != nil {
		return nil, err // handle error
	}

	return file, nil
}
