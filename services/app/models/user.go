package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"column:first_name;not null"`
	UserName  string `json:"username,omitempty" gorm:"column:username"`
	LastName  string `json:"last_name,omitempty" gorm:"column:last_name"`
	Photo     string `json:"photo,omitempty" gorm:"column:photo"`

	State    uint8  `json:"state,omitempty" gorm:"column:state;index:idx_state;default:0"`
	Status   bool   `json:"status,omitempty" gorm:"column:status;default:true"`
	Language string `json:"language,omitempty" gorm:"column:language"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
