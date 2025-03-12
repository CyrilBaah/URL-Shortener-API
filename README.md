# URL Shortener API

A simple URL shortener API built with Golang, Gorilla Mux, and Docker. This API allows users to shorten long URLs and redirect using generated short codes.

##  Features
- Shorten long URLs
- Redirect using short URLs
- In-memory storage for simplicity


##  Installation & Running

### 1. Clone the Repository
```bash
git clone https://github.com/CyrilBaah/URL-Shortener-API.git
cd URL-Shortener-API
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Run the Server
```bash
go run main.go
```

The server will start on **http://localhost:8080**.

##  API Endpoints

### Shorten a URL
**Endpoint:** `POST /shorten`

**Request:**
```json
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
```

**Response:**
```json
{
  "short_url": "Pwl3OV"
}
```

### Redirect to Original URL
**Endpoint:** `GET /{short_code}`

**Example:**
```bash
curl -L http://localhost:8080/Pwl3OV
```
Redirects to `https://example.com`.

## Running with Docker

### 1. Build the Docker Image
```bash
docker build -t url-shortener-api .
```

### 2. Run the Container
```bash
docker run -p 8080:8080 url-shortener-api
```

