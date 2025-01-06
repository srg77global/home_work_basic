-- name: InsertUser :one
INSERT INTO users (name, email, password)
VALUES ($1, $2, $3)
RETURNING id;

-- name: InsertProduct :one
INSERT INTO products (name, price)
VALUES ($1, $2)
RETURNING id;

-- name: UpdateUsernameByName :one
UPDATE users
SET name=$1
WHERE name=$2
RETURNING id;

-- name: UpdateProductpriceByName :one
UPDATE products
SET price=$1
WHERE name=$2
RETURNING id;

-- name: DeleteUserByName :one
DELETE FROM users
WHERE name=$1
RETURNING id;

-- name: DeleteProductByName :one
DELETE FROM products
WHERE name=$1
RETURNING id;


-- name: CreateOrderByUsernameProductnameAndQuantityFirst :one
INSERT INTO orders (user_id, order_date, total_amount)
VALUES ((SELECT id FROM users WHERE name=$1), now(), (SELECT price FROM products WHERE name=$2)*$3)
RETURNING $2;

-- name: CreateOrderByUsernameProductnameAndQuantitySecond :one
INSERT INTO orderProducts (order_id, product_id)
VALUES ((SELECT id FROM orders ORDER BY order_date DESC LIMIT 1), (SELECT id FROM products WHERE name=$1))
RETURNING order_id;

-- name: DeleteOrderByUsername :one
DELETE FROM orders
WHERE user_id=(SELECT id FROM users WHERE name=$1)
RETURNING $1;

-- name: SelectTwoUsersByUsername :many
SELECT * FROM users
WHERE name=$1 OR name=$2;

-- name: SelectProductsByPrices :many
SELECT name FROM products
WHERE price > $1 and price < $2;

-- name: SelectOrdersByUsername :many
SELECT u.name AS user_name, o.order_date, o.total_amount
FROM orders o, users u
WHERE o.user_id = u.id AND u.name=$1;

-- name: SelectUsersByOrders :many
SELECT u.name, o.total_amount / p.price AS units_bought
FROM users u, orders o, products p, orderProducts op
WHERE u.id = o.user_id AND o.id = op.order_id AND p.id = op.product_id;