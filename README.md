# Blog Post `API`

## Overview    

The Blog Post API is a web application built using Go (Golang). It allows users to manage articles and comments, supporting features such as creating, retrieving, and replying to comments. The application uses GORM for database operations and Gorilla Mux for routing.

## Features

### Articles:

- Create articles.
- Fetch article content by ID.
- List all articles.

### Comments:

- Add comments to articles.
- Reply to existing comments.
- Retrieve comments for a specific article.

## Technologies Used

- Backend: Go (Golang)
- Frameworks: Gorilla Mux, GORM
- Database: MySQL

## Environment Variables: godotenv

## Testing: testify/suite, httptest




Run the Application:

go run main.go

Access the Application:
The server runs on http://localhost:8080.

### API Endpoints

#### Articles

Create Article

Method: POST

Endpoint: /create/articles

Payload:

{
  "nickname": "JohnDoe",
  "title": "My First Article",
  "content": "This is the content of the article."
}

Response:

{
  "id": 1,
  "nickname": "JohnDoe",
  "title": "My First Article",
  "content": "This is the content of the article.",
  "created_at": "2023-12-23T10:00:00Z"
}

### Get Article Content

Method: GET

Endpoint: /articles/{id}

Response:

{
  "id": 1,
  "nickname": "JohnDoe",
  "title": "My First Article",
  "content": "This is the content of the article.",
  "created_at": "2023-12-23T10:00:00Z"
}

### Comments

#### Add Comment

Method: POST

Endpoint: /comment/article/{id}

Payload:

{
  "nickname": "JaneDoe",
  "comment": "This is a comment."
}

Response:

{
  "id": 1,
  "article_id": 1,
  "nickname": "JaneDoe",
  "comment": "This is a comment.",
  "created_at": "2023-12-23T10:30:00Z"
}

#### Add Reply

Method: POST

Endpoint: /comments/reply

Payload:

{
  "parent_comment_id": 1,
  "article_id": 1,
  "reply_comment": "This is a reply.",
  "nickname": "ReplyUser"
}

Response:

{
  "id": 2,
  "article_id": 1,
  "parent_comment_id": 1,
  "nickname": "ReplyUser",
  "comment": "This is a reply.",
  "created_at": "2023-12-23T10:35:00Z"
}

### Get Comments by Article ID

Method: GET

Endpoint: /comments/{article_id}

Response:

[
  {
    "id": 1,
    "article_id": 1,
    "nickname": "JaneDoe",
    "comment": "This is a comment.",
    "created_at": "2023-12-23T10:30:00Z"
  },
  {
    "id": 2,
    "article_id": 1,
    "parent_comment_id": 1,
    "nickname": "ReplyUser",
    "comment": "This is a reply.",
    "created_at": "2023-12-23T10:35:00Z"
  }
]

## Testing

### Unit Testing:

Tests are written using testify/suite.

Mock services simulate database interactions.

### Test Scenarios:

Create an article.

Get article content by ID.

Add a comment.

Add a reply to a comment.

Fetch comments for an article.

###  Future Enhancements

- Add user authentication.
- Implement pagination for articles and comments.
- Enhance error handling and logging.