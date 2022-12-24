# post-api
A simple Post API built with Gin and Gorm in the Go programming language and Postgres as a database. This is to further understand the basics of buiding of building REST APIs and this is also my subimission to the Slightly Techie Challenge

# Description

This is a REST API that allows you to create, read, update and delete posts. It is built with Gin and Gorm in the Go programming language and Postgres as a database.

# Live Environment

- The Deployed Version of the API can be found at `https://post-api-zc3z.onrender.com/`

- The API documentation for the deployment can be found at `https://documenter.getpostman.com/view/8806007/2s8Z6vYuAP`

# Usage

- Clone the repository

```bash
git clone https://github.com/buabaj/post-api.git
```

- Navigate to the project directory and add your database credentials to the `.env` file

```bash
cd post-api
touch .env
DB_URL=<your_postgresdb_connection_string>
```


- Run the application

```bash
go run main.go
```

- Test the endpoints using Postman or any other API testing tool

# Endpoints

| Name | Request Method | Endpoint | Parameters | Response code |
| --- | --- | --- | --- | --- |
| List all Posts | GET | localhost:8080/posts | None | 200 OK |
| Get single post | GET | localhost:8080/posts/{id} | id: int | 200 OK |
| Add Post | POST | localhost:8080/posts | None | 201 Created |
| Delete Post | DELETE | localhost:8080/posts/{id} | id: int | 200 OK |
| Update Post | PUT | localhost:8080/posts/{id} | id: int | 200 OK |

## Sample Data

Sample JSON data to test `Add Post` endpoint

```json
{
    "Title":"JavaScript- a dumb programming language, Use Go instead",
    "Body": "Lorem ipsum blah blah blah",
    "Author":"Mawuli"
}
```

Sample JSON response for `List All Posts` endpoint

```json
{
    "message": "posts fetched successfully",
    "posts": [
        {
            "ID": 1,
            "CreatedAt": "2022-12-24T16:46:55.910741Z",
            "UpdatedAt": "2022-12-24T16:46:55.910741Z",
            "DeletedAt": null,
            "title": "Python, a modern programming language",
            "body": "Lorem ipsum blah blah blah",
            "author": "Jerry Buaba"
        },
        {
            "ID": 2,
            "CreatedAt": "2022-12-24T16:47:23.795743Z",
            "UpdatedAt": "2022-12-24T16:47:23.795743Z",
            "DeletedAt": null,
            "title": "JavaScript, a dumb programming language",
            "body": "Lorem ipsum blah blah blah",
            "author": "Buaba"
        }
    ]
}
```

# Future Work

Learn how to perform database mocking in Go.
