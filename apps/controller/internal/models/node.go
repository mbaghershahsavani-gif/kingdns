package models

type Node struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Country string `json:"country"`
	IPAddress string `json:"ip_address"`
	Status string `json:"status"`
}
