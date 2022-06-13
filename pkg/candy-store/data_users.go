package candystore

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"

	utils "github.com/elulcao/candy-store/pkg/utils"
)

// ReadUsersFile reads a CSV file and returns a slice of users, snacks and eaten snacks
// The CSV file must have the following columns: name, favourite snack, total snacks
// Returns an error if the file cannot be read
func ReadUsersFile(file string) ([]*Users, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err // handle error
	}

	users, err := parseUsers(records)
	if err != nil {
		return nil, err
	}

	return sortStoreUsers(users), nil
}

// parseUsers receives a slice of strings from CSV file and returns a map of users
// Returns an error if slice contains invalid data from the CSV file
func parseUsers(records [][]string) (map[string]*Users, error) {
	users := make(map[string]*Users, 0)

	for idx, record := range records {
		if idx == 0 { // skip header
			continue
		}

		user := record[0]
		snack := record[1]
		eaten, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err // invalid data
		}

		_, userExist := users[user]
		if !userExist { // Create map entry if it doesn't exist
			users[user] = &Users{}
			users[user].HeapSnack = make(map[string]int64, 0)
		}

		users[user].NameUser = user       // User name
		users[user].EatenSnack += eaten   // Number of snacks eaten
		users[user].HeapSnack[snack] += 1 // Snack occurence
		users[user].FavSnack = utils.MaxHeapString(users[user].HeapSnack)
	}

	return users, nil
}

// sortStoreUsers receives a map of users and returns a sorted slice of users by total snacks eaten
func sortStoreUsers(users map[string]*Users) []*Users {
	keys := make([]string, 0)
	newUsers := make([]*Users, 0)

	for key := range users { // Create a slice of keys for sorting the map
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool { // Sort by eaten snacks
		return users[keys[i]].EatenSnack > users[keys[j]].EatenSnack
	})
	for _, u := range keys { // Create new map with sorted users
		newUsers = append(newUsers, users[u])
	}

	return newUsers
}
