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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    order_code VARCHAR(50) NOT NULL,
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
CREATE INDEX idx_order_id ON order_items USING btree(order_id);
CREATE INDEX idx_order_product_id ON order_items USING btree(product_id);
CREATE INDEX idx_order_code ON orders USING btree(order_code);



-- Path: seed.sql
INSERT INTO categories (name) VALUES
('Electronics'),
('Clothing'),
('Books'),
('Furniture'),
('Toys');


INSERT INTO products (name, category_id, price) 
VALUES
('iPhone 12', 1, 10000000.00),
('Samsung Galaxy S21', 1, 9000000.00),
('Macbook Pro', 1, 20000000.00),
('Dell XPS 15', 1, 15000000.00),
('Nike Air Max', 2, 500000.00),
('Adidas Superstar', 2, 400000.00),
('Levi''s Jeans', 2, 300000.00),
('H&M T-shirt', 2, 200000.00),
('The Alchemist', 3, 100000.00),
('Harry Potter', 3, 150000.00),
('The Da Vinci Code', 3, 120000.00),
('The Great Gatsby', 3, 110000.00),
('Sofa', 4, 3000000.00),
('Dining Table', 4, 2500000.00),
('Bed', 4, 2000000.00),
('Wardrobe', 4, 1500000.00),
('Lego', 5, 1000000.00),
('Barbie', 5, 800000.00),
('Hot Wheels', 5, 700000.00),
('Nerf', 5, 600000.00);