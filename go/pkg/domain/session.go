package domain

// sessionList セッションの情報保持 map[studentID]sessionIDてな感じで
var sessionList map[interface{}]*Session = map[interface{}]*Session{}

type Session struct {
	SessionID string
	StudentID string
	UserID    int
	Role      string
}

func (s *Session) SetSessionList() {
	sessionList[s.UserID] = s
}

func GetSessionByUserID(userID int) (*Session, error) {
	sess := sessionList[userID]
	return sess, nil
}

func GetSessionList() map[interface{}]*Session {
	return sessionList
}
