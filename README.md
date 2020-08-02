![Go](https://github.com/waznico/go-book-api/workflows/Go/badge.svg)

# Go Book API Example
A simple service to manage a book library with title, author and rating. This is an example API written in Go and using the frameworks fiber (API) and gorm (Database ORM Framework). It uses a SQLite3 database to store all data and supports all CRUD options on the book endpoint. I also implemented error handling and returning status codes to demonstrate how it could be done.

# Usage
You can build this programm on your platform by yourself or simply start it with "`go run .`". The application opens a socket at the port 3000.

The following endpoints are currently exposed:

| Endpoint                     | Description                         |
|------------------------------|-------------------------------------|
| /api/v1/books                | Lists all books                     |
| /api/v1/book/6               | Returns book with id 6 if available |
| /api/v1/book   (POST)        | Adds a new book                     |
| /api/v1/book/6 (PUT)         | Updates book with id 6 if available |
| /api/v1/book/6 (DELETE)      | Deletes book with id 6 if available |

# Examples
Here are some examples for the requests.

GET /api/v1/books

Response:
```json
[
    {
        "ID": 1,
        "CreatedAt": "2020-08-02T16:25:23.111696+02:00",
        "UpdatedAt": "2020-08-02T16:25:23.111696+02:00",
        "DeletedAt": null,
        "title": "1984",
        "author": "George Orwell",
        "rating": 5
    },
    {
        "ID": 6,
        "CreatedAt": "2020-08-02T16:37:06.757333+02:00",
        "UpdatedAt": "2020-08-02T17:01:25.807841+02:00",
        "DeletedAt": null,
        "title": "Alterra: Die Allianz der Drei",
        "author": "Maxime Chattam",
        "rating": 3
    }
]
```


GET /api/v1/book/6

Response:
```json
{
    "ID": 6,
    "CreatedAt": "2020-08-02T16:37:06.757333+02:00",
    "UpdatedAt": "2020-08-02T17:02:10.801443+02:00",
    "DeletedAt": null,
    "title": "Alterra: Die Allianz der Drei",
    "author": "Maxime Chattam",
    "rating": 4
}
```


POST /api/v1/book

Request:

```json
{
    "title": "Alterra: Die Allianz der Drei",
    "author": "Maxime Chattam",
    "rating": 5
}
```

Response:
```json
{
    "ID": 6,
    "CreatedAt": "2020-08-02T16:37:06.757333+02:00",
    "UpdatedAt": "2020-08-02T16:37:06.757333+02:00",
    "DeletedAt": null,
    "title": "Alterra: Die Allianz der Drei",
    "author": "Maxime Chattam",
    "rating": 5
}
```


PUT /api/v1/book/6

Request:

```json
{
    "title": "Alterra: Die Allianz der Drei",
    "author": "Maxime Chattam",
    "rating": 4
}
```

Response:

```json
{
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "title": "Alterra: Die Allianz der Drei",
    "author": "Maxime Chattam",
    "rating": 4
}
```

DELETE /api/v1/book/6

Response:

```text
OK
```
