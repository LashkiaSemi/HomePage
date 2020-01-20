package domain

// sessionList セッションの情報保持 map[studentID]sessionIDてな感じで
var sessionList map[string]string = map[string]string{}

type Session struct {
	SessionID string
	StudentID string
	UserID    int
}

func (s *Session) SetSessionList() {
	sessionList[s.StudentID] = s.SessionID
}

func GetSessionList() map[string]string {
	return sessionList
}
