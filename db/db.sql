/*create tables and Fks*/

CREATE TABLE `db`.`accounts`  (
    `account_id` int NOT NULL AUTO_INCREMENT,
    `account_number` int(10) NOT NULL,
    `password` varchar(255) NOT NULL,
    `token` varchar(255) NULL,
    `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`account_id`),
    UNIQUE INDEX `idx_account_token`(`token`) USING BTREE
);

CREATE TABLE `db`.`account_balance`  (
  `account_balance_id` int NOT NULL AUTO_INCREMENT,
  `account_id` int NULL,
  `balance` float(255, 2) NULL,
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`account_balance_id`),
  INDEX `idx_account_balance_balance`(`balance`) USING BTREE,
  CONSTRAINT `fk_balance_account_id` FOREIGN KEY (`account_id`) REFERENCES `db`.`accounts` (`account_id`) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE `db`.`transactions`  (
  `transactions_id` int NOT NULL AUTO_INCREMENT,
  `transactions_token` varchar(255) NOT NULL,
  `account_id` int NOT NULL,
  `value` float(255, 2) NULL DEFAULT 0,
  `description` varchar(255) NULL,
  `type_transactions` varchar(50) NOT NULL DEFAULT "AWAITING_RISK_ANALYSIS",
  `created_at` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`transactions_id`),
  INDEX `idx_transactions_type`(`type_transactions`) USING BTREE,
  INDEX `idx_transactions_account_id`(`account_id`) USING BTREE,
  UNIQUE INDEX `idx_transactions_transactions_token`(`transactions_token`) USING BTREE,
  CONSTRAINT `fk_transactions_accounts_id` FOREIGN KEY (`account_id`) REFERENCES `db`.`accounts` (`account_id`) ON DELETE CASCADE ON UPDATE CASCADE
);


/*create triggers*/

CREATE TRIGGER `trigger_create_line_account_balance` AFTER INSERT ON `accounts` FOR EACH ROW INSERT INTO db.account_balance (account_id, balance) VALUES (new.account_id, 0);