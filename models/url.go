package models

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Url struct {
	Base

	Url *string `gorm:"type:varchar(10);not null" json:"url"`
	TargetUrl *string `gorm:"type:varchar(255);not null" json:"target_url"`
	IsActive *bool `gorm:"type:boolean;not null;default:true" json:"is_active"`
	Clicks *int `gorm:"type:integer;not null" json:"clicks"`
}

func (u *Url) BeforeCreate(tx *gorm.DB) (err error) {
	u.Url = u.randomUrl(10)

	return
}

func (u *Url) randomUrl(lenght int) *string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, lenght)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	url := string(b)

	return &url
}
