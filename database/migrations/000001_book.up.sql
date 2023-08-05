CREATE TABLE books (
    id serial not null unique,
    title varchar(255) not null,
    description text not null,
    genre varchar(255) not null,
    author varchar(255) not null,
    page_count int not null,
    release_date date not null,
    price int not null
);