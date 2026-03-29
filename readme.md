# 🚀 GitHub Activity CLI (Go)

A simple and powerful Command Line Interface (CLI) tool built in Go to fetch and display recent GitHub activity of any user.

---

## 📌 Features

* 🔍 Fetch recent GitHub user activity
* ⚡ Fast and lightweight (no external libraries)
* 🧠 Clean architecture (cmd, api, models, utils)
* ❌ Graceful error handling
* 📊 Human-readable output
* 🚫 Handles invalid users & API errors
* 🔒 Handles GitHub API rate limits

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Packages Used:**

  * `net/http` → API calls
  * `encoding/json` → JSON parsing
  * `os` → CLI arguments
  * `fmt`, `io`, `strings`

---

## 📂 Project Structure

```
github-activity/
│
├── main.go            # Entry point
├── go.mod
│
├── cmd/
│   └── root.go        # CLI handling
│
├── api/
│   └── github.go      # API calls
│
├── models/
│   └── event.go       # JSON structs
│
├── utils/
│   └── formatter.go   # Output formatting
```

---

Project Url : https://roadmap.sh/projects/github-user-activity

## ⚙️ Installation & Setup

### 1. Clone the repository

```
git clone https://github.com/your-username/github-activity.git
cd github-activity
```

### 2. Initialize Go modules

```
go mod tidy
```

### 3. Run the project

```
go run main.go <github-username>
```

---

## ▶️ Usage

```
go run main.go kamranahmedse
```

### 💡 Example Output

```
📊 Recent Activity for kamranahmedse:

--------------------------------------------------
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
--------------------------------------------------
```

---

## ❗ Error Handling

The CLI handles:

* ❌ Missing username
* ❌ Invalid GitHub user
* 🚫 API rate limit exceeded
* ⚠️ No recent activity

---

## 🌐 GitHub API Used

```
https://api.github.com/users/<username>/events
```

---

## 🚀 Future Improvements

* 🔎 Filter activity by type (`--type=push`)
* 🎨 Add colored CLI output
* 💾 Cache API responses
* 📦 Build executable (`go build`)
* 🔐 Add GitHub authentication (increase rate limits)

---

## 🧠 What I Learned

* Building CLI applications in Go
* Working with REST APIs
* JSON parsing using structs
* Clean project architecture
* Error handling in Go

---

## 🤝 Contributing

Contributions are welcome! Feel free to fork the repo and submit a PR.

---

## 🙌 Acknowledgements

* GitHub API
* Go Documentation
