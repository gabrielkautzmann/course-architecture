# Scalable Go REST API

This is a scalable REST API implemented in Go, designed with modular services and layers.
It supports user management, posts, comments, image upload (mock), voting, and ranking.

## Features

- 🔐 JWT-based Authentication
- 📄 User Signup/Login
- 📝 Post Creation and Retrieval
- 💬 Commenting System
- 👍👎 Voting System
- 📈 Ranking Top Posts (Mocked)
- 🖼️ Image Upload (Mock)

## Tech Stack

- Go (v1.21+)
- GORM (PostgreSQL)
- Gorilla Mux (Router)
- JWT for Auth

## Project Structure

```
.
├── cmd/webapp              # Entry point
├── internal/               # Feature modules
│   ├── users
│   ├── posts
│   ├── comments
│   ├── votes
│   ├── images
│   └── ranking
├── api/                    # Route definitions
├── configs/                # Config loader
├── pkg/database/           # DB connection
├── scripts/                # DB migrations
├── go.mod
└── README.md
```

## Routes

### Users
- `POST /users/create` – Register user
- `POST /users/login` – Login (returns JWT)

### Posts
- `POST /posts` – Create post
- `GET /posts` – List posts
- `GET /posts/{post-id}` – Get post
- `DELETE /posts/{post-id}` – Delete post

### Comments
- `POST /posts/{post-id}/comments` – Add comment
- `GET /posts/{post-id}/comments` – List comments

### Images
- `POST /posts/{post-id}/images` – Upload image (mocked)

### Votes
- `POST /posts/{post-id}/vote` – Vote on post

### Ranking
- `GET /ranking/top` – Top posts (mocked)

## Running

### 1. Setup
Create `.env` file:
```
DB_DSN=host=localhost user=postgres password=postgres dbname=goapi port=5432 sslmode=disable
JWT_SECRET=mysecret
```

### 2. Initialize DB
```
psql -U postgres -d goapi -f scripts/init_db.sql
```

### 3. Run
```
go run cmd/webapp/main.go
```

## License
MIT
