package configs

const (
	JWTSecret         = "secret"
	JWTStudentIDClaim = "student_id"
	CookieName        = "_l_semi_homepage_session"
)

const (
	DateTimeFormat = "2006-01-02 15:4:5"
)

const (
	// TODO: 環境変数にした方がいい。絶対パスなら任意の場所に突っ込める
	SaveLectureFileDir = "./public/lectures"
)
