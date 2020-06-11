package domain

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AccountID User ID
type AccountID string

// NewAccountID Creates a new random AccountID
func NewAccountID() AccountID {
	guid := strings.Replace(uuid.New().String(), "-", "", -1)
	guid = strings.ToUpper(guid)
	return AccountID("acc_" + guid)
}

// Account represents either a phone or email contact
type Account struct {
	ID       AccountID `bson:"account_id" json:"user_id"`
	Email    string    `bson:"email" json:"email"`
	Verified bool      `bson:"verified" json:"value"`
}

// AccountVerification lookup used to verify contacts
type AccountVerification struct {
	AccountID    AccountID `bson:"account_id" json:"user_id"`
	Verification string    `bson:"verification" json:"verification"`
	CreatedAt    time.Time `bson:"created_on" json:"created_on"`
}

// NewAccount creates a new account with the given email. Assigns a randomly generated AccountID
func (account *Account) NewAccount(email string) error {
	account.ID = NewAccountID()
	account.Email = email
	account.Verified = true // FIXME default to true and kick off verification email here once infra is in place
	return account.Save()
}

// FindByID gets a board by ID
func (account *Account) FindByID(id AccountID) error {
	err := accountCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
	}
	return err
}

// FindByEmail gets a board by ID
func (account *Account) FindByEmail(email string) error {
	err := accountCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotFound
		}
	}
	return err
}

// Save creates a new board
func (account *Account) Save() error {
	if account.ID == "" || account.Email == "" {
		return ErrInvalidInput
	}
	_, err := accountCollection.InsertOne(context.Background(), account)
	return err
}
