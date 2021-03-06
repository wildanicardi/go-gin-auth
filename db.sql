CREATE DATABASE go_auth

CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
		updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp(),
		UNIQUE (email)
	);