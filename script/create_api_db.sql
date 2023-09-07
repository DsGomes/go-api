CREATE DATABASE api_todo;

-- go to api_todo database and run table creation
CREATE TABLE todos (
    id uuid primary key, 
    title varchar, 
    description text, 
    done bool default FALSE,
    created_at date default now(),
    updated_at date
);
