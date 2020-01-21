package domain

// Company 就職先企業
type Company struct {
	ID        int
	Company   string
	CreatedAt string
	UpdatedAt string
}

// Companies 就職先企業
type Companies []Company
