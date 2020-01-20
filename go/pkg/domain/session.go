package domain

// sessionList セッションの情報保持 map[studentID]sessionIDてな感じで
var sessionList map[interface{}]string = map[interface{}]string{}

type Session struct {
	SessionID string
	StudentID string
	UserID    int
}

func (s *Session) SetSessionList() {
	sessionList[s.UserID] = s.SessionID
}

func GetSessionList() map[interface{}]string {
	return sessionList
}
