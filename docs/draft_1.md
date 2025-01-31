# **First draft and high level architecture idea**

```
/chat-service      -> Main service: WebSocket chat, user management, message storage
    /api
    /internal
    /pkg
/grpc-service      -> Secondary service: spam filtering, message transformations
    /internal
    /proto
    /pkg
```

---

## **1) Chat Service (Main API Service)**

This service will:

- Handle **WebSocket** communication.
- Store chat messages in **PostgreSQL**.
- Provide **REST API** for fetching messages and user data.
- Implement **authentication (JWT/OAuth)**.
- Publish events to **Redis Pub/Sub** for multi-instance sync.

### **Project Structure (`/chat-service`)**

```
chat-service/
│── cmd/                # Entry point for main API
│    └── main.go        # Starts the server
│
├── api/                # HTTP handlers (REST API)
│    ├── chat.go        # Handles message send/receive
│    ├── user.go        # User registration/login
│    └── middleware.go  # Auth & request validation
│
├── internal/           # Business logic (use-case services)
│    ├── chat/          # Chat domain logic
│    │    ├── chat_service.go  # Core chat logic
│    │    ├── repository.go    # DB interactions
│    │    ├── websocket.go     # WebSocket handling
│    │    └── redis.go         # Redis pub/sub integration
│    │
│    └── user/          # User domain logic
│        ├── user_service.go
│        └── repository.go
│
├── pkg/                # Shared utilities & libraries
│    ├── logger/        # Structured logging setup
│    ├── config/        # Configuration loader (env, YAML)
│    ├── database/      # PostgreSQL connection pool
│    └── auth/          # JWT & OAuth helpers
│
├── test/               # Unit & integration tests
```

### **Core Components**

- `chat_service.go`: Business logic for handling messages.
- `repository.go`: Interacts with PostgreSQL (saving/retrieving messages).
- `websocket.go`: Manages WebSocket connections.
- `redis.go`: Handles Pub/Sub for syncing messages across instances.

---

## **2) gRPC Service (Supporting Service)**

This service will:

- Process messages for **spam filtering & transformations**.
- Provide **gRPC APIs** for the main chat service to call.
- Handle **background jobs** (e.g., delayed email notifications).

### **Project Structure (`/grpc-service`)**

```
grpc-service/
│── cmd/
│    └── main.go       # Starts gRPC server
│
├── proto/             # Protobuf definitions
│    └── chat.proto    # gRPC messages & services
│
├── internal/
│    ├── spam/         # Spam filtering logic
│    ├── transform/    # Text transformations (markdown, links)
│    └── jobs/         # Background workers (email notifications)
│
└── pkg/
     ├── logger/       # Logging utility
     └── config/       # Config loader
```

### **Core Components**

- `chat.proto`: Defines the gRPC methods.
- `spam/`: Implements message filtering (e.g., block spam, detect offensive words).
- `transform/`: Applies text transformations (e.g., markdown, auto-linking).
- `jobs/`: Manages delayed email notifications when a message is unread.

---

## **Inter-Service Communication**

- The **Chat Service** calls the **gRPC Service** for message processing.
- The **gRPC Service** runs asynchronously and can scale separately.

---

## **Tech Stack & Libraries**

| Feature                    | Library                                        |
| -------------------------- | ---------------------------------------------- |
| WebSocket handling         | `github.com/gorilla/websocket`                 |
| HTTP API (REST)            | `github.com/gin-gonic/gin`                     |
| Database (PostgreSQL)      | `gorm.io/gorm`                                 |
| Redis Pub/Sub              | `github.com/go-redis/redis/v8`                 |
| gRPC Service Communication | `google.golang.org/grpc`                       |
| Authentication (JWT)       | `github.com/golang-jwt/jwt` or some OAuth2 API |
| Background jobs (emails)   | `github.com/hibiken/asynq`                     |

---
