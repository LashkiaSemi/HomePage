package domain

// Activity 活動記録
type Activity struct {
	ID        int
	Date      string
	Activity  string
	CreatedAt string
	UpdatedAt string
}

// Activities 複数の活動記録
type Activities []Activity
