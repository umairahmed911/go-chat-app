# 💬 Go ChatApp – Real-Time Web-Based Chat Application

A lightweight, real-time web chat application built using **Go (Golang)**, **Gorilla WebSocket**, and standard **HTML/CSS/JavaScript**. This project demonstrates efficient handling of concurrent WebSocket connections, leveraging Go’s native concurrency support to deliver fast and scalable real-time communication.

---

## 🚀 Features

- ⚡ Real-time bi-directional messaging with WebSockets
- 🌐 Simple web-based interface using HTML/CSS/JS
- 🔄 Efficient message broadcasting using Go goroutines and channels
- 👥 Handles multiple users concurrently
- 💻 Lightweight – no database, no Docker, no external dependencies

---

## 🧱 Project Structure

```

go-chatapp/
├── main.go         # Go WebSocket server
├── static/
│   └── index.html  # Chat UI
├── go.mod          # Go module file
└── README.md       # Project documentation

````

---

## 📦 Tech Stack

- **Backend**: Go (Golang)
- **WebSocket Library**: [Gorilla WebSocket](https://github.com/gorilla/websocket)
- **Frontend**: HTML, CSS, Vanilla JavaScript
- **Architecture**: WebSocket server + static frontend

---

## ⚙️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-chatapp.git
cd go-chatapp
````

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

The server will be available at:
**[http://localhost:8080](http://localhost:8080)**

---

## 💬 How It Works

1. Client opens the web page and connects to the WebSocket server.
2. Server registers the new client and maintains the connection.
3. Messages are sent to the server and broadcasted to all connected clients.
4. Frontend updates messages in real-time using JavaScript and WebSocket events.

---

## 🖼️ Screenshots

![image1](chatApp01.png)
![image2](chatApp02.png)


* Live chat interface
* Console view of server logs

---

## 🔐 Future Enhancements

* ✅ User authentication & login
* 💾 Message history with persistent storage (e.g., SQLite, PostgreSQL)
* 🔐 Private 1-to-1 messaging support
* ☁️ Cloud deployment (Heroku, Fly.io, etc.)
* 🎨 Modern UI with TailwindCSS or React

---

## 📚 References

* [Go Programming Language](https://golang.org/)
* [Gorilla WebSocket Library](https://github.com/gorilla/websocket)
* [MDN WebSocket API](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)

---

## 👨‍💻 Author

**Umair Ahmed**
BS Artificial Intelligence Student


---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## ⭐ Show Your Support

If you found this project helpful, please give it a ⭐ and share it with others!
