package models

import (
	"time"
)

type Ad struct {
	ID string `json:"id"`

	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Image     string    `json:"image"`
	ExpiresAt time.Time `json:"expires_at"`
}
