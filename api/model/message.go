package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID `gorm:"type:uuid; default:gen_random_uuid()" json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
