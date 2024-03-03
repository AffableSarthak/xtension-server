package models

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		UserName  string `gorm:"unique" json:"name,omitempty"`
		Hp        string `json:"hp,omitempty"`
		SessionID uint
		Session   Session
		Bookmark  []Bookmark `gorm:"foreignKey:RedditorName;references:UserName"`
	}

	Session struct {
		gorm.Model
		UserSessionID string
	}

	Bookmark struct {
		Link          string
		SubRedditName string
		Title         string
		RedditorName  string `gorm:"primaryKey"`
	}
)
