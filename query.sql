CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE(product_id, user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'Pending',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- INDEXES
CREATE INDEX idx_category ON products USING btree(category_id);
CREATE INDEX idx_product_user_id ON cart_items USING btree(user_id);
CREATE INDEX idx_product_id ON cart_items USING btree(product_id);
CREATE INDEX idx_order_user_id ON orders USING btree(user_id);



-- Path: seed.sql
INSERT INTO categories (name) VALUES
('Electronics'),
('Clothing'),
('Books'),
('Furniture'),
('Toys');


INSERT INTO products (name, category_id, price) VALUES
('iPhone 12', 1, 799.99),
('Samsung Galaxy S21', 1, 699.99),
('Macbook Pro', 1, 1299.99),
('Dell XPS 13', 1, 999.99),
('Nike Air Max', 2, 99.99),
('Adidas Superstar', 2, 79.99),
('Levi''s Jeans', 2, 49.99),
('Ray-Ban Sunglasses', 2, 149.99),
('The Great Gatsby', 3, 9.99),
('To Kill a Mockingbird', 3, 7.99),
('1984', 3, 8.99),
('The Catcher in the Rye', 3, 6.99),
('Sofa', 4, 499.99),
('Dining Table', 4, 399.99),
('Office Chair', 4, 199.99),
('Bed', 4, 299.99),
('Lego Set', 5, 49.99),
('Barbie Doll', 5, 29.99),
('Hot Wheels', 5, 19.99),
('Nerf Gun', 5, 39.99);