# 🚀 [FEATURE] Scaffold Production-Ready Golang API

## 📌 Background

Dibutuhkan sebuah scaffolding backend service menggunakan Golang yang siap digunakan untuk environment production dengan karakteristik:

* High traffic
* Deployable di Kubernetes
* Menggunakan SQL Server sebagai database utama
* Menggunakan Redis untuk caching / session
* Terintegrasi dengan structured logging (ELK-ready)

Tujuan utama adalah menyediakan baseline project yang **clean, scalable, dan maintainable** untuk pengembangan lebih lanjut.

---

## 🎯 Objectives

Membuat project REST API Golang dengan:

* Clean Architecture (handler → service → repository)
* Dependency injection sederhana (tanpa framework DI)
* Production-ready setup (timeout, logging, graceful shutdown)
* SQL-first data access (bukan full ORM)

---

## 🧰 Tech Stack (Mandatory)

* HTTP Server: `net/http`
* Router: `github.com/go-chi/chi/v5`
* Database:

  * `database/sql`
  * SQL Server driver: `github.com/denisenkom/go-mssqldb`
  * Query helper: `github.com/jmoiron/sqlx`
* Redis: `github.com/go-redis/redis/v8`
* Logger: `go.uber.org/zap`

---

## 🏗️ Project Structure

```id="1kdnuh"
cmd/api/main.go
internal/handler/
internal/service/
internal/repository/
internal/repository/sqlserver/
internal/model/
internal/middleware/
pkg/logger/
pkg/redis/
pkg/response/
config/
routes/
```

---

## ⚙️ Functional Requirements

### 1. HTTP Server

* Gunakan `http.Server`
* Wajib konfigurasi:

  * ReadTimeout
  * WriteTimeout
  * IdleTimeout
* Implement graceful shutdown (SIGTERM & SIGINT)

---

### 2. Routing

Implement endpoint berikut:

* `GET /health`
* `GET /users`
* `POST /users`

---

### 3. Middleware

Wajib implement:

* Logging (request path + duration)
* Recovery (panic handler)
* Timeout (gunakan context.WithTimeout)

---

### 4. Database (SQL Server)

* Gunakan `database/sql` + `sqlx`
* Implement connection pooling
* Semua query wajib menggunakan `context.Context`
* Buat repository untuk entity `User`

---

### 5. Redis

* Setup Redis client
* Implement fungsi dasar:

  * Get
  * Set
* Gunakan context di setiap operasi

---

### 6. Logging

* Gunakan `zap` (JSON format)
* Minimal logging:

  * Request masuk
  * Error
  * Startup & shutdown

---

### 7. Response Format

```json id="nnd1c4"
{
  "data": {},
  "error": ""
}
```

---

### 8. Kubernetes Readiness

* Endpoint `/health` wajib tersedia
* Graceful shutdown harus berjalan dengan benar

---

### 9. Docker

* Gunakan multi-stage build
* Final image menggunakan alpine (atau minimal base image)

---

## 🗄️ Data Access Strategy (SQL-first, bukan ORM penuh)

### Pendekatan:

Gunakan SQL eksplisit (bukan ORM seperti GORM)

### Library:

* `github.com/jmoiron/sqlx`

---

### Requirements:

* Semua query ditulis manual (SQL eksplisit)
* Gunakan struct mapping (`db` tag)
* Gunakan parameter binding (hindari SQL injection)
* Gunakan context di semua query

---

### Contoh Query yang Wajib Ada:

* SELECT users
* INSERT user

---

### Catatan:

* Tidak menggunakan ORM seperti GORM
* Fokus pada performa, kontrol query, dan transparansi

---

## 🧪 Unit Testing

### 🎯 Tujuan

Memastikan setiap layer dapat diuji secara terisolasi dan deterministic.

---

### 🧰 Tools

* Testing: `testing` (standard library)
* Assertion (opsional): `github.com/stretchr/testify`
* Mocking:

  * manual mock
  * atau `testify/mock`

---

### 📦 Scope Testing

#### 1. Service Layer (WAJIB)

* Test semua business logic
* Gunakan mock repository
* Tidak boleh akses DB langsung

#### 2. Handler Layer

* Test HTTP response
* Gunakan `httptest`

#### 3. Repository Layer (Optional)

* Boleh integration test (gunakan DB test container)

---

### 📌 Rules

* Test harus deterministic
* Tidak boleh bergantung external service
* Semua dependency harus bisa di-mock
* Gunakan context

---

### ⚠️ Non-Functional Requirement

* Test dapat dijalankan dengan:

```bash id="8e63cq"
go test ./...
```

* Tidak boleh ada flaky test

---

### ⭐ Optional

* [ ] Coverage minimal 60%
* [ ] Table-driven test
* [ ] Benchmark test (`go test -bench`)

---

## ⚠️ Non-Functional Requirements

* Semua function harus menerima `context.Context`
* Tidak boleh menggunakan global variable untuk dependency
* Gunakan dependency injection sederhana
* Tidak menggunakan framework berat (Gin, Fiber, dll)
* Project harus bisa dijalankan dengan:

```bash id="mb88q5"
go run ./cmd/api
```

* Tidak boleh ada error compile

---

## 📦 Deliverables

* [ ] Struktur folder lengkap
* [ ] Implementasi semua file (bukan pseudo-code)
* [ ] `go.mod`
* [ ] `.env.example`
* [ ] `Dockerfile`
* [ ] Contoh cara menjalankan project
* [ ] Unit test minimal untuk service & handler

---

## 🧪 Acceptance Criteria

* Project berhasil build tanpa error
* Endpoint `/health` dapat diakses
* Endpoint `/users` berfungsi (GET & POST)
* Logging muncul dalam format JSON
* Graceful shutdown berjalan saat SIGTERM
* Redis & DB client berhasil diinisialisasi
* Unit test dapat dijalankan (`go test ./...`)

---

## 🚫 Anti-Pattern yang Harus Dihindari

* Menggunakan ORM berat (GORM, dll)
* Tidak menggunakan context
* Query tidak eksplisit (hidden query)
* Tidak ada timeout
* Test langsung ke DB untuk unit test
* Hardcoded delay / sleep di test

---

## 🚀 Next Phase (Out of Scope PR Ini)

Akan ditambahkan pada issue terpisah:

* Rate limiter berbasis Redis
* Circuit breaker
* Validation request
* Distributed tracing (OpenTelemetry)
* Authentication (JWT / mTLS)
* Integration test dengan Docker

---
