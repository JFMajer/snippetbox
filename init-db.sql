CREATE DATABASE IF NOT EXISTS snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE snippetbox;

CREATE TABLE IF NOT EXISTS snippets (
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

-- Insert sample data
INSERT INTO snippets (title, content, created, expires) VALUES
('First snippet', 'This is the content for the first snippet.', NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY)),
('Second snippet', 'Content for the second snippet.', NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY)),
('Third snippet', 'Here is a bit more content for the third snippet.', NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY)),
('Fourth snippet', 'Fourth snippet content goes here.', NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY)),
('Fifth snippet', 'And this is the fifth snippet.', NOW(), DATE_ADD(NOW(), INTERVAL 7 DAY));

CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';