package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/charlesmarvin/agile-tools/domain"
	"github.com/charlesmarvin/agile-tools/models"
	"github.com/google/uuid"
)

// CreateBoardHandler creates new Board
func CreateBoardHandler(w http.ResponseWriter, req *http.Request) {
	var createBoardRequest models.CreateBoardRequest
	err := json.NewDecoder(req.Body).Decode(&createBoardRequest)
	if err != nil {
		log.Printf("Error decoding request payload.")
		Error(w, http.StatusBadRequest, "Error decoding request payload.")
		return
	}
	var newBoard models.Board
	// explicitly set ID
	newBoard.ID = uuid.New().String()
	newBoard.Name = createBoardRequest.Name
	newBoard.Type = createBoardRequest.Type
	newBoard.Passcode = createBoardRequest.Passcode
	_, err = domain.SaveBoard(newBoard)
	if err != nil {
		log.Printf("Error decoding request payload.")
		Error(w, http.StatusBadRequest, "Error decoding request payload.")
	} else {
		Success(w, http.StatusCreated, newBoard)
	}
}

// GetBoardHandler gets Board by ID
func GetBoardHandler(w http.ResponseWriter, req *http.Request) {
	boardID := strings.TrimPrefix(req.URL.Path, "/api/v1/boards/")
	board, err := domain.GetBoard(boardID)
	if err != nil {
		Error(w, http.StatusNotFound, "Board "+boardID+" does not exist.")
	} else {
		Success(w, http.StatusCreated, board)
	}
}
