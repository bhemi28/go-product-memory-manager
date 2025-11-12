package types

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id                uuid.UUID      `json:"id"`
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	OriginalPrice     float64        `json:"original_price"`
	Link              string         `json:"link"`
	Website           string         `json:"website"`
	Category          string         `json:"category"`
	EstimatedPrice    float64        `json:"estimated_price"`
	SendNotifications bool           `json:"notifications"`
	Availability      bool           `json:"availability"`
	PriceHistory      []PriceHistory `json:"price_history"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

type ProductCreate struct {
	Id                uuid.UUID `json:"id"`
	Link              string    `json:"link"`
	Category          string    `json:"category"`
	EstimatedPrice    float64   `json:"estimated_price"`
	SendNotifications bool      `json:"notifications"`
}

type PriceHistory struct {
	Price float64   `json:"price"`
	Date  time.Time `json:"date"`
}
