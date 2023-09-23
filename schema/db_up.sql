CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    first_name VARCHAR(225),
    last_name VARCHAR(225),
    email VARCHAR(254) UNIQUE,
    password_hash VARCHAR(500),
    phone_number VARCHAR(20),
    latitude NUMERIC(10, 7),
    longtitude NUMERIC(10, 7)
);

CREATE TABLE categories (
    id SERIAL NOT NULL UNIQUE
);

CREATE TABLE categories_translations (
    id SERIAL NOT NULL UNIQUE,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE NOT NULL,
    language_code VARCHAR(5) NOT NULL,
    name VARCHAR(225) NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE products (
    id SERIAL NOT NULL UNIQUE,
    rating NUMERIC(3, 2),
    category_id INT REFERENCES categories(id) ON DELETE CASCADE NOT NULL,
    image_urls TEXT[]
);

CREATE TABLE products_translations (
    id SERIAL NOT NULL UNIQUE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    language_code VARCHAR(5) NOT NULL,
    name VARCHAR(225) NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE consignments (
    id SERIAL NOT NULL UNIQUE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    expiration_date DATE,
    quantity INTEGER NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE discounts (
    id SERIAL NOT NULL UNIQUE,
    consignments_id INT REFERENCES consignments(id) ON DELETE CASCADE NOT NULL,
    name VARCHAR(225) NOT NULL,
    precent NUMERIC(5, 2) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE discounts_translations (
    id SERIAL NOT NULL UNIQUE,
    discount_id INT REFERENCES discounts(id) ON DELETE CASCADE NOT NULL,
    language_code VARCHAR(5) NOT NULL,
    name VARCHAR(225) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE clients_discounts (
    id SERIAL NOT NULL UNIQUE,
    client_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    consignment_id INT REFERENCES consignments(id) ON DELETE CASCADE NOT NULL,
    precent NUMERIC(5,2) NOT NULL
);

CREATE TABLE clients_discounts_translations (
    id SERIAL NOT NULL UNIQUE,
    clients_discount_id INT REFERENCES clients_discounts(id) ON DELETE CASCADE NOT NULL,
    language_code VARCHAR(5) NOT NULL,
    name VARCHAR(225) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE orders (
    id SERIAL NOT NULL UNIQUE,
    client_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    order_date TIMESTAMPTZ NOT NULL,
    total NUMERIC(10, 2) NOT NULL,
    status VARCHAR(25)
);

CREATE TABLE order_items (
    id SERIAL NOT NULL UNIQUE,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE NOT NULL,
    product_id INT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
    consignment_id INT REFERENCES consignments(id) ON DELETE CASCADE NOT NULL,
    quantity INT NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE payments (
    id SERIAL NOT NULL UNIQUE,
    payment_date TIMESTAMPTZ NOT NULL,
    total NUMERIC(10, 2) NOT NULL,
    payment_method VARCHAR(25) NOT NULL,
    status VARCHAR(25)
);

-- Inserting sample data into the 'categories' table
INSERT INTO categories DEFAULT VALUES;

-- Inserting sample data into the 'categories_translations' table
INSERT INTO categories_translations (category_id, language_code, name, description)
VALUES
    (1, 'en', 'Fruits', 'Fresh and delicious fruits'),
    (1, 'es', 'Frutas', 'Frutas frescas y deliciosas');

-- Inserting sample data into the 'products' table
INSERT INTO products (rating, category_id)
VALUES
    (4.5, 1),
    (3.8, 1),
    (4.2, 1),
    (3.6, 1),
    (4.8, 1),
    (3.9, 1),
    (4.3, 1),
    (4.1, 1),
    (3.7, 1),
    (4.6, 1);

-- Inserting sample data into the 'products_translations' table
INSERT INTO products_translations (product_id, language_code, name, description)
VALUES
    -- Product 1
    (1, 'en', 'Apples', 'Crisp and juicy apples'),
    (1, 'ru', 'Яблоки', 'Сочные и хрустящие яблоки'),

    -- Product 2
    (2, 'en', 'Oranges', 'Sweet and tangy oranges'),
    (2, 'ru', 'Апельсины', 'Сладкие и кислые апельсины'),

    -- Product 3
    (3, 'en', 'Bananas', 'Ripe and delicious bananas'),
    (3, 'ru', 'Бананы', 'Спелые и вкусные бананы'),

    -- Product 4
    (4, 'en', 'Grapes', 'Sweet and succulent grapes'),
    (4, 'ru', 'Виноград', 'Сладкий и сочный виноград'),

    -- Product 5
    (5, 'en', 'Peaches', 'Juicy and fragrant peaches'),
    (5, 'ru', 'Персики', 'Сочные и ароматные персики'),

    -- Product 6
    (6, 'en', 'Pears', 'Sweet and tender pears'),
    (6, 'ru', 'Груши', 'Сладкие и нежные груши'),

    -- Product 7
    (7, 'en', 'Berries Mix', 'A delightful mix of assorted berries'),
    (7, 'ru', 'Смесь Ягод', 'Восхитительная смесь разнообразных ягод'),

    -- Product 8
    (8, 'en', 'Mangoes', 'Exotic and refreshing mangoes'),
    (8, 'ru', 'Манго', 'Экзотические и освежающие манго'),

    -- Product 9
    (9, 'en', 'Watermelons', 'Cool and hydrating watermelons'),
    (9, 'ru', 'Арбузы', 'Охлаждающие и увлажняющие арбузы'),

    -- Product 10
    (10, 'en', 'Kiwi', 'Tangy and vitamin-packed kiwi'),
    (10, 'ru', 'Киви', 'Кислое и богатое витаминами киви');

-- Inserting sample data into the 'consignments' table
INSERT INTO consignments (product_id, expiration_date, quantity, price, description)
VALUES
    -- Consignment for Product 1
    (1, '2023-09-01', 100, 10.99, 'Fresh apples from the orchard'),

    -- Consignment for Product 2
    (2, '2023-08-31', 50, 7.99, 'Organic oranges from sunny groves'),

    -- Consignment for Product 3
    (3, '2023-09-05', 80, 4.49, 'Ripe bananas ready to eat'),

    -- Consignment for Product 4
    (4, '2023-09-10', 60, 8.99, 'Succulent grapes perfect for snacking'),

    -- Consignment for Product 5
    (5, '2023-08-29', 70, 6.99, 'Fragrant peaches bursting with flavor'),

    -- Consignment for Product 6
    (6, '2023-09-03', 40, 5.49, 'Tender pears for a sweet treat'),

    -- Consignment for Product 7
    (7, '2023-09-08', 90, 9.99, 'Assorted berries for a delightful mix'),

    -- Consignment for Product 8
    (8, '2023-08-30', 65, 7.49, 'Exotic mangoes with a refreshing taste'),

    -- Consignment for Product 9
    (9, '2023-09-06', 55, 3.99, 'Cool watermelons to beat the heat'),

    -- Consignment for Product 10
    (10, '2023-09-15', 75, 2.99, 'Tangy kiwi packed with vitamins');

-- Inserting sample data into the 'discounts' table
INSERT INTO discounts (consignments_id, name, precent, description)
VALUES
    (1, 'Summer Fruit Sale', 15.00, 'Special discount on fresh apples'),
    (2, 'Citrus Delight', 10.50, 'Discount on juicy oranges');

-- -- Inserting sample data into the 'clients_discounts' table
-- INSERT INTO clients_discounts (client_id, consignment_id, precent)
-- VALUES
--     (1, 1, 10.00),
--     (2, 2, 5.00);

-- -- Inserting sample data into the 'orders' table
-- INSERT INTO orders (client_id, order_date, total, status)
-- VALUES
--     (1, NOW(), 20.99, 'Pending'),
--     (2, NOW(), 15.49, 'Completed');

-- -- Inserting sample data into the 'order_items' table
-- INSERT INTO order_items (order_id, product_id, consignment_id, quantity, price)
-- VALUES
--     (1, 1, 1, 2, 10.99),
--     (2, 2, 2, 3, 7.99);

-- Inserting sample data into the 'payments' table
INSERT INTO payments (payment_date, total, payment_method, status)
VALUES
    (NOW(), 20.99, 'Credit Card', 'Success'),
    (NOW(), 15.49, 'PayPal', 'Success');


-- Inserting sample data into the 'discounts_translations' table
INSERT INTO discounts_translations (discount_id, language_code, name, description)
VALUES
    (1, 'en', 'Summer Sale', 'Special discounts for the summer season'),
    (1, 'es', 'Venta de Verano', 'Descuentos especiales para la temporada de verano');

-- -- Inserting sample data into the 'clients_discounts_translations' table
-- INSERT INTO clients_discounts_translations (clients_discount_id, language_code, name, description)
-- VALUES
--     (1, 'en', 'Client Discount', 'Discount for loyal customers'),
--     (1, 'es', 'Descuento para Clientes', 'Descuento para clientes leales');
