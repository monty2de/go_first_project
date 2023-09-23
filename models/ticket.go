package models

type Ticket struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
	UserID  int    `json:"userID"` // Foreign key, refers to the User
}
