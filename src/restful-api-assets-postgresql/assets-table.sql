CREATE DATABASE go-assets;

CREATE TABLE assets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    website VARCHAR(255),
    description TEXT
);

-- data demo
INSERT INTO assets (name, code, price, website, description)
VALUES
    ('Asset 1', 'CODE001', 10.99, 'https://example.com/asset1', 'Description for Asset 1'),
    ('Asset 2', 'CODE002', 19.99, 'https://example.com/asset2', 'Description for Asset 2'),
    ('Asset 3', 'CODE003', 15.99, 'https://example.com/asset3', 'Description for Asset 3'),
    ('Asset 4', 'CODE004', 12.49, 'https://example.com/asset4', 'Description for Asset 4'),
    ('Asset 5', 'CODE005', 8.99, 'https://example.com/asset5', 'Description for Asset 5'),
    ('Asset 6', 'CODE006', 14.99, 'https://example.com/asset6', 'Description for Asset 6'),
    ('Asset 7', 'CODE007', 9.99, 'https://example.com/asset7', 'Description for Asset 7'),
    ('Asset 8', 'CODE008', 11.99, 'https://example.com/asset8', 'Description for Asset 8'),
    ('Asset 9', 'CODE009', 13.49, 'https://example.com/asset9', 'Description for Asset 9'),
    ('Asset 10', 'CODE010', 17.99, 'https://example.com/asset10', 'Description for Asset 10'),
    -- Thêm các dòng dữ liệu cho 90 Asset còn lại
    -- ...
    ('Asset 99', 'CODE099', 8.99, 'https://example.com/asset99', 'Description for Asset 99'),
    ('Asset 100', 'CODE100', 12.99, 'https://example.com/asset100', 'Description for Asset 100');