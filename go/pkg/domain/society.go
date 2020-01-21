package domain

// Society 学会発表
type Society struct {
	ID        int
	Title     string
	Author    string
	Society   string
	Award     string
	Date      string
	CreatedAt string
	UpdatedAt string
}

// Societies 学会発表。複数あるとき
type Societies []Society
