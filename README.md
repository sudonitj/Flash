# Flash 🕸️

**Flash** is a fast and lightweight web crawler written in Go. It recursively visits web pages, extracts links, and can be configured for various use cases like scraping, indexing, or auditing websites.

---

## 🚀 Features

- ⚡ Fast and concurrent crawling
- 🌐 Recursive URL discovery
- ⛔ Optional robots.txt respect
- 📂 Modular and extensible
- 🐳 Docker support for easy deployment

---

## 🛠️ Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.20 or later
- [Docker](https://www.docker.com/) (optional)

---

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/ekas-7/flash.git
cd flash
go mod tidy
```

## 🧪 Running Flash

### ▶️ Locally (Direct Execution)

Run with `go run`:

```bash
go run main.go https://example.com
```

Or build and execute:

```bash
go build -o flash
./flash https://example.com
```

You'll see the crawler visiting pages and printing links and information to the console.

### 🐳 Running Flash with Docker

1. 🛠️ Build the Docker Image

```bash
docker build -t flash-crawler .
```

2. ▶️ Run the Web Crawler

```bash
docker run --rm flash-crawler https://example.com
```

3. ⚙️ Optional Flags (if implemented)

If your crawler accepts flags, you can pass them like:

```bash
docker run --rm flash-crawler --depth=2 --concurrency=5 https://example.com
```

4. 📂 .dockerignore (Optional but Recommended)

To improve build times, create a `.dockerignore` file:

```
.git
*.log
*.tmp
*.md
```

## 📁 Project Structure

```
flash/
├── main.go        # Entry point: starts the crawler
├── crawler.go     # Core crawling logic
├── utils.go       # Helper functions
├── go.mod         # Go module definition
├── Dockerfile     # Docker configuration
├── .dockerignore  # Docker exclusions
└── README.md      # Project documentation
```

## 📄 License

This project is licensed under the MIT License.

