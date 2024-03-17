# Junior Test task. Rest API CRUD for books

This project provides a REST API for managing a collection of books, allowing basic CRUD (Create, Read, Update, Delete) operations. It's built using [your technology stack, e.g., Node.js, Express, MongoDB] and is containerized using Docker for easy setup and deployment.

## Features

- CRUD operations for books
- Dockerized environment for easy setup
- Swagger documentation for API endpoints

## Prerequisites

Before you begin, ensure you have the following installed:
- Docker and Docker Compose
- Make (for running Makefile commands)

## Installation

1. Create `.env` file based on `.env.example`. Change values if needed
2. Run docker container

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

API Base URL: [localhost:8000](http://localhost:8000)\
Swagger Documentation: For a detailed look at the available API endpoints and their specifications, visit the Swagger documentation at: [localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)