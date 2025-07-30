# 🗓️ Appointment Scheduling Platform

This is a full-featured appointment scheduling platform built in Go and React. Users can authenticate using their email and password, manage availability, create patients for their clinics with address, email, phones and names.

Functionalities: 
  The patient can:
    Reset their password by verifying the code received via email.
    Receive an email 24 hours before their appointment to confirm or cancel it.
  The clinic can:
    Create appointments by linking them to previously registered patients.
    Register new patients for their clinic.


✅ Checkout the demo:
https://appointment-platform-two.vercel.app/login
(Note: Due to inactivity, the first load might be a bit slow)

Deployment: 
Database: Atlas MongoDB
Backend: Render
Frontend: Vercel
---

## 🚀 Features

- 🔐 Secure authentication via JWT
- ✅ Custom time slot availability and appointment booking  
- 📊 Personal dashboard (upcoming appointments, availability overview)  
- 🔧 RESTful API built with Gin  
- ⚙️ Scalable and modular architecture  
- 🐳 Docker-ready (optional for deployment)

---

## 🧰 Tech Stack

| Layer     | Technology                         |
|-----------|------------------------------------|
| Backend   | Go, Gin, SendGrid (email sender)   |
| API       | SendGrid                           |
| Database  | MongoDB                            |
| Auth      | JWT                                |
| Deploy    | Docker                             |

---

## 🔧 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/YOUR_USERNAME/appointment-platform.git
cd appointment-platform
```
### 2. Copy the envs to the project root, as the example bellow
```bash
  APP_ENV=local
  APP_PORT=8080
  JWT_SECRET=your_jwt_secret
  MONGO_DSN=mongodb://mongo:27017
  MONGO_DB=appointment-platform
  SEND_GRID_API_KEY=your_sendgrid_api_key
  GOOGLE_CLIENT_ID=your_google_client_id
```

### 3. Run the build command:
```bash
  docker-compose build
```

### 4. Run the up command:
```bash
  docker-compose up -d
```

### 4. Or, if you prefer, use the Makefile command:
```bash
  make up
```
