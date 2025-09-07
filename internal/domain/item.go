package domain

type Item struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null;default:''"`
}
