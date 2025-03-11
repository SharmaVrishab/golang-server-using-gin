📘 Course Management API (Go + Gin)
A lightweight and efficient RESTful API for managing courses, built using Golang and the Gin framework. This API provides endpoints for creating, retrieving, updating, and deleting courses, making it ideal for learning and production use.

🚀 Features
Retrieve all courses (GET /courses)
Fetch a specific course by ID (GET /course/:id)
Create a new course (POST /createcourse)
Update an existing course (POST /updatecourse)
Delete a course (POST /deletecourse)
Home route (GET /)
🛠 Tech Stack
Golang – Backend language
Gin Framework – HTTP router & middleware
JSON Handling – Struct-based serialization
Random ID Generation – Unique course identifiers
📦 Installation & Setup
Clone the repository:
sh
Copy
Edit
git clone https://github.com/your-username/course-management-api.git  
cd course-management-api  
Install dependencies:
sh
Copy
Edit
go mod tidy  
Run the application:
sh
Copy
Edit
go run main.go  
The API will be available at:
arduino
Copy
Edit
http://localhost:4000  
📌 Example Request: Create a Course
Endpoint: POST /createcourse
Request Body (JSON):

json
Copy
Edit
{
  "coursename": "Kubernetes Essentials",
  "price": 2499,
  "author": {
    "fullname": "Michael Scott",
    "website": "https://michaelscott.com"
  }
}
