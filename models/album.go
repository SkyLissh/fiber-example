package models

type Album struct {
	Base

	Title *string `gorm:"type:varchar(100);not null" json:"title"`
	Artist *string `gorm:"type:varchar(100);not null" json:"artist"`
	Price *float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}
