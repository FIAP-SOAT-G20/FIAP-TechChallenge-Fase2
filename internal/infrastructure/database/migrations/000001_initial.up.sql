CREATE TABLE IF NOT EXISTS customers
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    email      VARCHAR   NOT NULL UNIQUE,
    cpf        VARCHAR   NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS categories
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS staffs
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    role VARCHAR CHECK (role IN ('COOK', 'ATTENDANT', 'MANAGER')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR        NOT NULL,
    description VARCHAR,
    price       DECIMAL(19, 2) NOT NULL,
    category_id INT            NOT NULL REFERENCES categories (id),
    image_url   VARCHAR,
    staff_id    INT REFERENCES staffs (id),
    active      BOOLEAN                 DEFAULT true,
    created_at  TIMESTAMP      NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP      NOT NULL DEFAULT now()
);

DROP TYPE IF EXISTS order_status;
CREATE TYPE order_status AS ENUM ('OPEN','CANCELLED','PENDING','RECEIVED', 'PREPARING', 'READY', 'COMPLETED');

CREATE TABLE IF NOT EXISTS orders
(
    id          SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers (id),
    status     order_status DEFAULT 'OPEN',
    total_bill  DECIMAL(19, 2), -- TODO: REMOVE
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS order_products
(
    order_id   INT REFERENCES orders (id),
    product_id INT REFERENCES products (id),
    quantity   INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (order_id, product_id)
);

CREATE TABLE IF NOT EXISTS order_histories
(
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders (id) NOT NULL,
    staff_id   INT REFERENCES staffs (id) NULL,
    status     order_status DEFAULT 'OPEN',
    created_at TIMESTAMP NOT NULL                                                        DEFAULT now(),
    updated_at TIMESTAMP NOT NULL                                                        DEFAULT now()
);

CREATE TABLE IF NOT EXISTS payments
(
    id                  SERIAL PRIMARY KEY,
    status              VARCHAR CHECK (status IN ('PROCESSING', 'CONFIRMED', 'ABORTED', 'FAILED')) DEFAULT 'PROCESSING',
    external_payment_id VARCHAR,
    order_id            INT REFERENCES orders (id),
    qr_data             VARCHAR,
    created_at          TIMESTAMP NOT NULL                                             DEFAULT now(),
    updated_at          TIMESTAMP NOT NULL                                             DEFAULT now()
);

INSERT INTO categories (name)
VALUES ('Lanches'),
       ('Bebidas'),
       ('Sobremesas'),
       ('Acompanhamentos'),
       ('Combos');

INSERT INTO staffs (name, role)
VALUES ('João Silva', 'COOK'),
       ('Maria Oliveira', 'ATTENDANT'),
       ('Pedro Santos', 'MANAGER'),
       ('Ana Costa', 'COOK'),
       ('Carlos Pereira', 'ATTENDANT');

INSERT INTO products (name, description, price, category_id, image_url, staff_id, active)
VALUES ('X-Burger', 'Hambúrguer com queijo, alface e tomate', 25.90, 1, 'https://example.com/xburger.jpg', 1, true),
       ('Coca-Cola 350ml', 'Refrigerante Coca-Cola lata', 6.90, 2, 'https://example.com/coca.jpg', 2, true),
       ('Sundae', 'Sorvete com calda de chocolate', 12.90, 3, 'https://example.com/sundae.jpg', 1, true),
       ('Batata Frita', 'Porção de batata frita crocante', 15.90, 4, 'https://example.com/batata.jpg', 4, true),
       ('Combo Big', 'X-Burger + Batata + Refrigerante', 42.90, 5, 'https://example.com/combo.jpg', 1, true);

INSERT INTO customers (name, email, cpf)
VALUES ('Lucas Mendes', 'lucas@email.com', '123.456.789-00'),
       ('Julia Santos', 'julia@email.com', '987.654.321-00'),
       ('Rafael Costa', 'rafael@email.com', '456.789.123-00'),
       ('Mariana Lima', 'mariana@email.com', '789.123.456-00'),
       ('Bruno Oliveira', 'bruno@email.com', '321.654.987-00');

INSERT INTO orders (customer_id, total_bill, status)
VALUES (1, 32.80, 'OPEN'),
       (2, 42.90, 'PENDING'),
       (3, 25.90, 'CANCELLED'),
       (4, 58.70, 'RECEIVED'),
       (5, 19.80, 'PREPARING'),
       (1, 25.90, 'READY'),
       (2, 6.90, 'COMPLETED');


INSERT INTO order_products (order_id, product_id, quantity)
VALUES (1, 1, 1),
       (1, 2, 1),
       (2, 5, 1),
       (3, 1, 1),
       (4, 1, 1),
       (4, 2, 1),
       (4, 3, 1),
       (4, 4, 1),
       (5, 2, 2),
       (5, 3, 1),
       (6, 1, 1),
       (7, 2, 1);


INSERT INTO order_histories (order_id, staff_id, status)
VALUES (1, null, 'OPEN'),
       (2, null, 'OPEN'),
       (2, null, 'PENDING'),
       (3, null, 'OPEN'),
       (3, null, 'PENDING'),
       (3, null, 'CANCELLED'),
       (4, null, 'OPEN'),
       (4, null, 'PENDING'),
       (4, null, 'RECEIVED'),
       (5, null, 'OPEN'),
       (5, null, 'PENDING'),
       (5, null, 'RECEIVED'),
       (5, 1, 'PREPARING'),
       (6, null, 'OPEN'),
       (6, null, 'PENDING'),
       (6, null, 'RECEIVED'),
       (6, 1, 'PREPARING'),
       (6, 2, 'READY'),
       (7, null, 'OPEN'),
       (7, null, 'PENDING'),
       (7, null, 'RECEIVED'),
       (7, 2, 'PREPARING'),
       (7, 2, 'READY'),
       (7, 2, 'COMPLETED');

INSERT INTO payments (id, status, external_payment_id, order_id, qr_data)
VALUES (1, 'PROCESSING', '09d92b11-cd55-4a72-b2ee-7377ceefe265', 2, 'QR_DATA_345'),
       (2, 'FAILED', 'b7fa4bee-fc25-4bb4-b948-5139af948a39', 3, 'QR_DATA_789'), 
       (3, 'CONFIRMED', '5c272292-4ba4-41e9-83d8-dea99afe5194', 4, 'QR_DATA_123'),
       (4, 'CONFIRMED', 'ac174c5e-c9ef-4407-a3b3-bceeb4163af3', 5, 'QR_DATA_456'),
       (5, 'CONFIRMED', '09d92b11-cd55-4a72-b2ee-7377ceefe265', 6, 'QR_DATA_345'),
       (6, 'CONFIRMED', '26e24f2a-5b00-4687-800f-a7be71104b2a', 7, 'QR_DATA_789');
       

SELECT setval('categories_id_seq', (SELECT MAX(id) FROM categories));
SELECT setval('staffs_id_seq', (SELECT MAX(id) FROM staffs));
SELECT setval('products_id_seq', (SELECT MAX(id) FROM products));
SELECT setval('customers_id_seq', (SELECT MAX(id) FROM customers));
SELECT setval('orders_id_seq', (SELECT MAX(id) FROM orders));
SELECT setval('order_histories_id_seq', (SELECT MAX(id) FROM order_histories));
SELECT setval('payments_id_seq', (SELECT MAX(id) FROM payments));
