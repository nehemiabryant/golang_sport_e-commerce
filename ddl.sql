---------------------------------------------------------
-- 1. TABLE ROLES
---------------------------------------------------------
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Insert default roles
INSERT INTO roles (name) VALUES ('user'), ('admin');


---------------------------------------------------------
-- 2. TABLE USERS
---------------------------------------------------------
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE RESTRICT
);


---------------------------------------------------------
-- 3. TABLE CATEGORIES
---------------------------------------------------------
CREATE TABLE categories (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);


---------------------------------------------------------
-- 4. TABLE PRODUCTS
---------------------------------------------------------
CREATE TABLE products (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    price NUMERIC(12,2) NOT NULL,
    stock INTEGER NOT NULL
);


---------------------------------------------------------
-- 5. TABLE PRODUCT_CATEGORIES (many-to-many)
---------------------------------------------------------
CREATE TABLE product_categories (
    product_id VARCHAR(10) REFERENCES products(id) ON DELETE CASCADE,
    category_id VARCHAR(10) REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);


---------------------------------------------------------
-- 6. TABLE CART
---------------------------------------------------------
CREATE TABLE cart (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);


---------------------------------------------------------
-- 7. TABLE CART_ITEM
---------------------------------------------------------
CREATE TABLE cart_item (
    id SERIAL PRIMARY KEY,
    cart_id INTEGER REFERENCES cart(id) ON DELETE CASCADE,
    product_id VARCHAR(10) REFERENCES products(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL
);


---------------------------------------------------------
-- 8. TABLE PAYMENT
---------------------------------------------------------
CREATE TABLE payment (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    total_amount NUMERIC(14,2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW()
);


---------------------------------------------------------
-- 9. TABLE PAYMENT_DETAIL
---------------------------------------------------------
CREATE TABLE payment_detail (
    id SERIAL PRIMARY KEY,
    payment_id INTEGER REFERENCES payment(id) ON DELETE CASCADE,
    product_id VARCHAR(10) REFERENCES products(id) ON DELETE CASCADE,
    price NUMERIC(12,2) NOT NULL,
    subtotal NUMERIC(14,2) NOT NULL
);