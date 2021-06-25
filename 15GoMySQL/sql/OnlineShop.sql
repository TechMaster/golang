CREATE TABLE `products` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `madein` varchar(2),
  `price` int,
  `manufacturer_id` int,
  `category_id` int
);

CREATE TABLE `product_properies` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `key` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `type` ENUM ('0', '1', '2', '3', '4')
);

CREATE TABLE `countries` (
  `code` varchar(2) PRIMARY KEY NOT NULL,
  `name` varchar(255) NOT NULL
);

CREATE TABLE `product_prices` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `price` int,
  `created_at` datetime DEFAULT (now())
);

CREATE TABLE `product_medias` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `product_id` int,
  `uri` varchar(255) NOT NULL,
  `media_type` ENUM ('photo', 'vIDeo', 'PDF')
);

CREATE TABLE `manufacturers` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `country_code` varchar(2)
);

CREATE TABLE `categories` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `parent_id` int
);

CREATE TABLE `relate_products` (
  `product_id` int,
  `relate_id` int,
  `relation` ENUM ('oldversion', 'newversion', 'similar', 'recommend') NOT NULL
);

CREATE TABLE `users` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) UNIQUE NOT NULL,
  `mobile` varchar(255) UNIQUE,
  `password` varchar(255) NOT NULL
);

CREATE TABLE `customers` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int
);

CREATE TABLE `addresses` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `customer_id` int NOT NULL,
  `addr` varchar(255) NOT NULL,
  `city_id` int
);

CREATE TABLE `cities` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL
);

ALTER TABLE `products` ADD FOREIGN KEY (`madein`) REFERENCES `countries` (`code`);

ALTER TABLE `products` ADD FOREIGN KEY (`manufacturer_id`) REFERENCES `manufacturers` (`ID`);

ALTER TABLE `products` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`ID`);

ALTER TABLE `product_properies` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `product_prices` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `product_medias` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `manufacturers` ADD FOREIGN KEY (`country_code`) REFERENCES `countries` (`code`);

ALTER TABLE `relate_products` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`ID`);

ALTER TABLE `relate_products` ADD FOREIGN KEY (`relate_id`) REFERENCES `products` (`ID`);

ALTER TABLE `customers` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`ID`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`ID`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`city_id`) REFERENCES `cities` (`ID`);
