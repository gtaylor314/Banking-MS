CREATE TABLE IF NOT EXISTS `users` (
  `username` VARCHAR(255) NOT NULL PRIMARY KEY,
  `password` VARCHAR(255) NOT NULL,
  `role` VARCHAR(255) NOT NULL,
  `customer_id` INT,
  `created_on` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX `user_cust_fk` (`customer_id`),
  CONSTRAINT `user_cust_fk` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
);

INSERT INTO `users` VALUES
  ('1', 'passw0rd123', 'user', 1, '2022-08-17 08:15:05'),
  ('admin', 'passw0rd123', 'admin', NULL, '2022-08-17 08:15:10');