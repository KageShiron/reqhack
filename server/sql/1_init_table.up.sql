CREATE DATABASE IF NOT EXISTS `reqhack`;
USE reqhack;

CREATE TABLE `bin` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR (32) NOT NULL UNIQUE,
);

CREATE TABLE `request` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `bin` INT NOT NULL,
  `data` JSON,
  FOREIGN KEY (bin) REFERENCES bin(id)
);
