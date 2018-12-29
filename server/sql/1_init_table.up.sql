USE reqhack;

CREATE TABLE `bin` (
                     `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                     `name` VARCHAR (32) NOT NULL UNIQUE,
                     `create_at` DATETIME default current_timestamp
);

CREATE TABLE `request` (
                         `bin` INT NOT NULL,
                         `id` INT NOT NULL AUTO_INCREMENT,
                         `data` JSON,
                         FOREIGN KEY (bin) REFERENCES bin(id),
                         PRIMARY KEY(id,bin)
);
