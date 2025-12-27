# Go Product Memory Manager

The Go Product Memory Manager is a backend system for tracking products in the system. It provides user and product management features with authentication, database interactions, and RESTful API endpoints. The system operates with PostgreSQL and MongoDB while being extensible for additional functionalities.

---

## Features

- **User Management:**
  - Create and manage users with secure password hashing.
  - Login functionality using JWT-based authentication.
- **Product Management:**
  - Track products' details and maintain product-related operations.
  - Store and retrieve product price histories.
- **Authentication Middleware:**
  - Includes middleware for handling secure authentication via API keys and tokens.
- **Database Support:**
  - PostgreSQL and MongoDB support for structured and document-based data storage.
- **Extensible APIs:**
  - RESTful endpoints for users, products, and sessions.

---

## Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/bhemi28/go-product-memory-manager.git
   cd go-product-memory-manager
   ```

2. **Setup Environment Variables:**
   Create a `.env` file in the root directory with variables such as:
   ```
   POSTGRES_URL=your_postgresql_connection_string
   MONGODB_URL=your_mongodb_connection_string
   JWT_SECRET=your_secret_key
   ```

3. **Install Dependencies:**
   Ensure Go is installed, then run:
   ```bash
   go mod tidy
   ```

4. **Run the Application:**
   ```bash
   go run cmd/main.go
   ```

5. **Using the Makefile (Optional):**
   Build and run the project via:
   ```bash
   make build
   make run
   ```

---

## Project Structure

The project follows a modular structure, making it easy to navigate and extend.

- **`cmd/`:**
  - Contains the entry point of the application and API initialization.
- **`db/`:**
  - Handles database-specific connectivity for both PostgreSQL and MongoDB.
  - Includes various data models.
- **`internal/`:**
  - Contains handlers for internal logic and database queries.
- **`service/`:**
  - Holds services for authentication, users, and product management.
  - `auth`, `product`, and `user` submodules manage domain-specific responsibilities.
- **`types/`:**
  - Defines the structures for database and API entities such as `User` and `Product`.
- **`utils/`:**
  - Contains utility functions such as JSON parsing for HTTP interactions.

---

## API Endpoints

### User Routes (JWT Protected)
- **Create User**:
  - `POST /user`
  - Payload:
    ```json
    {
      "username": "example-user",
      "email": "user@example.com",
      "password": "password123"
    }
    ```

- **User Login (Obtain JWT)**:
  - `POST /user/login`
  - Payload:
    ```json
    {
      "username": "example-user",
      "password": "password123"
    }
    ```

### Product Routes
- **Create Product**:
  - `POST /product`
  - Payload:
    ```json
    {
      "name": "Product Name",
      "link": "https://product.com",
      "category": "Electronics"
    }
    ```

- **List Products**:
  - `GET /product`

---

## Technologies Used

- **Language:** Go
- **Frameworks/Libraries:**
  - `go-chi`: HTTP routing.
  - `sqlc`: PostgreSQL database queries generation.
  - `mongodb`: MongoDB database operations.
  - `golang-jwt`: Authentication via JWT.
  - `bcrypt`: Password hashing and verification.
- **Databases:**
  - PostgreSQL: For structured data.
  - MongoDB: For document-oriented data.

---

## Contributing

1. **Fork the Repository:**
   [Fork](https://github.com/bhemi28/go-product-memory-manager/fork)
2. Clone your fork:
   ```bash
   git clone https://github.com/<your-username>/go-product-memory-manager.git
   ```
3. Make your changes.
4. Submit a pull request.

---

## License

This repository currently lacks a license. Contact the repository owner for usage inquiries.

---

## Author

- **[bhemi28](https://github.com/bhemi28)**  

Feel free to reach out via GitHub for questions, feedback, or suggestions.
