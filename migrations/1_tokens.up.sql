CREATE TABLE tokens
(
    id         bigint AUTO_INCREMENT,
    name       VARCHAR(256) NOT NULL DEFAULT '',
    address    VARCHAR(256) NOT NULL DEFAULT '',
    symbol     VARCHAR(256) NOT NULL DEFAULT '',
    logo_url   VARCHAR(256) NOT NULL DEFAULT '',
    decimals   bigint       NOT NULL DEFAULT 0,
    created_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
