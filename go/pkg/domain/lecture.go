package domain

// Lecture レクチャー
type Lecture struct {
	ID        int
	Title     string
	File      string
	Comment   string
	CreatedAt string
	UpdatedAt string
	User      User
}

// Lectures レクチャー
type Lectures []Lecture
