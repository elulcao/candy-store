# Usage

## Build

```bash
make build
```

## Run

```bash
./candy-store --file test.csv
```

## Test

```bash
make test
```

# Summary

Input is a CSV file with the following columns:
* `Name`: The name of the user
* `Candy`: The name of the candy
* `Eaten`: The number of candies eaten

**Note:** The first line of the CVS is ignored, assuming it is a header.

Ouput is a JSON file with the user's name and the number of candies eaten, sorted by number of candies eaten.

```json
[
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
]
```

**Note:** The output is a JSON output, not a JavaScript JSON object.