# 🗓️ Appointment Scheduling Platform with Google Calendar Integration

This is a full-featured appointment scheduling platform built in Go, with Google Calendar integration via OAuth 2.0. Users can authenticate using their Google accounts, manage availability, and automatically sync events to their Google Calendar.

---

## 🚀 Features

- 🔐 Secure authentication via Google OAuth 2.0  
- 📆 Real-time Google Calendar integration (read/write)  
- ✅ Custom time slot availability and appointment booking  
- 📊 Personal dashboard (upcoming appointments, availability overview)  
- 🔧 RESTful API built with Gin  
- ⚙️ Scalable and modular architecture  
- 🐳 Docker-ready (optional for deployment)

---

## 🧰 Tech Stack

| Layer     | Technology                         |
|-----------|------------------------------------|
| Backend   | Go, Gin, OAuth2                    |
| API       | Google Calendar API, Google User Info API |
| Database  | PostgreSQL or MongoDB (pluggable)  |
| Auth      | Google OAuth 2.0 (JWT planned)     |
| Deploy    | Docker (optional), .env config     |

---

## 📸 Screenshots

_(coming soon)_ — Calendar UI, availability management, dashboard summary.

---

## 🔧 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/YOUR_USERNAME/appointment-platform.git
cd appointment-platform

### 2. Copy the envs to the project root, as the example bellow

```bash
    APP_ENV=local
    APP_PORT=8080
    JWT_SECRET=your_jwt_secret
    MONGO_DSN=mongodb://mongo:27017
    MONGO_DB=appointment-platform
    SEND_GRID_API_KEY=your_sendgrid_api_key
    GOOGLE_CLIENT_ID=your_google_client_id
