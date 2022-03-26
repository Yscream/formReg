CREATE TABLE IF NOT EXISTS tokens(
	users_id INT REFERENCES users_data (id),
	token VARCHAR
);