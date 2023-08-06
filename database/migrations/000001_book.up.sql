CREATE TABLE IF NOT EXISTS books (
    id serial not null unique,
    title varchar(255) not null,
    description text not null,
    genre varchar(255) not null,
    author varchar(255) not null,
    page_count int not null,
    year int not null,
    price int not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);