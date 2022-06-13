package candystore

// Users represents a user, his favourite snack and the number of snacks eaten
type Users struct {
	NameUser   string           `json:"name"`
	FavSnack   string           `json:"favouriteSnack"`
	EatenSnack int64            `json:"totalSnacks"`
	HeapSnack  map[string]int64 `json:"-"` // Emulate a Max Heap, to get the most frequent snack
}
