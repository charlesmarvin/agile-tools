package httpapi

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/charlesmarvin/agile-tools/domain"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// CreateBoardHandler creates new Board
func CreateBoardHandler(w http.ResponseWriter, req *http.Request) {
	var createBoardRequest domain.CreateBoardRequest
	err := json.NewDecoder(req.Body).Decode(&createBoardRequest)
	if err != nil {
		log.Printf("Error decoding request payload.")
		Error(w, http.StatusBadRequest, "Error decoding request payload.")
		return
	}
	newBoard, err := domain.CreateBoard(createBoardRequest)
	if err != nil {
		log.Println("Error decoding request payload.")
		Error(w, http.StatusBadRequest, "Error decoding request payload.")
	} else {
		Success(w, http.StatusCreated, newBoard)
	}
}

// GetBoardHandler gets Board by ID
func GetBoardHandler(w http.ResponseWriter, req *http.Request) {
	boardID := strings.TrimPrefix(req.URL.Path, "/api/v1/boards/")
	board := new(domain.Board)
	values, ok := req.URL.Query()["passcode"]
	var maybePasscode string
	if ok && len(values[0]) > 1 {
		maybePasscode = values[0]
	}
	err := board.Get(boardID, maybePasscode)
	if err != nil {
		if err == domain.ErrAuth {
			Error(w, http.StatusUnauthorized, "Board "+boardID+" does not exist.")
		} else {
			Error(w, http.StatusNotFound, "Board "+boardID+" does not exist.")
		}
	} else {
		Success(w, http.StatusOK, board)
	}
}

type command struct {
	Type     string            `json:"type"`
	Request  map[string]string `json:"request,omitempty"`
	Response interface{}       `json:"response,omitempty"`
	Error    string            `json:"error,omitempty"`
}
type NewSessionResponse struct {
	SessionID domain.SessionID `json:"sessionId"`
}
type JoinBoardResponse struct {
	BoardID domain.BoardID     `json:"boardId"`
	Members []domain.SessionID `json:"members"`
}

var sessionManager = NewSessionManager()

type SessionManager struct {
	ConnectionBySessionID map[domain.SessionID]*ConnectionManager
	sessionsByBoards      map[domain.BoardID][]domain.SessionID
	sessionBoardIndex     map[domain.SessionID]domain.BoardID
	votesByBoard          map[domain.BoardID]map[domain.SessionID]string
}

func NewSessionManager() *SessionManager {
	var sessionMgr = new(SessionManager)
	sessionMgr.ConnectionBySessionID = make(map[domain.SessionID]*ConnectionManager)
	sessionMgr.sessionsByBoards = make(map[domain.BoardID][]domain.SessionID)
	sessionMgr.sessionBoardIndex = make(map[domain.SessionID]domain.BoardID)
	sessionMgr.votesByBoard = make(map[domain.BoardID]map[domain.SessionID]string)
	return sessionMgr
}
func (sm *SessionManager) InitSession(conn net.Conn) domain.SessionID {
	sessionID := domain.NewSessionID()
	sm.ConnectionBySessionID[sessionID] = NewConnectionManager(conn)
	return sessionID
}

func (sm *SessionManager) GetConnectionManager(sessionID domain.SessionID) *ConnectionManager {
	return sm.ConnectionBySessionID[sessionID]
}

func (sm *SessionManager) JoinBoard(boardID domain.BoardID, sessionID domain.SessionID) {
	sm.sessionsByBoards[boardID] = append(sm.sessionsByBoards[boardID], sessionID)
	sm.sessionBoardIndex[sessionID] = boardID
	var newMemberMessage command
	newMemberMessage.Type = "member_joined"
	newMemberMessage.Response = map[string]interface{}{"sessionId": string(sessionID), "members": sm.sessionsByBoards[boardID]}
	sm.Broadcast(boardID, newMemberMessage)
}

func (sm *SessionManager) CastVote(boardID domain.BoardID, sessionID domain.SessionID, vote string) {
	if sm.votesByBoard[boardID] == nil {
		sm.votesByBoard[boardID] = make(map[domain.SessionID]string)
	}
	sm.votesByBoard[boardID][sessionID] = vote
	// reveal votes if all participants have casted their votes
	log.Printf("%v -- %v == %v", boardID, len(sm.sessionsByBoards[boardID]), len(sm.votesByBoard[boardID]))
	if len(sm.sessionsByBoards[boardID]) == len(sm.votesByBoard[boardID]) {
		var voteRevealMessage command
		voteRevealMessage.Type = "vote_reveal"
		voteRevealMessage.Response = map[string]interface{}{"votes": sm.votesByBoard[boardID]}
		sm.Broadcast(boardID, voteRevealMessage)
	}
}

func (sm *SessionManager) EndSession(sessionID domain.SessionID) {
	delete(sm.ConnectionBySessionID, sessionID)
	boardID := sm.sessionBoardIndex[sessionID]
	delete(sm.sessionBoardIndex, sessionID)
	sessions := sm.sessionsByBoards[boardID]
	for sessionIdx, sid := range sessions {
		if sessionID == sid {
			sessions[sessionIdx] = sessions[len(sessions)-1]
			sessions[len(sessions)-1] = sid
			sm.sessionsByBoards[boardID] = sessions[:len(sessions)-1]
			break
		}
	}
	var newMemberMessage command
	newMemberMessage.Type = "member_left"
	newMemberMessage.Response = map[string]interface{}{"sessionId": string(sessionID), "members": sm.sessionsByBoards[boardID]}
	sm.Broadcast(boardID, newMemberMessage)
}

func (sm *SessionManager) Broadcast(boardID domain.BoardID, message interface{}) {
	sessions := sm.sessionsByBoards[boardID]
	for _, sessionID := range sessions {
		connMgr := sm.ConnectionBySessionID[sessionID]
		connMgr.WriteJson(message)
	}
}

type ConnectionManager struct {
	Reader  *wsutil.Reader
	Writer  *wsutil.Writer
	Encoder *json.Encoder
	Decoder *json.Decoder
}

func NewConnectionManager(conn net.Conn) *ConnectionManager {
	var connMgr = new(ConnectionManager)
	connMgr.Reader = wsutil.NewReader(conn, ws.StateServerSide)
	connMgr.Writer = wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	connMgr.Encoder = json.NewEncoder(connMgr.Writer)
	connMgr.Decoder = json.NewDecoder(connMgr.Reader)
	return connMgr
}

func (connMgr *ConnectionManager) ReadJson(out interface{}) error {
	hdr, err := connMgr.Reader.NextFrame()
	if err != nil {
		return err
	}
	if hdr.OpCode == ws.OpClose {
		return err
	}
	if err := connMgr.Decoder.Decode(&out); err != nil {
		log.Printf("Error decoding request payload. %v", err)
		return err
	}
	log.Printf("Got message %v", out)
	return nil
}

func (connMgr *ConnectionManager) WriteJson(in interface{}) error {
	if err := connMgr.Encoder.Encode(&in); err != nil {
		log.Printf("Error encoding response payload. %v", err)
		return err
	}
	if err := connMgr.Writer.Flush(); err != nil {
		return err
	}
	return nil
}

// SocketHandler handles websocket connection
func SocketHandler(w http.ResponseWriter, req *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(req, w)
	if err != nil {
		// handle errors
	}

	go func() {
		defer conn.Close()
		sessionID := sessionManager.InitSession(conn)
		for {
			var cmd command
			if err := sessionManager.GetConnectionManager(sessionID).ReadJson(&cmd); err != nil {
				sessionManager.EndSession(sessionID)
				break
			}
			log.Printf("Got message %v", cmd)

			var response command
			switch cmd.Type {
			case "init":
				boardID := domain.BoardID(cmd.Request["boardId"])
				sessionManager.JoinBoard(boardID, sessionID)
				response.Type = cmd.Type
				response.Response = map[string]interface{}{"sessionId": string(sessionID), "members": sessionManager.sessionsByBoards[boardID]}
				log.Printf("Initialized new session: %v", response)
				sessionManager.GetConnectionManager(sessionID).WriteJson(response)
			case "vote":
				boardID := domain.BoardID(cmd.Request["boardId"])
				// sessionID := domain.SessionID(cmd.Request["sessionId"])
				vote := string(cmd.Request["vote"])
				log.Printf("New vote received: %v %v %v", boardID, sessionID, vote)
				sessionManager.CastVote(boardID, sessionID, vote)
			}
		}
	}()
}
