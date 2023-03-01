package store

type Metos struct {
	MetaID    int
	UserID    int
	PartOneID int
	PartTwoID int
	PartThrID int
	Key       string
}

type Parts struct {
	PartID int
	Data   []byte
	Node   string
}
