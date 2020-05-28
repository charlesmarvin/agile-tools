package models

// BoardType is a enumeration of board types [Fibonacci|Standard|Tshirt]
type BoardType string

// Board Model object for a board
type Board struct {
	Name     string    `bson:"name" json:"name"`
	Type     BoardType `bson:"type" json:"type"`
	ID       string    `bson:"_id" json:"id"`
	Passcode string    `bson:"passcode" json:"passcode"`
}

// ContactType is a enumeration of contact types [Phone|Email]
type ContactType string

// Contact represents either a phone or email contact
type Contact struct {
	Type  ContactType `bson:"contactType" json:"contactType"`
	Value string      `bson:"value" json:"value"`
}

// CreateBoardRequest is the input for creating a new board
type CreateBoardRequest struct {
	Name     string    `json:"name"`
	Type     BoardType `json:"type"`
	Passcode string    `json:"passcode"`
	Members  []Contact `json:"members"`
}
