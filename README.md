# 📝 Go-Gin Notes API

Welcome! This is a clean, modular REST API I built using **Golang** to manage notes. I focused on creating a scalable structure that separates logic into repositories, handlers, and routes for better maintainability.

---

## 🛠 The Tech Stack
* **Language:** [Go (Golang)](https://golang.org/)
* **Framework:** [Gin Gonic](https://gin-gonic.com/) (High-performance routing)
* **Database:** [MongoDB](https://www.mongodb.com/) (NoSQL storage)
* **Live Reload:** [Air](https://github.com/cosmtrek/air) (For a smooth development workflow)
* **Config:** Environment-based configuration via `.env`

---

## 📂 Internal Architecture
I've organized this project using a layered approach to keep things clean:
* **`cmd/api/`**: The main entry point where the application starts.
* **`internal/notes/`**: The heart of the app:
    * `note_model.go`: The blueprint for our note data.
    * `notes_repo.go`: Handles all the MongoDB database queries.
    * `notes_handler.go`: The middleman that processes requests and sends back JSON.
    * `notes_routes.go`: Where the `/notes` endpoints are defined.
* **`server/`**: Sets up the Gin router and server engine.

---

## 🚀 How to Run Locally

### 1. Clone the repository
```bash
git clone [https://github.com/shubhtrek/go-gin-mongodb-api.git](https://github.com/shubhtrek/go-gin-mongodb-api.git)
cd go-gin-mongodb-api
```

## 📝 Example Request (POST /notes)
```json
{
    "title": "My First Note",
    "content": "This is a note saved in MongoDB!",
    "tags": ["golang", "gin"]
}
