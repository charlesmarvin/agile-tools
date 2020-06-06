package domain

import (
	"log"
	"strings"

	"github.com/google/uuid"
)

// SessionID Session ID
type SessionID string

// NewSessionID Creates a new random sessionID
func NewSessionID() SessionID {
	guid := strings.Replace(uuid.New().String(), "-", "", -1)
	guid = strings.ToUpper(guid)
	return SessionID("sid_" + guid)
}

// PlanningSession Keeps track of Planning Sessions
type PlanningSession struct {
	boardSessions map[BoardID][]SessionID
	sessionBoard  map[SessionID]BoardID
}

// NewPlanningSession Creates a new PlanningSession struct
func NewPlanningSession() *PlanningSession {
	var ps PlanningSession
	ps.boardSessions = make(map[BoardID][]SessionID)
	ps.sessionBoard = make(map[SessionID]BoardID)
	return &ps
}

// Init Initializes a new planning session for  the given board
func (ps PlanningSession) Init(boardID BoardID) SessionID {
	var sessionID = NewSessionID()
	ps.boardSessions[boardID] = append(ps.boardSessions[boardID], sessionID)
	ps.sessionBoard[sessionID] = boardID
	return SessionID(sessionID)
}

// Close Closes and removes a planning session
func (ps PlanningSession) Close(sessionID SessionID) {
	log.Printf("Removing session %s", sessionID)
	delete(ps.sessionBoard, sessionID)
}
