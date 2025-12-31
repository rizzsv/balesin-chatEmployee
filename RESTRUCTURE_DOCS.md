# Project Structure - Balesin Chat Employee

## Struktur Folder yang Telah Direorganisasi

```
balesin-chatEmploye/
├── cmd/
│   ├── server/
│   │   └── main.go              # Application entry point
│   └── seed/
│       └── seed.go              # Database seeder
│
├── internal/
│   ├── config/                   # ✅ Configuration layer
│   │   ├── database.go          # PostgreSQL connection
│   │   ├── jwt.go               # JWT configuration
│   │   ├── websocket.go         # WebSocket upgrader config
│   │   └── redis.go             # Redis config (placeholder)
│   │
│   ├── domain/                   # ✅ Domain/Business entities
│   │   ├── user/
│   │   │   ├── user.entity.go   # User entity struct
│   │   │   ├── user.repository.go # User repository interface
│   │   │   └── user.service.go  # User business logic
│   │   │
│   │   └── chat/
│   │       ├── chat.entity.go   # Chat entity
│   │       ├── message.entity.go # Message entity
│   │       ├── chat.repository.go # Chat repository interface
│   │       └── chat.service.go  # Chat business logic
│   │
│   ├── repository/               # ✅ Data access implementations
│   │   └── postgres/
│   │       ├── user_repository_pg.go  # User DB operations
│   │       └── chat_repository_pg.go  # Chat DB operations
│   │
│   ├── transport/                # ✅ Transport layer (HTTP, WebSocket)
│   │   ├── http/
│   │   │   └── auth_handler.go  # HTTP auth endpoints
│   │   │
│   │   └── websocket/
│   │       ├── chat_handler.go  # WebSocket chat handler
│   │       ├── registry.go      # WebSocket connection registry (Hub)
│   │       └── middleware_jwt_ws.go # WebSocket JWT middleware
│   │
│   ├── middleware/               # HTTP middleware
│   │   ├── jwt.go               # JWT HTTP authentication
│   │   └── role.go              # Role-based access control
│   │
│   ├── security/                 # Security utilities
│   │   ├── jwt.go               # JWT token generation/parsing
│   │   └── password.go          # Password hashing
│   │
│   └── shared/                   # ✅ Shared utilities
│       ├── response.go          # Standard API responses
│       ├── errors.go            # Common error definitions
│       └── constants.go         # Application constants
│
├── pkg/                          # Public packages
│   └── logger/
│       ├── logger.go            # Logger initialization
│       └── http_middleware.go  # HTTP request logger
│
├── .env                          # Environment variables (gitignored)
├── .gitignore
├── go.mod
└── README.md
```

## Prinsip Arsitektur

### 1. **Domain-Driven Design (DDD)**
- Domain entities berada di `internal/domain/`
- Setiap domain memiliki entity, repository interface, dan service
- Business logic terisolasi di domain layer

### 2. **Dependency Inversion**
- Repository interfaces didefinisikan di domain layer
- Implementasi konkret (PostgreSQL) di `internal/repository/postgres/`
- Domain tidak bergantung pada infrastruktur

### 3. **Separation of Concerns**
- **Config**: Semua konfigurasi terpusat
- **Domain**: Business logic murni
- **Repository**: Data access layer
- **Transport**: HTTP/WebSocket handlers
- **Shared**: Utilities yang digunakan bersama

### 4. **Clean Code Principles**
- Single Responsibility Principle
- Interface Segregation
- Dependency Injection

## Import Paths Baru

```go
// Old
import "balesin-chatEmployee/internal/database"
import "balesin-chatEmployee/internal/domain"
import "balesin-chatEmployee/internal/repository"
import "balesin-chatEmployee/internal/handler/http"
import "balesin-chatEmployee/internal/handler/websocket"
import "balesin-chatEmployee/internal/service"

// New
import "balesin-chatEmployee/internal/config"
import "balesin-chatEmployee/internal/domain/user"
import "balesin-chatEmployee/internal/domain/chat"
import "balesin-chatEmployee/internal/repository/postgres"
import "balesin-chatEmployee/internal/transport/http"
import "balesin-chatEmployee/internal/transport/websocket"
import "balesin-chatEmployee/internal/shared"
```

## File-File yang Sudah Diperbaharui

### ✅ Sudah Dipindahkan ke Struktur Baru:
- `internal/config/*` - Configuration files
- `internal/domain/user/*` - User domain
- `internal/domain/chat/*` - Chat domain
- `internal/repository/postgres/*` - PostgreSQL implementations
- `internal/transport/http/*` - HTTP handlers
- `internal/transport/websocket/*` - WebSocket handlers
- `internal/shared/*` - Shared utilities
- `internal/middleware/jwt.go` - Updated imports
- `cmd/server/main.go` - Fully refactored with new structure

### ⚠️ File Lama (Bisa Dihapus):
- `internal/database/postgre.go` → Diganti `internal/config/database.go`
- `internal/domain/user.go` → Diganti `internal/domain/user/user.entity.go`
- `internal/repository/user_repository.go` → Diganti `internal/repository/postgres/user_repository_pg.go`
- `internal/handler/http/auth_handler.go` → Diganti `internal/transport/http/auth_handler.go`
- `internal/handler/websocket/*` → Diganti `internal/transport/websocket/*`
- `internal/service/auth_service.go` → Diganti `internal/domain/user/user.service.go`

## Cara Menjalankan

```bash
# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go

# Build
go build -o bin/server.exe cmd/server/main.go

# Run binary
./bin/server.exe
```

## Testing Build

✅ Build sukses tanpa error!
✅ Semua import paths sudah diupdate
✅ Struktur folder profesional dan scalable

## Next Steps (Opsional)

1. Hapus file-file lama yang sudah tidak digunakan
2. Tambahkan unit tests di setiap domain
3. Implementasi repository Redis untuk caching
4. Tambahkan usecase layer untuk orchestration logic yang kompleks
5. Tambahkan migration files untuk database schema
6. Setup CI/CD pipeline
