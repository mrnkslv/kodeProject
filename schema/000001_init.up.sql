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

insert into users(name, username, password_hash)
values ('testusername111', 'test111', '6a666768646c6b6f76696f6a6b66646b6fe38ad214943daad1d64c102faec29de4afe9da3d');
