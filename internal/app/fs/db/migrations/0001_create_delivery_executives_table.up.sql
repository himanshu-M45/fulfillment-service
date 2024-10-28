-- db/migrations/0001_create_delivery_executives_table.up.sql
CREATE TABLE IF NOT EXISTS delivery_executives (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255),
    is_available BOOLEAN,
    assigned_order_id INT
);