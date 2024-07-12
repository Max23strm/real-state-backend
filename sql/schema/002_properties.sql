-- +goose Up
CREATE TABLE properties (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    type TEXT NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    bedrooms INT,
    bathrooms INT,
    price INT NOT NULL,
    square_meters INT NOT NULL,
    description TEXT NOT NULL,
    long_description VARCHAR(500) NOT NULL,
    location VARCHAR(18),
    is_available BOOLEAN NOT NULL,
    realtors VARCHAR(18) ARRAY[10],
    images VARCHAR(18) ARRAY[20]
);

-- +goose Down
DROP TABLE properties;