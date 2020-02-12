package domain

// sessionList セッションの情報保持 map[studentID]sessionIDてな感じで
var sessionList map[interface{}]*Session = map[interface{}]*Session{}

// Session セッションを管理
type Session struct {
	SessionID string
	StudentID string
	UserID    int
	Role      string
}

// SetSessionList サーバで保持しておくセッションリストを登録/更新
func (s *Session) SetSessionList() {
	sessionList[s.UserID] = s
}

// GetSessionByUserID userIDからセッションを取得
func GetSessionByUserID(userID int) (*Session, error) {
	sess := sessionList[userID]
	return sess, nil
}

// GetSessionList サーバで保持してる、ユーザのセッションを記録しているセッションリストを取得
func GetSessionList() map[interface{}]*Session {
	return sessionList
}
