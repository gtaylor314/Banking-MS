CREATE TABLE IF NOT EXISTS `customers` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `date_of_birth` date NOT NULL,
  `city` varchar(255) NOT NULL,
  `zipcode` varchar(255) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '1'
) AUTO_INCREMENT=6;

INSERT INTO `customers` VALUES
  (1, 'Jake', '1985-01-10', 'Brooklyn', '11219', 1),
  (2, 'Amy', '1986-06-14', 'Brooklyn', '11218', 1),
  (3, 'Gina', '1985-07-21', 'Brooklyn', '11219', 0),
  (4, 'Charles', '1980-02-04', 'Staten Island', '10306', 1),
  (5, 'Rosa', '1984-03-05', 'Brooklyn', '11218', 1);

CREATE TABLE IF NOT EXISTS `accounts` (
  `account_id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `customer_id` int(11) NOT NULL,
  `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `account_type` varchar(255) NOT NULL,
  `amount` DECIMAL(20,2) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '1',
  INDEX `acct_cust_fk` (`customer_id`),
  CONSTRAINT `acct_cust_fk` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
) AUTO_INCREMENT=107;

INSERT INTO `accounts` VALUES
  (101, 1, '2010-09-15 10:20:00', 'Saving', '1000.00', 1),
  (102, 2, '2000-01-13 08:15:05', 'Saving', '15000.00', 1),
  (103, 3, '2008-05-20 12:30:30', 'Checking', '2084.00', 0),
  (104, 4, '2009-02-03 09:13:47', 'Saving', '4244.00', 1),
  (105, 2, '2001-05-20 08:10:04', 'Checking', '18854.00', 1),
  (106, 5, '2007-09-09 10:10:09', 'Saving', '8029.00', 1);

CREATE TABLE IF NOT EXISTS `transactions` (
  `transaction_id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` int(11) NOT NULL,
  `amount` DECIMAL(20,2) NOT NULL,
  `transaction_type` varchar(255) NOT NULL,
  `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX `trans_acct_fk` (`account_id`),
  CONSTRAINT `trans_acct_fk` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
);
