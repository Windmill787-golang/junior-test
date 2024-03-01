# Junior Test task. Rest API CRUD for books

## Installation

1. Create `.env` file based on `.env.example`. Change values if needed
2. Run docker container with postgres

```bash
make docker-up
```

3. Apply migrations

```bash
make migrate-up
```

4. Run application

```bash
make run
```

## Usage

API url: [localhost:8000](http://localhost:8000) (port can be different, depending on `.env`)\
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

```json
{
  "username": "test",
  "password": "123"
}
```

After registration, you can sign in to receive jwt token\
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

```json
{
  "title": "Cthulhu",
  "author": "Lovecraft",
  "description": "Very scary book",
  "genre": "horror",
  "page_count": 546,
  "year": 1910,
  "price": 5000
}
```
