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
