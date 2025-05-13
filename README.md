# Flash 🕸️

**Flash** is a fast and lightweight web crawler written in Go. It recursively visits web pages, extracts links, and can be configured for SEO auditing, content discovery, or data collection.

![Flash Web Crawler](https://img.shields.io/badge/Flash-Web%20Crawler-orange)
![Go Version](https://img.shields.io/badge/Go-1.24-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## 🚀 Features

- ⚡ **Fast Concurrent Crawling**: Efficiently processes multiple pages
- 🌐 **Recursive URL Discovery**: Automatically follows links to explore websites
- 🧩 **Simple & Modular**: Easy to extend with your own functionality
- 📊 **URL Normalization**: Handles various URL formats consistently
- 🐳 **Docker Support**: Ready for containerized deployments

## 📋 Requirements

- [Go](https://golang.org/dl/) 1.24 or later
- [Docker](https://www.docker.com/) (optional for containerization)

## 🛠️ Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/sudonitj/Flash.git
cd Flash

# Install dependencies
go mod tidy

# Build the application
go build -o flash
```

### With Docker

```bash
# Build the Docker image
docker build -t flash-crawler .
```

## 💻 Usage

### Basic Usage

```bash
# Run from compiled binary
./flash https://example.com

# Or using go run
go run main.go https://example.com
```

### Docker Usage

```bash
# Run with Docker
docker run --rm flash-crawler https://example.com
```

## 📂 Project Structure

```
flash/
├── main.go                  # Entry point
├── crawler/                 # Core crawler package
│   ├── crawler.go           # Crawling logic
│   ├── get_urls.go          # HTML parsing for URLs
│   ├── normalize_url.go     # URL normalization
│   ├── get_url_test.go      # Tests for URL extraction
│   └── normalize_url_test.go# Tests for normalization
├── go.mod                   # Module definition
├── go.sum                   # Dependency checksums
├── Dockerfile               # Docker configuration
└── README.md                # This file
```

## ⚙️ How It Works

1. The crawler starts with a given URL (the seed)
2. It visits the page, extracts all links from the HTML
3. Links are normalized to prevent duplicates
4. Each discovered URL is added to the queue if not already visited
5. The process repeats until no more URLs are left to visit

Flash uses the standard Go library for HTTP requests and HTML parsing, making it lightweight with minimal dependencies.

## 🧪 Running Tests

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./crawler
```

## 🔄 Extending Flash

Flash is designed to be modular. You can extend its functionality by:

1. Adding more analysis during crawling
2. Implementing depth limits
3. Adding domain filtering
4. Implementing rate limiting
5. Adding data extraction capabilities

## 📊 Example Use Cases

- **SEO Auditing**: Find broken links and analyze site structure
- **Content Discovery**: Map out all accessible pages on a website
- **Data Collection**: Extract specific types of content from pages
- **Site Monitoring**: Track changes to pages over time

## 🛡️ Responsible Usage

When using Flash, please:

- Respect website terms of service
- Consider adding rate limiting to avoid overloading servers
- Add support for robots.txt to respect crawling directives
- Only crawl sites you have permission to access

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- Go community for excellent HTTP and HTML parsing libraries
- All contributors who help improve Flash