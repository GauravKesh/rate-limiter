
# 🚀 Redis Rate Limiter (Go)

A high-performance **distributed rate limiter** built using Go, Redis, and Lua scripting.
Implements the **Token Bucket algorithm** with atomic operations for accurate and scalable request throttling.

---

## ✨ Features

* ⚡ High-performance rate limiting using Go
* 🧠 Token Bucket algorithm (smooth + burst handling)
* 🔒 Atomic operations using Redis Lua scripts
* 🌐 Middleware-based integration (Gin)
* 📊 Rate limit headers (`X-RateLimit-*`)
* 🔄 Distributed & horizontally scalable
* 🧪 Load testing support
* ⚙️ Configurable via environment variables

---

## 🏗️ Architecture

Client → API (Gin) → Rate Limiter Middleware → Redis (Lua)

* Redis stores token state
* Lua ensures atomic updates
* Middleware enforces limits before request reaches handler

---

## 📁 Project Structure

```
rate-limiter-go/
│
├── cmd/server/main.go
├── internal/
│   ├── config/
│   ├── limiter/
│   ├── handlers/
│   ├── router/
│
├── scripts/loadtest.go
├── docker/docker-compose.yml
├── .env
├── README.md
```

---

## ⚙️ Setup

### 1. Clone repository

```
git clone https://github.com/GauravKesh/rate-limiter
cd rate-limiter
```

---

### 2. Setup environment variables

Create `.env` file:

```
PORT=3000
REDIS_HOST=localhost
REDIS_PORT=6379
RATE_LIMIT_CAPACITY=10
RATE_LIMIT_REFILL=1
```

---

### 3. Start Redis

Using Docker:

```
docker run -p 6379:6379 redis
```

---

### 4. Run server

```
go run cmd/server/main.go
```

Server will start at:

```
http://localhost:3000
```

---

## 🧪 Load Testing

Run:

```
go run scripts/loadtest.go
```

Example output:

```
Total Requests: 200
Success (200): 10
Rate Limited (429): 190
Failed: 0
```

---

## 📡 API Response Headers

| Header                | Description          |
| --------------------- | -------------------- |
| X-RateLimit-Limit     | Max requests allowed |
| X-RateLimit-Remaining | Remaining tokens     |

---

## 🧠 Algorithm

### Token Bucket

* Each user has a bucket with limited tokens
* Tokens refill over time
* Each request consumes a token
* If no tokens → request rejected (429)

Benefits:

* Smooth rate limiting
* Allows bursts
* More accurate than fixed window

---

## 🔧 Configuration

Environment variables:

| Variable            | Description | Default   |
| ------------------- | ----------- | --------- |
| PORT                | Server port | 3000      |
| REDIS_HOST          | Redis host  | localhost |
| REDIS_PORT          | Redis port  | 6379      |
| RATE_LIMIT_CAPACITY | Max tokens  | 10        |
| RATE_LIMIT_REFILL   | Tokens/sec  | 1         |

---

## 🚀 Future Improvements

* 🔑 API key / user-based rate limiting
* 📊 Metrics (Prometheus + Grafana)
* 🌍 Redis Cluster support
* 🧱 Circuit breaker (Redis fallback)
* 🖥️ Admin dashboard (Next.js)
* ⚡ k6 benchmarking

---

## 📜 License

MIT License
