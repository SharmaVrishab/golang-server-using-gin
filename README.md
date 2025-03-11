# Course API with Go and Gin

## Overview
This is a simple RESTful API for managing courses using the Go programming language and the Gin framework.

## Features
- Get all courses
- Get a single course by ID
- Create a new course
- Update an existing course
- Delete a course

## Tech Stack
- **Go**
- **Gin Web Framework**

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/your-repo.git
   ```
2. Change into the project directory:
   ```sh
   cd your-repo
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

| Method | Endpoint            | Description               |
|--------|---------------------|---------------------------|
| GET    | `/`                 | Home Page                 |
| GET    | `/courses`          | Get all courses           |
| GET    | `/course/:id`       | Get a single course       |
| POST   | `/createcourse`     | Create a new course       |
| POST   | `/updatecourse/:id` | Update an existing course |
| POST   | `/deletecourse/:id` | Delete a course           |

## Example JSON Request for Creating a Course
```json
{
  "coursename": "New Course",
  "price": 1200,
  "author": {
    "fullname": "John Doe",
    "website": "https://johndoe.com"
  }
}
```



