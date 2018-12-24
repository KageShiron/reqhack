use reqhack;
DELETE FROM `bin` WHERE (created_at < CURRENT_TIMESTAMP - INTERVAL 2 DAY);
