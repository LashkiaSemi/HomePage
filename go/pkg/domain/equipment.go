package domain

// Equipment 備品
type Equipment struct {
	ID        int
	Name      string
	Stock     int
	Note      string
	CreatedAt string
	UpdatedAt string
	Tag       Tag
}

// Equipments 複数の備品
type Equipments []Equipment
