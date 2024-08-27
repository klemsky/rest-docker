CREATE TABLE IF NOT EXISTS records (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    marks INTEGER[],
    created_at TIMESTAMPTZ
);

INSERT INTO records (name, marks, created_at) VALUES
('Alex', ARRAY[20, 30, 50, 40], '2024-08-26T08:30:00Z'),
('Brie', ARRAY[60, 40, 100, 90], '2024-08-26T08:15:00Z'),
('Clockie', ARRAY[100, 90, 80, 100], '2024-08-26T08:45:00Z'),
('Dome', ARRAY[10, 30, 20, 10], '2024-08-26T08:30:00Z'),
('Elaine', ARRAY[100, 100, 100, 100], '2024-08-26T08:00:00Z');