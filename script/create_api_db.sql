CREATE DATABASE api_todo;

CREATE TABLE todos (id serial primary key, title varchar, description text, done bool default FALSE);
