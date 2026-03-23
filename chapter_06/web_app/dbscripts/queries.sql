create table products (
    id int unique,
    product_name varchar(255) not null,
    color varchar(255),
    price int default 5
);