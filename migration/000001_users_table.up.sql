 CREATE TABLE users (
     id SERIAL PRIMARY KEY, 
     email VARCHAR(255) NOT NULL,
     full_name VARCHAR(255) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, version INT DEFAULT 1
      );