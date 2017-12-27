-- +goose Up
CREATE TABLE profiles (
	id SERIAL PRIMARY KEY,
	first_name varchar(50) NOT NULL,
	last_name varchar(50) NOT NULL,
	avatar text
);

INSERT INTO profiles (first_name, last_name) VALUES('Yahor', 'Zhylinski');
INSERT INTO profiles (first_name, last_name) VALUES('Darya', 'Zhylinskaya');

-- +goose Down
DROP TABLE profiles;
