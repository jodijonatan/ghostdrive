# GhostDrive

A lightweight, secure file storage application inspired by cloud drive services. GhostDrive provides a simple yet powerful platform for users to store, manage, and access their files through a web interface, with a robust backend API ensuring data integrity and user authentication.

## Features

- **User Authentication**: Secure login system using JWT tokens for session management.
- **File Management**: Full CRUD operations for files, including upload, download, update, and deletion.
- **RESTful API**: Clean and intuitive API endpoints for seamless integration.
- **Cross-Origin Support**: Configured CORS for flexible client-side interactions.
- **Database Integration**: SQLite database with GORM for efficient data persistence.
- **Modern Frontend**: Responsive React application built with Vite and styled with Tailwind CSS.
- **Real-time Updates**: Client-side state management for dynamic file operations.

## Tech Stack

### Backend

- **Go**: Programming language for high-performance server-side logic.
- **Gin**: Web framework for building RESTful APIs.
- **GORM**: ORM library for database interactions.
- **SQLite**: Lightweight database for data storage.
- **JWT**: Token-based authentication for secure user sessions.
- **CORS**: Middleware for handling cross-origin requests.

### Frontend

- **React**: JavaScript library for building user interfaces.
- **Vite**: Fast build tool and development server.
- **Tailwind CSS**: Utility-first CSS framework for styling.
- **Axios**: HTTP client for API communication.
- **IndexedDB**: Client-side storage for offline capabilities.

## Installation

### Prerequisites

- Go 1.24.4 or later
- Node.js 18 or later
- npm or yarn

### Backend Setup

1. Navigate to the backend directory:

   ```bash
   cd backend
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

3. Run the backend server:
   ```bash
   go run cmd/main.go
   ```
   The server will start on `http://localhost:8080`.

### Frontend Setup

1. Navigate to the frontend directory:

   ```bash
   cd frontend
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```
   The application will be available at `http://localhost:5173`.

## Usage

1. Start both the backend and frontend servers as described in the installation steps.
2. Open your browser and navigate to the frontend URL.
3. Register or log in to access your personal file storage.
4. Upload files, view your file list, and manage your documents through the intuitive web interface.

## API Endpoints

### Authentication

- `POST /login`: Authenticate user and receive JWT token.

### Protected Routes (Require JWT Token)

- `GET /me`: Retrieve current user information.
- `GET /files`: List all files for the authenticated user.
- `POST /files`: Upload a new file.
- `PUT /files/:id`: Update an existing file.
- `DELETE /files/:id`: Delete a file.

### Request/Response Examples

#### Login

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "user", "password": "pass"}'
```

#### Get Files

```bash
curl -X GET http://localhost:8080/files \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Project Structure

```
ghostdrive/
├── backend/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── auth/
│   │   ├── database/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── models/
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── assets/
│   │   ├── db/
│   │   ├── pages/
│   │   └── services/
│   ├── package.json
│   ├── vite.config.js
│   └── index.html
└── README.md
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

<!--
## License

This project is licensed under the MIT License - see the LICENSE file for details. -->
