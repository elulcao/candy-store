package candystore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUsers(t *testing.T) {
	tests := []struct {
		data         [][]string
		expectedData string
	}{
		{
			data: [][]string{
				{"Name", "Candy", "Eaten"},
				{"Annika", "Geisha", "100"},
			},
			expectedData: `[
  {
    "name": "Annika",
    "favouriteSnack": "Geisha",
    "totalSnacks": 100
  }
]`,
		},
		{
			data: [][]string{
				{"Name", "Candy", "Eaten"},
				{"Annika", "Geisha", "100"},
				{"Jonas", "Kexchoklad", "10"},
				{"Jane", "Center", "1"},
			},
			expectedData: `[
  {
    "name": "Annika",
    "favouriteSnack": "Geisha",
    "totalSnacks": 100
  },
  {
    "name": "Jonas",
    "favouriteSnack": "Kexchoklad",
    "totalSnacks": 10
  },
  {
    "name": "Jane",
    "favouriteSnack": "Center",
    "totalSnacks": 1
  }
]`,
		},
		{
			data: [][]string{
				{"Name", "Candy", "Eaten"},
				{"Annika", "Geisha", "100"},
				{"Jonas", "Geisha", "200"},
				{"Jonas", "Kexchoklad", "100"},
				{"Aadya", "Nötchoklad", "2"},
				{"Jonas", "Nötchoklad", "3"},
				{"Jane", "Nötchoklad", "17"},
				{"Annika", "Geisha", "100"},
				{"Jonas", "Geisha", "700"},
				{"Jane", "Nötchoklad", "4"},
				{"Aadya", "Center", "7"},
				{"Jonas", "Geisha", "900"},
				{"Jane", "Nötchoklad", "1"},
				{"Jonas", "Kexchoklad", "12"},
				{"Jonas", "Plopp", "40"},
				{"Jonas", "Center", "27"},
				{"Aadya", "Center", "2"},
				{"Annika", "Center", "8"},
			},
			expectedData: `[
  {
    "name": "Jonas",
    "favouriteSnack": "Geisha",
    "totalSnacks": 1982
  },
  {
    "name": "Annika",
    "favouriteSnack": "Geisha",
    "totalSnacks": 208
  },
  {
    "name": "Jane",
    "favouriteSnack": "Nötchoklad",
    "totalSnacks": 22
  },
  {
    "name": "Aadya",
    "favouriteSnack": "Center",
    "totalSnacks": 11
  }
]`,
		},
	}

	for _, test := range tests {
		users, _ := parseUsers(test.data)
		sortedUsers := sortStoreUsers(users)
		reportUsers, _ := ReportUsers(sortedUsers)

		assert.Equal(t, test.expectedData, string(reportUsers))
	}
}
