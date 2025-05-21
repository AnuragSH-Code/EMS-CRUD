# EMS-CRUD Backend

A simple Employee Management System (EMS) backend built with Go and MongoDB.

---

## 📦 Prerequisites

Make sure you have the following installed before running the application:

- [Go](https://golang.org/dl/) (version 1.20 or higher)
- [Git](https://git-scm.com/)
- [MongoDB Shell (`mongodbsh`)](https://www.mongodb.com/try/download/shell) installed and running on default port `27017`

> **Note:** Ensure MongoDB is running locally on `mongodb://localhost:27017`

---

## 🚀 Getting Started

Follow these steps to set up and run the backend server:

### 1. Clone the Repository

```bash
git clone https://github.com/AnuragSH-Code/EMS-CRUD.git
```
Replace `<REPO_URL>` with the actual URL of this repository.  
### 2. Navigate to the Backend Directory

```bash
cd EMS-CRUD/backend/
```

### 3. Run the Application

```bash
go run ./cmd/api/*.go
```

> This command starts the Go backend server.

---

## 🛠 Project Structure

```text
EMS-CRUD/
├── backend/
│   ├── cmd/
│   │   ├── api/
│   │   │   └── *.go       # Entry point files for the API server
│   ├── internal/storage   # Internal application packages
│   └── go.mod             # Go module file
```

---

## 🧪 Environment Configuration

If your project supports environment variables, you can create a `.env` file in the `backend/` directory:

```env
MONGO_URI=mongodb://localhost:27017
```

---

## 📡 API Endpoints

Below are the available HTTP routes for the EMS backend:

### ➕ Create Employee

**POST**  
`http://localhost:8080/v1/employees`

**Request Body:**

```json
{
  "firstname": "Harshit",
  "lastname": "Taneja",
  "role": "Software Engineer",
  "department": "Contacts Tribe",
  "email": "mail@mail.in",
  "contact_no": "1234567890",
  "manager": "Harshit Sangwan"
}
```

Creates a new employee record.

---

### 📥 Get Employees (Paginated)

**GET**  
`http://localhost:8080/v1/employees?limit=2&offset=0`

Returns a paginated array of all employee records in JSON format.

---

### ✏️ Update Employee

**PUT**  
`http://localhost:8080/v1/employees?id=${id}`

**Request Body:**

```json
{
  "email": "mail@mail.in"
}
```

Updates the employee record with the specified ID.

---

### ❌ Delete Employee

**DELETE**  
`http://localhost:8080/v1/employees?id=${id}`

Deletes the employee record with the specified ID.

---

---

## ✅ Verifying MongoDB

Ensure MongoDB is up and running using the shell:

```bash
mongodbsh
```

Then run:

```mongodb
show dbs
```

You should see your stud database listed once the app has inserted data.

---
