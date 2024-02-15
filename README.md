# Go Web Service Template

## Overview
This Go project is structured as a RESTful API service designed to handle user management and authentication, including operations such as user creation, deletion, updating, and retrieval. It employs JWT (JSON Web Tokens) for secure authentication and includes middleware for token validation. The service interacts with a PostgreSQL database, and Docker is used for environment setup and deployment.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go 1.22.0
- Docker and Docker Compose
- PostgreSQL (setup via Docker Compose)
- PgAdmin4 (setup via Docker Compose)

### Installing

1. **Clone the repository**
   ```bash
   git clone git@github.com:HarrisonHemstreet/go-web-service-template.git
   ```

2. **Navigate to the project directory**
   ```bash
   cd go-web-service-template
   ```

3. **Start the PostgreSQL database with Docker Compose**
   ```bash
   cd deployments
   docker-compose up -d
   ```

4. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

5. **Run the application**
   ```bash
   go run ./cmd/main.go
   ```

## Project Structure
- `cmd/main.go`: Entry point of the application, initializing the server and dependencies.
- `deployments`: Contains Docker Compose and PostgreSQL initialization scripts for setting up the environment.
- `internal/database/db.go`: Handles database connection and operations.
- `internal/handler`: Contains error handling, login logic, and user-related operations (create, delete, update, get).
- `internal/middleware/jwt/validate_token.go`: Middleware for JWT validation.
- `internal/model`: Defines data models for error responses, JWT claims, products, and users.
- `internal/service/user`: Services for user authentication, fetching, and management.
- `internal/utils`: Utilities for password hashing, JWT creation, and route registration.

## Features
- **User Authentication**: Secure login mechanism using JWT for handling user sessions.
- **User Management**: APIs for creating, updating, deleting, and fetching user details.
- **Product Management**: Basic structure for managing product information (extendable).
- **Security**: Password hashing and token validation for secure access.
- **Dockerization**: Easy deployment with Docker and PostgreSQL setup.

## API Endpoints
Document specific endpoints provided by your application, such as:
- `/user` `POST`: Create a new user.
- `/user` `DELETE`: Delete an existing user.
- `/user` `PUT`: Update details of an existing user.
- `/user` `GET`: Retrieve details of a specific user.
- `/login` `POST`: Authenticate a user and return a JWT.

## Security
This project uses JWT for authentication and bcrypt for password hashing, ensuring secure data handling and user authentication.

## Contributing
Please read `CONTRIBUTING.md` for details on our code of conduct, and the process for submitting pull requests to the project.

## Authors
- **Harrison Hemstreet** - *Initial work*

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments
- Inspired by all the time spent bootstrapping web services from scratch.
