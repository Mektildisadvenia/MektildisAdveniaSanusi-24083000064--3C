# Yeyen Full-Stack Project

This is a full-stack web application skeleton structured around **Vite (React + TS)** and **Golang (Go 1.22)**, adhering to a Clean Architecture pattern, persisting data on **PostgreSQL**, and containerized using **Docker / Docker Compose**.

## Requirements
- Docker and Docker Compose
- Node.js (v20+)
- Go (1.22+)
- `golang-migrate` CLI (if you want to run standalone migrations)

## Architectures

### Backend (Golang)
Located in `/backend`, built around separation of concerns:
- `domain`: Interfaces and structs
- `repository`: Implementation of PostgreSQL databases.
- `usecase`: Business logic implementations.
- `handler/http`: REST API controllers and native Go `net/http` router logic.

### Frontend (Vite)
Located in `/frontend`. Uses React with TypeScript.

## Setup Instructions

### 1. Initialize environments
```bash
cp .env.example .env
```
Check and modify `.env` if needed. Default values should be acceptable for local development.

### 2. Install dependencies (Optional if fully using Docker)
```bash
make install
```

### 3. Running with Docker (Development)
Spin up the `frontend`, `backend`, and `postgres` containers in development mode:
```bash
make dev
```
- The frontend runs at `http://localhost:5173`
- The backend API runs at `http://localhost:8080` (with live reload via Air)
- The database is exposed at `localhost:5432`

### 4. Database Migrations
Migrations will be needed for the backend to run properly. 
Inside the container, or locally, run the migration up scripts:
```bash
make migrate-up
```

### 5. API Endpoints
- **Users**: `/users` (GET, POST), `/users/{id}` (GET, PUT, DELETE)
- **Products**: `/products` (GET, POST), `/products/{id}` (GET, PUT, DELETE)

### 6. Staging / Production Simulation
```bash
make staging
```
Spins up production-ready containers simulating staging.
