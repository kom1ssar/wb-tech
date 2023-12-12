CREATE TABLE orders
(
    order_uid UUID PRIMARY KEY,
    track_number TEXT NOT NULL ,
    entry TEXT NOT NULL ,
    locale TEXT NOT NULL ,
    internal_signature TEXT NOT NULL  DEFAULT '',
    customer_id TEXT NOT NULL ,
    delivery_service TEXT NOT NULL,
    shardkey TEXT NOT NULL,
    sm_id INT NOT NULL ,
    date_created TIMESTAMPTZ NOT NULL ,
    oof_shard TEXT NOT NULL
);




CREATE TABLE delivery
(
    id SERIAL PRIMARY KEY,
    name VARCHAR (70) NOT NULL ,
    phone VARCHAR(11) NOT NULL ,
    zip VARCHAR(10) NOT NULL ,
    city VARCHAR(100) NOT NULL ,
    address TEXT NOT NULL ,
    region varchar(100) NOT NULL ,
    email varchar(150) NOT NULL,
    order_uid UUID NOT NULL ,
    CONSTRAINT fk_delivery_order FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);



CREATE TABLE  payment
(
    id SERIAL PRIMARY KEY,
    transaction TEXT NOT NULL UNIQUE ,
    request_id text NOT NULL  DEFAULT '',
    currency VARCHAR(30) NOT NULL,
    provider TEXT NOT NULL,
    amount  INT NOT NULL,
    payment_dt TIMESTAMP NOT NULL ,
    bank TEXT NOT NULL ,
    delivery_cost INT NOT NULL ,
    goods_total INT NOT NULL ,
    custom_fee INT NOT NULL,
    order_uid UUID NOT NULL,
    CONSTRAINT fk_payment_order FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);


CREATE TABLE item
(
    id SERIAL PRIMARY KEY,
    chrt_id INT NOT NULL,
    track_number TEXT NOT NULL,
    price  INT NOT NULL ,
    rid TEXT NOT NULL ,
    name TEXT NOT NULL,
    sale INT NOT NULL ,
    size TEXT NOT NULL ,
    total_price INT NOT NULL ,
    nm_id INT NOT NULL ,
    brand TEXT NOT NULL ,
    status INT NOT NULL
);





CREATE TABLE orders_item
(
    item_id  INT REFERENCES  item (id),
    order_uid UUID REFERENCES  orders(order_uid),
    CONSTRAINT orders_item_pkey PRIMARY KEY (item_id, order_uid)
);









INSERT INTO orders (order_uid, track_number, entry, internal_signature, locale, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES (
           'f0255086-28da-45b2-9433-7753ee76aafd','WBILMTESTTRACK', 'WBIL','', 'en', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1'
       );





INSERT INTO delivery (name, phone, zip, city, address, region, email, order_uid)
VALUES(
          'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com', 'f0255086-28da-45b2-9433-7753ee76aafd'
      );


INSERT INTO payment (transaction, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid)
VALUES (
           'b563feb7b2b84b6test', 'USD', 'wbpay', 1817, '2021-11-26T06:22:19Z', 'alpha', 1500,317, 0, 'f0255086-28da-45b2-9433-7753ee76aafd'
       );


INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES (
           9934930, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202
       );


INSERT INTO orders_item (item_id, order_uid)
VALUES(
          1, 'f0255086-28da-45b2-9433-7753ee76aafd'
      );



SELECT * FROM orders
                  LEFT JOIN payment
                            ON orders.order_uid = payment.order_uid

                  LEFT JOIN delivery
                            ON orders.order_uid = delivery.order_uid

                  LEFT JOIN  orders_item
                             ON orders.order_uid = orders_item.order_uid
                  LEFT JOIN item
                            ON orders_item.item_id = item.id

WHERE orders.order_uid = 'f0255086-28da-45b2-9433-7753ee76aafd'


DROP TABLE orders_item;
DROP TABLE payment;
DROP TABLE item;
DROP TABLE  delivery;
DROP TABLE  orders;


