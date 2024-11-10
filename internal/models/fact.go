package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"id"`
	Text  string `gorm:"not null" json:"text"`
	Count uint   `gorm:"default:0" json:"count"`
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
