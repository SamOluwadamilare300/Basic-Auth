# Go Auth App

A simple authentication API built with [Fiber](https://gofiber.io/) and MongoDB, designed for user registration and login functionality.

## 📂 Project Structure

```
/go-auth-app
│── /config
│   ├── database.go          # MongoDB connection setup
│── /handlers
│   ├── auth.go              # Authentication handlers (register, login)
│── /models
│   ├── user.go              # User model
│   ├── book.go              # Book model (if needed)
│── /middleware
│   ├── auth_middleware.go   # Middleware for authentication
│── main.go                  # Main entry point
│── go.mod                   # Go module file
│── go.sum                   # Dependency lock file
│── render.yaml              # Render service configuration (optional)
```

## 🚀 Getting Started

### 1️⃣ Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/dl/)
- [MongoDB](https://www.mongodb.com/try/download/community)
- [Git](https://git-scm.com/downloads)

### 2️⃣ Clone the Repository
```sh
git clone https://github.com/YOUR_USERNAME/go-auth-app.git
cd go-auth-app
```

### 3️⃣ Install Dependencies
```sh
go mod tidy
```

### 4️⃣ Configure Environment Variables
Create a `.env` file and add your MongoDB connection string:
```
MONGO_URI=mongodb+srv://your_username:your_password@cluster.mongodb.net/dbname
PORT=3000
```

### 5️⃣ Run the Application
```sh
go run main.go
```
Your API will be running at: `http://localhost:3000`

## 🌍 API Endpoints

| Method | Endpoint       | Description         |
|--------|--------------|---------------------|
| `GET`  | `/`          | Welcome message    |
| `POST` | `/auth/register` | Register a user |
| `POST` | `/auth/login` | User login        |

## 🛠 Deploying to Render

### 1️⃣ Push to GitHub
```sh
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/go-auth-app.git
git push -u origin main
```

### 2️⃣ Deploy on Render
1. Go to [Render](https://dashboard.render.com/)
2. Click **New Web Service** → Connect GitHub repository
3. Set **Environment** to **Go**
4. Add **Environment Variables** (`MONGO_URI`, `PORT`)
5. Click **Deploy**

### 3️⃣ Test Your Deployment
```sh
curl -X GET https://your-app.onrender.com/
```

## 📜 License
This project is licensed under the MIT License.

---
💡 **Need help?** Feel free to open an issue!