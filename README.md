# Rest API CRUD for books

Manipulate books via API

## Instalation

1. Copy `.env.example` file to create `.env` file
2. Run docker container with postgres

```
make up
```

3. Apply migrations

```
make migrate-up
```

4. Run application

```
make run
```

## Usage

API url: [localhost:8000](http://localhost:8000)\
Endpoints for books observation:

```
GET /books    - Get list of all books
GET /book/:id - Get single book by id
```

Guest user can only view list of books or single book.\
Authentication is needed to `create/update/delete` books\
Available actions for authentication:

```
POST /auth/sign-up - Registration
POST /auth/sigh-in - Login
```

Registration and authentication example body:

```
{
    "username": "test",
    "password": "123"
}
```

After registration you can sign in to receive jwt token\
Then place your token to Authentication header

```
Authentication: Bearer <your_token_here>
```

If you set valid token, you are able to `create/update/delete`books\
Endpoints for books manipulations:

```
POST /book       - Create book
PUT /book/:id    - Update book
DELETE /book/:id - Delete book
```

User can `update/delete` only books he created.\
Books example body:

```
{
    "title": "Ktulhu",
    "author": "Lovecraft",
    "description": "Very scary book",
    "genre": "horror",
    "page_count": 546,
    "year": 1910,
    "price": 5000
}
```
