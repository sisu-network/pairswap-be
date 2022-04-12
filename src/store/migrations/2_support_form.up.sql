CREATE TABLE support_forms
(
    id         bigint AUTO_INCREMENT,
    name       VARCHAR(255)  NOT NULL DEFAULT '',
    email      VARCHAR(255)  NOT NULL DEFAULT '',
    tx_url     VARCHAR(255)  NOT NULL DEFAULT '',
    comment    VARCHAR(1023) NOT NULL DEFAULT '',
    created_at timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)
