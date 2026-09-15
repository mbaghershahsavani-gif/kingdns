package models

type Service struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Enabled bool `json:"enabled"`
}
