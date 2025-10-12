# URL Shortener

A simple, fast, and efficient URL shortening service written in Go.

This project is a minimalistic API for creating short links and redirecting to the original URLs. It is ideal for learning the basics of creating web services in Go.

## 🚀 Features

*   **Short URL creation:** Send a long URL and receive a short code.
*   **Redirection:** Use the short code to redirect to the original URL.
*   **Graceful Shutdown:** Correctly shut down the server without interrupting active requests.
* **In-memory storage:** (or specify your own database: Redis, PostgreSQL) Simple data storage for a quick start.
* **Minimal dependencies:** Written using the standard Go library.

## ⚙️ Installation

### Requirements
- Go 1.22+

### Steps

1.  **Clone repo:**
    ```bash
    git clone https://github.com/cobrich/url-shortener.git
    cd url-shortener
    ```

2.  **Build project:**
    ```bash
    # This will create an executable named 'url-shortener' (on Linux/macOS)
    # or 'url-shortener.exe' (on Windows)
    go build
    ```
    *or just run for development:*
    ```bash
    go run .
    ```

## ▶️ Usage

After running server can be available on address `http://localhost:8080`.

### 1. Creating short url

Send `POST` request to endpoint `/shorten` with your URL in request body.

**Request:**
```bash
curl -X POST http://localhost:8080/shorten \
-H "Content-Type: application/json" \
-d '{"url": "https://www.google.com/search?q=how+to+write+a+good+readme"}'
```

**Respone:**
```json
{
  "short_code": "http://localhost:8080/aB3xZ9"
}
```

### 2. Redirecting to original URL

Just navigate with generated short link in browser or use `curl`:

```bash
curl -vL http://localhost:8080/aB3xZ9
```
You would redirected (`302 Found`) to original URL.

## 🔧 Configuration

The service can be configured using environment variables:

| Variable | Definition                | Default value |
|------------|-------------------------|-----------------------|
| `PORT`     | Порт, на котором работает сервер | `8080`                |

<!-- 
## 🧪 Запуск тестов

Для запуска всех тестов в проекте выполните команду:

```bash
go test ./... -v
``` -->

## 📖 API Endpoints

### `POST /shorten`
Creates a new short URL.

*   **Request Body:**
    ```json
    {
      "url": "https://your-long-url.com/goes/here"
    }
    ```
*   **Success Response (201 Created):**
    ```json
    {
      "short_code": "aB3xZ9"
    }
    ```
*   **Error Responses:**
    *   `400 Bad Request`: If the JSON is invalid or the URL is not reachable.
    *   `500 Internal Server Error`: If a unique code cannot be generated.

### `GET /{short_code}`
Redirects to the original long URL.

*   **Example:** `GET /aB3xZ9`
*   **Success Response:** `302 Found` redirect to the original URL.
*   **Error Response:** `404 Not Found` if the code does not exist.

## 🤝 Contributing to the project

I always welcome help! If you want to improve the project:
1.  Fork the repository.
2.  Create a new branch (`git checkout -b feature/AmazingFeature`).
3.  Make your changes.
4.  Commit (`git commit -m ‘Add some AmazingFeature’`).
5.  Push your changes to your fork (`git push origin feature/AmazingFeature`).
6.  Create a Pull Request.

Please report any bugs in the [Issues](https://github.com/cobrich/url-shortener/issues) section.

## 📄 License

This project is distributed under the MIT license. For details, see the `LICENSE` file.