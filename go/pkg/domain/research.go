package domain

// Research 卒業研究
type Research struct {
	ID        int
	Title     string
	Author    string
	File      string
	Comment   string
	CreatedAt string
	UpdatedAt string
}

// Researches 卒業研究
type Researches []Research
