CREATE TABLE IF NOT EXISTS credentials(
	users_id INT UNIQUE REFERENCES users_data (id),
	salt VARCHAR,
	hash VARCHAR
);