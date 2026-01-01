# Balesin Chat Employee

A real-time chat application built with Go, featuring WebSocket communication, JWT authentication, and PostgreSQL database. This project follows clean architecture principles with domain-driven design.

## 🚀 Features

- ✅ **Real-time Chat** - WebSocket-based instant messaging
- 🔐 **JWT Authentication** - Secure token-based authentication
- 👥 **User Management** - User registration and authentication
- 💬 **Direct Messaging** - One-to-one chat between users
- 🏗️ **Clean Architecture** - Domain-driven design with clear separation of concerns
- 📦 **PostgreSQL Database** - Reliable data persistence
- 📝 **Structured Logging** - Comprehensive logging with zerolog

## 🛠️ Tech Stack

- **Language**: Go 1.21+
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **WebSocket**: Gorilla WebSocket
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Logging**: zerolog
- **Database Driver**: pgx/v5

## 📁 Project Structure

```
balesin-chatEmploye/
├── cmd/
│   ├── server/          # Application entry point
│   └── seed/            # Database seeder
│
├── internal/
│   ├── config/          # Configuration (DB, JWT, WebSocket)
│   ├── domain/          # Business entities and logic
│   │   ├── user/        # User domain
│   │   └── chat/        # Chat domain
│   ├── repository/      # Data access layer
│   │   └── postgres/    # PostgreSQL implementations
│   ├── transport/       # HTTP & WebSocket handlers
│   │   ├── http/        # HTTP REST API
│   │   └── websocket/   # WebSocket handlers
│   ├── middleware/      # HTTP middleware
│   ├── security/        # JWT & password utilities
│   └── shared/          # Shared utilities & constants
│
└── pkg/
    └── logger/          # Logging utilities
```

## 🔧 Installation & Setup

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12+
- Git

### 1. Clone the repository

```bash
git clone https://github.com/fizzsv/balesin-chatEmployee.git
cd balesin-chatEmployee
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Setup PostgreSQL Database

Create a PostgreSQL database:

```sql
CREATE DATABASE balesin_chat_employee;
```

Create the required tables:

```sql
-- Users table
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Chats table
CREATE TABLE chats (
    id VARCHAR(36) PRIMARY KEY,
    participant1 VARCHAR(36) NOT NULL,
    participant2 VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (participant1) REFERENCES users(id),
    FOREIGN KEY (participant2) REFERENCES users(id)
);

-- Messages table
CREATE TABLE messages (
    id VARCHAR(36) PRIMARY KEY,
    chat_id VARCHAR(36) NOT NULL,
    from_user VARCHAR(36) NOT NULL,
    to_user VARCHAR(36) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_read BOOLEAN DEFAULT false,
    FOREIGN KEY (chat_id) REFERENCES chats(id),
    FOREIGN KEY (from_user) REFERENCES users(id),
    FOREIGN KEY (to_user) REFERENCES users(id)
);

CREATE INDEX idx_messages_chat_id ON messages(chat_id);
CREATE INDEX idx_messages_created_at ON messages(created_at);
```

### 4. Configure Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DATABASE_URL=postgres://username:password@localhost:5432/balesin_chat_employee?sslmode=disable

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Server Configuration
PORT=8080
```

### 5. Seed Database (Optional)

```bash
go run cmd/seed/seed.go
```

### 6. Run the Application

```bash
go run cmd/server/main.go
```

Or build and run:

```bash
go build -o bin/server.exe cmd/server/main.go
./bin/server.exe
```

The server will start on `http://localhost:8080`

## 📚 API Documentation

### Authentication

#### Login
```http
POST /auth/login
Content-Type: application/json

{
  "email": "admin@company.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Protected Routes

All `/api/*` routes require JWT authentication via Bearer token:

```http
Authorization: Bearer <your-jwt-token>
```

#### Get Current User
```http
GET /api/me
Authorization: Bearer <token>
```

**Response:**
```json
{
  "user_id": "04bb24ce-bcf7-479c-9577-9e0b0eecd8cf"
}
```

## 🔌 WebSocket Documentation

### Connect to WebSocket

```javascript
const token = "your-jwt-token";
const ws = new WebSocket(`ws://localhost:8080/ws/chat?token=${token}`);

ws.onopen = () => {
  console.log("Connected to chat");
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log("Received:", data);
  // { "from": "user-id", "message": "Hello!" }
};
```

### Send Message

```javascript
ws.send(JSON.stringify({
  "to": "target-user-id",
  "message": "Hello, how are you?"
}));
```

### Message Format

**Outgoing (Client → Server):**
```json
{
  "to": "user-id-of-recipient",
  "message": "Your message here"
}
```

**Incoming (Server → Client):**
```json
{
  "from": "user-id-of-sender",
  "message": "Received message"
}
```

## 🏗️ Architecture Overview

This project follows **Clean Architecture** principles with **Domain-Driven Design**:

### Layers

1. **Domain Layer** (`internal/domain/`)
   - Contains business entities and logic
   - Defines repository interfaces
   - Independent of external frameworks

2. **Repository Layer** (`internal/repository/`)
   - Implements data access
   - PostgreSQL implementation
   - Follows repository pattern

3. **Transport Layer** (`internal/transport/`)
   - HTTP REST API handlers
   - WebSocket handlers
   - Request/Response handling

4. **Config Layer** (`internal/config/`)
   - Database connections
   - JWT configuration
   - WebSocket setup

5. **Shared Layer** (`internal/shared/`)
   - Common utilities
   - Error definitions
   - Constants

### Design Principles

- ✅ **Dependency Inversion** - High-level modules don't depend on low-level modules
- ✅ **Single Responsibility** - Each module has one reason to change
- ✅ **Interface Segregation** - Small, focused interfaces
- ✅ **Separation of Concerns** - Clear boundaries between layers

## 🔒 Security

- Passwords are hashed using bcrypt
- JWT tokens for stateless authentication
- Token expiration set to 24 hours
- WebSocket connections validated with JWT
- SQL injection prevention with parameterized queries

## 📝 Development

### Run Tests

```bash
go test ./...
```

### Build

```bash
go build -o bin/server.exe cmd/server/main.go
```

### Format Code

```bash
go fmt ./...
```

### Lint

```bash
golangci-lint run
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👤 Author

**Rizqs**
- GitHub: [@fizzsv](https://github.com/fizzsv)

## 🙏 Acknowledgments

- Gin Web Framework
- Gorilla WebSocket
- PostgreSQL
- Go Community

---

Made with ❤️ using Go
