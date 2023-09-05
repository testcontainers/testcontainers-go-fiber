CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    fullname varchar,
    email varchar
);

INSERT INTO users(fullname, email) values ('Manu', 'manu@mail.com');
INSERT INTO users(fullname, email) values ('Siva', 'siva@mail.com');
