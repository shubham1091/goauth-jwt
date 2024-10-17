### Project Overview

This project is a RESTful API built with Go, designed for user authentication and authorization using JSON Web Tokens (JWT). It includes features like user registration, login, and role-based access control.

### Local Setup Instructions

#### Prerequisites

1. **Go**: Ensure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

2. **Docker**: Install Docker to run MongoDB locally. You can download it from the [Docker website](https://www.docker.com/products/docker-desktop).

3. **Postman**: Use Postman to test API endpoints. Download it from the [Postman website](https://www.postman.com/downloads/).

#### Steps to Set Up Locally

1. **Clone the Repository**

   First, fork and clone the project repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/goauth-jwt.git
   cd goauth-jwt
   ```

2. **Install Dependencies**

   Use Go modules to install the necessary dependencies:

   ```bash
   go mod tidy
   ```

3. **Run MongoDB Locally Using Docker**

   You can use Docker to run MongoDB locally. Pull the MongoDB image and run it:

   ```bash
      docker run -d \
     --name mongodb-container \
     -e MONGO_INITDB_DATABASE=yourDatabaseName \
     -p 27017:27017 \
     mongo
   ```

   This will start MongoDB on port 27017.

4. **Set Up Environment Variables**

   Create a `.env` file in the root directory and add your configuration details:

   ```
   MONGODB_URL=mongodb://localhost:27017/yourDatabaseName
   SECRET_KEY=your_secret_key
   ```

5. **Run the Application Locally**

   You can run the application using Go:

   ```bash
   go run main.go
   ```

   This will start your Go application, which will connect to the MongoDB instance running in Docker.

6. **Test with Postman**

   Open Postman and import the `postman_collection.json` file to test the API endpoints.

By following these steps, you can set up and run the Go project locally, use Docker to manage MongoDB, and test the application using Postman. Docker remains optional for the application itself, but it provides a convenient way to run MongoDB without installing it directly on your machine.
