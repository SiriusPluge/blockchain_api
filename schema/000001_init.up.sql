CREATE TABLE blockchain_list (
    id  SERIAL NOT NULL UNIQUE,
    symbol  VARCHAR(10) UNIQUE PRIMARY KEY,
    price_24h  NUMERIC,
    volume_24h NUMERIC,
    last_trade_price NUMERIC
);
