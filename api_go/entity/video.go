package entity

import (
    "time"
    "gorm.io/gorm"
)

type Video struct {
    ID          int           `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID      int           `gorm:"not null;index" json:"user_id"`
    Title       string         `gorm:"size:255;not null" json:"title"`
    Description string         `gorm:"type:text" json:"description"`
    URL         string         `gorm:"size:255;not null" json:"url"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
