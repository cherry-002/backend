package data

import "time"

type Flower struct {
	ID        int64
	CreatedAt time.Time `json:"Created_at"`
	Name      string
	Origin    string		`json:",omitempty"` // will omit the field if its empty
	Details   string
	Price     uint64
	Version   int32		`json:"-"` // hides this info from user 
}
