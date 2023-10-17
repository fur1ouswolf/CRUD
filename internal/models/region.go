package models

type Region struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
