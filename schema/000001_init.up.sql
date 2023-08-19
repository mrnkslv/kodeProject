CREATE TABLE users 
(
    id serial not null unique,
    name varchar (255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null 
);

CREATE TABLE notes
(
    id serial not null unique,
    user_id integer references users (id),
    text varchar(255) not null,
    description varchar(255) not null
);

