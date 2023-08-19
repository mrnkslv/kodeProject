# notes-app REST API Documentation

* [Introduction](#introduction)
* [Access](#access)
* [Notes](#notes)
* [Token](#token)
* [Test Data for Manual Testing](#test-data)

## Introduction

This documentation provides a description of all available API handlers.

To use most of the API handlers you need to be logged in using /auth/sign-in, see [POST /auth/sign-in](#access) section for more information.

1. To create a user, use [POST /auth/sign-up](#access).
2. Then log in with created user's credentials via [POST /auth/sign-in].
3. You are in!

Test user with test data is automatically created on migrate-up, check out [test data section](#test-data) for more information.

## access

The access section. Used to authenticate the api user. 

### POST auth/sign-up

Returns user id.

Body
```
{
    "name": "name",
    "username": "testusername",
    "password": "password"
}
```

Success response 
```
{
  "id": 1
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | name, username or password are not valid 
| Internal Server Error | 500 | Server error. Possibly name/username already exists

### POST auth/sign-in

Used to log in the user. Returs access token (for 12 hours)

Body
```
{
    "username": "testusername",
    "password": "password"
}
```

Success response 
```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI1MjUxMzIsImlhdCI6MTY5MjQ4MTkzMiwiVXNlcklkIjo0fQ.F1ZC8icUUuTGwIKW47sFkWbESYBUu4FRvXutF4ILTgY"
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | invalid characters in input
| Internal Server Error | 500 | Server error. Possibly password/username incorrect.


## notes
🔑 Requires user to be logged in. (please use your access token).
Used by user to create note or get all of them. 

### POST api/notes

Create a new note. Mistakes in text and description are corrected by using Yandex Speller API automatically. 

Body
```
{
  "text": "паход в университет",
  "description": "зделать домашку"
}
```

Success response
```
{
  "new note": {
    "id": 15,
    "UserId": 3,
    "text": "поход в университет",
    "description": "сделать домашку"
  }
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | invalid characters in input
| Unauthorized      | 401  | user isn't logged in or the access token has expired
| Internal Server Error | 500 | Server error

### GET api/notes

Get all the user's notes


Success response
```
{
  "data": [
    {
      "id": 1,
      "UserId": 0,
      "text": "поход в магазин",
      "description": "купить яблоки, груши, грибы"
    },
    {
      "id": 2,
      "UserId": 0,
      "text": "поход в университет",
      "description": "сделать домашку"
    },
  ]
}

```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Unauthorized      | 401  | token has expired or user not authorized
| Internal Server Error | 500 | Server error


## Token 

When testing api by postman or other tool please add your token without quotes in Auth->Bearer Token. 


## Test Data

On migrate-up, the user and some test user data is created.

To start testing with test user, log in via [POST auth/sign-in](#access).
Body
```
{
    "username": "test111",
    "password": "password1"
}
```