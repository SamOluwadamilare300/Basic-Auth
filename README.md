# 📚 Go Library API

A robust Library API built with [Fiber](https://gofiber.io/) and MongoDB, designed for user authentication, book management, and borrowing functionality.

## 📂 Project Structure

```
/go-library-api
│── /config
│   ├── database.go          # MongoDB connection setup
│── /handlers
│   ├── auth.go              # Authentication handlers (register, login)
│   ├── book.go              # Book management handlers
│   ├── borrow.go            # Borrow and return book handlers
│── /models
│   ├── user.go              # User model
│   ├── book.go              # Book model
│── /middleware
│   ├── auth_middleware.go   # Middleware for authentication
│── main.go                  # Main entry point
│── go.mod                   # Go module file
│── go.sum                   # Dependency lock file
│── render.yaml              # Render service configuration
│── .env                     # Environment variables
```

## 🚀 Features
- **User Authentication** (Register, Login, JWT Auth)
- **Role-based Access Control** (Admin/User)
- **Book Management** (Add, View, Update, Delete Books)
- **Borrowing & Returning Books**
- **User Book Records** (Reading List, Completed Books)
- **Secure Middleware & Rate Limiting**

## 🛠 Getting Started

### 1️⃣ Prerequisites
Ensure you have the following installed:
- [Go](https://go.dev/dl/)
- [MongoDB](https://www.mongodb.com/try/download/community)
- [Git](https://git-scm.com/downloads)

### 2️⃣ Clone the Repository
```sh
git clone https://github.com/YOUR_USERNAME/go-library-api.git
cd go-library-api
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
JWT_SECRET=your_secret_key
```

### 5️⃣ Run the Application
```sh
go run main.go
```
Your API will be running at: `http://localhost:3000`

## 🌍 API Endpoints

### 🔐 Authentication
| Method | Endpoint           | Description |
|--------|-------------------|-------------|
| `POST` | `/auth/register`  | Register a new user |
| `POST` | `/auth/login`     | User login |

### 📚 Books
| Method | Endpoint        | Description |
|--------|----------------|-------------|
| `POST` | `/books`       | Add a new book |
| `GET`  | `/books`       | Fetch all books |
| `GET`  | `/books/:id`   | Get a single book |
| `PUT`  | `/books/:id`   | Update book details |
| `DELETE` | `/books/:id` | Remove a book |

### 🔄 Borrowing Books
| Method | Endpoint       | Description |
|--------|--------------|-------------|
| `POST` | `/borrow`    | Borrow a book |
| `POST` | `/return`    | Return a book |

## 🛠 Deploying to Render

### 1️⃣ Push to GitHub
```sh
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/go-library-api.git
git push -u origin main
```

### 2️⃣ Deploy on Render
1. Go to [Render](https://dashboard.render.com/)
2. Click **New Web Service** → Connect GitHub repository
3. Set **Environment** to **Go**
4. Add **Environment Variables** (`MONGO_URI`, `PORT`, `JWT_SECRET`)
5. Click **Deploy**

### 3️⃣ Test Your Deployment
```sh
curl -X GET https://your-app.onrender.com/
```

## 📜 License
This project is licensed under the MIT License.

---
💡 **Need help?** Feel free to open an issue!

