package models

type Person struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Forename   string  `json:"forename" binding:"required"`
	Surname    string  `json:"surname" binding:"required"`
	Patronymic string  `json:"patronymic" binding:"required"`
	Gender     bool    `json:"gender" gorm:"default:true"`
	RegionID   *uint   `json:"region_id"`
	Region     *Region `gorm:"foreignKey:RegionID"`
}
