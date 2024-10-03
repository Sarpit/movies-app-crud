
# Movie Management API Documentation

## Overview

This application is a Go-based RESTful API server for managing a movie database. It uses MySQL as the database and the Gorilla Mux router for handling HTTP requests.

## Endpoints

### 1. Home

- **URL:** `/`
- **Method:** `GET`
- **Description:** Returns a welcome message.
- **Response:** Plain text welcome message.

### 2. Get All Movies

- **URL:** `/movies`
- **Method:** `GET`
- **Description:** Retrieves all movies from the database.
- **Response:** JSON array of movie objects.

### 3. Get Single Movie

- **URL:** `/movies/{id}`
- **Method:** `GET`
- **Description:** Retrieves a specific movie by its ID.
- **URL Parameters:** `id=[integer]` where `id` is the ID of the movie.
- **Success Response:** JSON object of the movie.
- **Error Response:**
    - Code: 404 NOT FOUND if the movie doesn't exist.
    - Code: 400 BAD REQUEST if the ID is invalid.

### 4. Create Movie

- **URL:** `/movies`
- **Method:** `POST`
- **Description:** Creates a new movie.
- **Request Body:** JSON object of the movie to be created.
- **Success Response:** JSON object of the created movie.
- **Error Response:**
    - Code: 400 BAD REQUEST if the request payload is invalid.

### 5. Update Movie

- **URL:** `/movies/{id}`
- **Method:** `PUT`
- **Description:** Updates an existing movie.
- **URL Parameters:** `id=[integer]` where `id` is the ID of the movie to update.
- **Request Body:** JSON object of the updated movie details.
- **Success Response:** JSON object of the updated movie.
- **Error Response:**
    - Code: 400 BAD REQUEST if the ID is invalid or the payload is invalid.

### 6. Delete Movie

- **URL:** `/movies/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a specific movie.
- **URL Parameters:** `id=[integer]` where `id` is the ID of the movie to delete.
- **Success Response:** JSON object with a success message.
- **Error Response:**
    - Code: 400 BAD REQUEST if the ID is invalid.

## Error Handling

The API uses standard HTTP response codes to indicate the success or failure of requests. In case of errors, a JSON response is returned with an "error"

## Running the Application

To run the application:

1. Ensure you have Go installed on your system.
2. Set up your MySQL database and update the connection details in const.go.
   ```sql
   package main
   const DBUser = "[DATABASE USER]"
   const DBName = "[DATABASE NAME]"
   const DBPassword = "[DATABASE PASSWORD]"
   const DBHost = "[DATABASE HOST]"
   ```

3. Install the required dependencies:
   ```
   go get github.com/go-sql-driver/mysql
   go get github.com/gorilla/mux
   ```
4. Run the application:
   ```
   go build
   ./movies-app-crud
   ```

The server will start on `localhost:3000`.

## Notes

- The application uses the `database/sql` package with the MySQL driver for database operations.
- Gorilla Mux is used for routing HTTP requests.
- JSON is used for data exchange in request and response bodies.
