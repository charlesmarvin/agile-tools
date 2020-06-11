package domain

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// BoardType is a enumeration of board types [Fibonacci|Standard|Tshirt]
type BoardType string

// BoardID the board id
type BoardID string

// NewBoardID Creates a new random BoardID
func NewBoardID() BoardID {
	guid := strings.Replace(uuid.New().String(), "-", "", -1)
	guid = strings.ToUpper(guid)
	return BoardID("brd_" + guid)
}

// SessionID Session ID
type SessionID string

// NewSessionID Creates a new random sessionID
func NewSessionID() SessionID {
	guid := strings.Replace(uuid.New().String(), "-", "", -1)
	guid = strings.ToUpper(guid)
	return SessionID("sid_" + guid)
}

// Board Model object for a board
type Board struct {
	Name      string    `bson:"name" json:"name"`
	Type      BoardType `bson:"type" json:"type"`
	ID        BoardID   `bson:"_id" json:"id"`
	Shortcode string    `bson:"shortcode" json:"shortcode"`
	Passcode  string    `bson:"passcode" json:"passcode"`
	CreatedAt time.Time `bson:"created_on" json:"created_on"`
	CreatedBy AccountID `bson:"created_by" json:"created_by"`
}

// CreateBoardRequest is the input for creating a new board
type CreateBoardRequest struct {
	Name      string    `json:"name"`
	Type      BoardType `json:"type"`
	Passcode  string    `json:"passcode"`
	AccountID AccountID `json:"account_id"`
	Email     string    `json:"email"`
}

func generateRandomShortcode() string {
	// set a new see every time
	rand.Seed(time.Now().UnixNano())
	// generate a random number [0, 100,000,000)
	numericShortcode := rand.Intn(100000000)
	return fmt.Sprintf("%08d", numericShortcode)
}

// CreateBoard creates a new board from the CreateBoardRequest
func CreateBoard(createBoardRequest CreateBoardRequest) (*Board, error) {
	var newBoard Board
	// explicitly set ID
	newBoard.ID = NewBoardID()
	newBoard.Name = createBoardRequest.Name
	newBoard.Type = createBoardRequest.Type
	newBoard.Passcode = createBoardRequest.Passcode
	newBoard.Shortcode = generateRandomShortcode()
	newBoard.CreatedAt = time.Now().UTC()
	account := new(Account)
	if createBoardRequest.AccountID != "" {
		err := account.FindByID(createBoardRequest.AccountID)
		if err != nil {
			return nil, err
		}
	} else {
		err := account.FindByEmail(createBoardRequest.Email)
		if err != nil {
			if err == ErrNotFound {
				err = account.NewAccount(createBoardRequest.Email)
			}
			return nil, err
		}
	}
	newBoard.CreatedBy = account.ID
	newBoard.Save()
	return &newBoard, nil
}

// Save creates a new board
func (board *Board) Save() error {
	if board.ID == "" || board.Shortcode == "" {
		return ErrInvalidInput
	}
	// ignore the inserted ID since we are setting it explicitly
	_, err := boardCollection.InsertOne(context.Background(), board)
	return err
}

// Get gets a board by ID
func (board *Board) Get(id string, maybePasscode string) error {
	err := boardCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(board)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
	}
	if len(board.Passcode) > 1 && board.Passcode != maybePasscode {
		return ErrAuth
	}
	return err
}
