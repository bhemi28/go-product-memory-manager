-- CREATE TABLE IF NOT EXISTS "products" (
--     "id" UUID PRIMARY KEY,
--     "name" VARCHAR(50) NOT NULL UNIQUE,
--     "description" VARCHAR(100) NOT NULL UNIQUE,
--     "original_price" FLOAT NOT NULL,
--     "link" VARCHAR(255) NOT NULL,
--     "website" VARCHAR(255) NOT NULL,
--     "category" VARCHAR(50) NOT NULL,
--     "estimated_price" FLOAT NOT NULL,
--     "notifications" BOOLEAN NOT NULL DEFAULT FALSE,
--     "availability" BOOLEAN NOT NULL DEFAULT FALSE,
--     "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
--     "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
-- );
--

-- name: CreateProduct: one
INSERT INTO products (id, name, description, original_price, link, website, category, estimated_price, notifications, availability)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;
