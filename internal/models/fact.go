package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"text"`
	Content string `json:"content"`
	Teg     string `json:"teg"`
}

func (Fact) TableName() string {
	return "facts"
}

func (Fact) Relationships() map[string]string {
	return map[string]string{
		"ID":    "ID",
		"Count": "Count",
	}
}
