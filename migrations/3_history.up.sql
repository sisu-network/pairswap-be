CREATE TABLE histories
(
  id            bigint AUTO_INCREMENT,
  address       VARCHAR(255)  NOT NULL DEFAULT '',
  recipient     VARCHAR(255)  NOT NULL DEFAULT '',
  src_chain     VARCHAR(255)  NOT NULL DEFAULT '',
  dest_chain    VARCHAR(255)  NOT NULL DEFAULT '',
  token_symbol  VARCHAR(255)  NOT NULL DEFAULT '',
  amount        VARCHAR(255)  NOT NULL DEFAULT '',
  src_hash      VARCHAR(255)  NOT NULL DEFAULT '',
  dest_hash     VARCHAR(255)  NOT NULL DEFAULT '',
  src_link      VARCHAR(255)  NOT NULL DEFAULT '',
  dest_link     VARCHAR(255)  NOT NULL DEFAULT '',
  created_at    timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)
