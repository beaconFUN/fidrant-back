create table if not exists users(
    id serial primary key,
    name text
);

INSERT INTO users(id, name) VALUES(1, 'sumeshi1');
INSERT INTO users(id, name) VALUES(2, 'sumeshi2');