create table product (
                         id SERIAL primary key,
                         product_name VARCHAR(50) not null,
                         price NUMERIC(15,2) not NULL
)

    insert into product (product_name, price) values ('Sushi', 100.99)

select * from product p 