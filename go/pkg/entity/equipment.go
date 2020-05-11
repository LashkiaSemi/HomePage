package entity

// Equipment 備品
type Equipment struct {
	ID        int
	Name      string
	Stock     int
	Comment   string
	Tag       *Tag
	CreatedAt string
	UpdatedAt string
}
