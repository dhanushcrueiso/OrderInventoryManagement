CREATE TABLE customers (
    id uuid PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    "name" varchar not null,
    "mobile" varchar not null
);


CREATE TABLE users (
    id uuid PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    "name" varchar not null,
    "mobile" varchar not null,
    role varchar default 'executive'
);


CREATE TABLE products (
    id uuid PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,

);



CREATE TABLE inventory (
    id uuid PRIMARY KEY,
    product_id uuid REFERENCES products(id),
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


alter table inventory add column quantity INTEGER ;

CREATE TABLE orders (
    customer_id uuid not null,
    id uuid not null,
    product_id uuid not null,
    quantity_ordered INT,
    total_price DECIMAL(10, 2) NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (customer_id, id, product_id),
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

alter table orders add column order_id varchar not null;


create table "tokens" (
id uuid primary key not null,
	"token" varchar not null,
	account_id uuid not null,
	expires_at  TIMESTAMP WITHOUT TIME zone not null
	
	);
	
	
alter table "tokens" ADD CONSTRAINT token_constraint UNIQUE ("token");