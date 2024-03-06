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
		Bookmark  []Bookmark
		// Bookmark  []Bookmark `gorm:"foreignKey:RedditorName;references:UserName"`
	}
	Session struct {
		gorm.Model
		UserSessionID string
	}

	Bookmark struct {
		Link          string `json:"link"`
		SubRedditName string `json:"subRedditName"`
		Title         string `json:"title"`
		UserID        uint
	}
)
