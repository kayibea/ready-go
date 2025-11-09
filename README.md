# Go Learning Roadmap (Standard Library Only)

## 1. Hello CLI / Basic Calculator
**Goal:** Command-line calculator that adds, subtracts, multiplies, and divides two or more numbers.

**Directives:**
- Accept any number of numeric arguments from `os.Args`.
- Convert all arguments to `float64`.
- Perform:
  - Sum (add all)
  - Subtraction (left-to-right)
  - Multiplication (multiply all)
  - Division (left-to-right)
- Print results neatly aligned with `fmt.Printf`.
- Handle division by zero and invalid inputs gracefully.

**Focus:** `fmt`, `os`, `strconv`, `path`

---

## 2. File Line Counter
**Goal:** Count lines, words, and characters in a text file.

**Directives:**
- Accept a filename via CLI argument.
- Open file with `os.Open`.
- Read using `bufio.Scanner`.
- For each line:
  - Increment line count.
  - Count words using `strings.Fields`.
  - Count characters using `len(line)`.
- Print total lines, words, and characters.

**Focus:** `os`, `bufio`, `strings`, `fmt`

---

## 3. Simple REST API
**Goal:** Basic task management API (in-memory).

**Directives:**
- Define a `Task` struct with ID, Title, Done.
- Implement endpoints:
  - `GET /tasks` → list all
  - `POST /tasks` → create
  - `DELETE /tasks/{id}` → delete
- Store tasks in a global slice.
- Encode/decode JSON manually.

**Focus:** `net/http`, `encoding/json`, `strconv`, `fmt`

---

## 4. Todo CLI App with Persistence
**Goal:** CLI tool to manage tasks stored in a JSON file.

**Directives:**
- Commands: `add`, `list`, `done <id>`, `delete <id>`.
- Tasks stored in a slice of structs.
- Read/write JSON file for persistence.
- Load file at start, save after each change.

**Focus:** `encoding/json`, `os`, `ioutil` or `os.WriteFile`, `fmt`

---

## 5. Concurrent Downloader
**Goal:** Download multiple URLs at once.

**Directives:**
- Read URLs from a text file or arguments.
- Start a goroutine per download.
- Save each response to `<index>.html`.
- Use `sync.WaitGroup` to wait for all.
- Handle network errors cleanly.

**Focus:** `net/http`, `os`, `io`, `sync`, `fmt`

---

## 6. Chat Server
**Goal:** TCP-based chat server and client.

**Directives:**
- Server:
  - Accept TCP connections with `net.Listen`.
  - Handle each client in a goroutine.
  - Broadcast messages to all connected clients.
- Client:
  - Connect with `net.Dial`.
  - Read input and send to server.
  - Print received messages.

**Focus:** `net`, `bufio`, `fmt`, `os`, `sync`

---

## 7. Web Scraper
**Goal:** Extract all links from a web page.

**Directives:**
- Fetch HTML via `http.Get`.
- Parse HTML manually using `golang.org/x/net/html` (skip if avoiding external).
  - Alternative: simple `strings.Contains` and regex for `href="..."`.
- Collect and print unique URLs.

**Focus:** `net/http`, `strings`, `regexp`, `fmt`

---

## 8. JWT-Based Authentication API (No external deps)
**Goal:** Basic token authentication simulation (custom JWT-like).

**Directives:**
- Register and login endpoints.
- On login, generate a fake token (e.g., `base64(user:timestamp)`).
- Store user credentials in memory.
- Middleware checks for "Authorization" header with a valid token.
- **Note:** Do NOT use real JWT libs (use `encoding/base64`, `time`).

**Focus:** `net/http`, `encoding/base64`, `time`, `fmt`

---

## 9. Mini ORM (Conceptual, SQLite-less)
**Goal:** Mimic simple ORM using structs and reflection.

**Directives:**
- Create struct tags like `db:"column_name"`.
- Write a function to print SQL-like statements:
  - `INSERT INTO table (...) VALUES (...)`
  - `SELECT * FROM table;`
- Use `reflect.Type` and `reflect.Value` to read struct fields.

**Focus:** `reflect`, `fmt`, `strings`

---

## 10. Real-Time Web App (Pseudo WebSocket)
**Goal:** Simulate real-time updates over long-polling (no external libs).

**Directives:**
- Server:
  - `GET /poll` → holds connection for a few seconds before replying.
  - `POST /message` → broadcast new messages to waiting clients.
- Clients send new requests every few seconds.
- Manage connected clients with goroutines and channels.

**Focus:** `net/http`, `sync`, `time`, `fmt`

---

## Notes
- Use **only the Go standard library** for every project.
- Avoid any third-party dependencies.
- Each project should teach a **core Go concept**:
  - CLI parsing
  - File I/O
  - JSON and HTTP
  - Concurrency
  - Networking
  - Reflection
- Always format your code with `go fmt` and document key functions.

---
