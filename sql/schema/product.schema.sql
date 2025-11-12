-- type Product struct {
-- 	Id                uuid.UUID      `json:"id"`
-- 	Name              string         `json:"name"`
-- 	Description       string         `json:"description"`
-- 	OriginalPrice     float64        `json:"original_price"`
-- 	Link              string         `json:"link"`
-- 	Website           string         `json:"website"`
-- 	Category          string         `json:"category"`
-- 	EstimatedPrice    float64        `json:"estimated_price"`
-- 	SendNotifications bool           `json:"notifications"`
-- 	Availability      bool           `json:"availability"`
-- 	PriceHistory      []PriceHistory `json:"price_history"`
-- 	CreatedAt         time.Time      `json:"created_at"`
-- 	UpdatedAt         time.Time      `json:"updated_at"`
}

CREATE TABLE IF NOT EXISTS "products" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE,
    "description" VARCHAR(100) NOT NULL UNIQUE,
    "original_price" FLOAT NOT NULL,
    "link" VARCHAR(255) NOT NULL,
    "website" VARCHAR(255) NOT NULL,
    "category" VARCHAR(50) NOT NULL,
    "estimated_price" FLOAT NOT NULL,
    "notifications" BOOLEAN NOT NULL DEFAULT FALSE,
    "availability" BOOLEAN NOT NULL DEFAULT FALSE,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
