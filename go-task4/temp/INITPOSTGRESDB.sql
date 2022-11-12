CREATE USER pg;
CREATE DATABASE pg;
GRANT ALL PRIVILEGES ON DATABASE pg TO pg;
CREATE TABLE accounts (
	userid varchar PRIMARY KEY,
    name varchar not null,
    surname varchar not null
);
CREATE TABLE phones (
	id serial PRIMARY KEY,
	userid  varchar,
	phones varchar,
	FOREIGN KEY (userid) REFERENCES accounts (userid) 
);