CREATE TABLE products (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(255),
    description TEXT, 
    price NUMERIC(10,2)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    username TEXT NOT NULL, 
    email TEXT NOT NULL, 
    password TEXT NOT NULL
);

CREATE TABLE credit_cards (	
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id), 
    card_number CHARACTER VARYING(16) NOT NULL,
    expiry_month CHARACTER VARYING(2) NOT NULL, 
    expiry_year CHARACTER VARYING(4) NOT NULL,
    cvv CHARACTER VARYING(3) NOT NULL, 
    name_on_card CHARACTER VARYING(50) NOT NULL
);


-- PRODUCT TABLE
-- id: a serial primary key column that will automatically generate unique integer values for each row
-- name: a text column that stores the name
-- Description: a text column that stores the product description
-- price: a column that stores the price

-- USERS TABLE
-- id: a serial primary key column that will automatically generate unique integer values for each row
-- username: a text column that stores the user's username
-- email: a text column that stores the user's email address
-- password: a text column that stores the user's hashed password

-- CREDIT CARD TABLE
-- id: a serial primary key column that auto-increments
-- user_id: an integer column that references the id column in the users table
-- card_number: a character varying column that stores the credit card number and is not nullable
-- expiry_month: a character varying column that stores the expiry month of the credit card and is not nullable
-- expiry_year: a character varying column that stores the expiry year of the credit card and is not nullable
-- cvv: a character varying column that stores the CVV of the credit card and is not nullable
-- name_on_card: a character varying column that stores the name on the credit card and is not nullable
