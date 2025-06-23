# 🌐 Ghi chú về URL trong Go (`net/url`)

Go cung cấp package `net/url` để làm việc với URL một cách hiệu quả. Struct chính là `url.URL`, chứa nhiều trường và phương thức để phân tích, truy vấn, hoặc chỉnh sửa URL.

---

## 🧩 Các trường (Fields) trong `url.URL`

| Trường     | Giải thích                                                 |
| ---------- | ---------------------------------------------------------- |
| `Scheme`   | Trả về giao thức, ví dụ: `http`, `https`, `ftp`.           |
| `Host`     | Trả về `hostname[:port]`, ví dụ: `example.com:8080`.       |
| `RawQuery` | Trả về chuỗi query (chưa parse), ví dụ: `id=123&sort=asc`. |
| `Path`     | Trả về đường dẫn, ví dụ: `/products/item1`.                |
| `Fragment` | Trả về phần sau dấu `#`, ví dụ: `top`.                     |

---

## ⚙️ Các phương thức (Methods) hữu ích

| Phương thức  | Giải thích                                           |
| ------------ | ---------------------------------------------------- |
| `Hostname()` | Trả về phần hostname (không có port).                |
| `Port()`     | Trả về port nếu có (chuỗi).                          |
| `Query()`    | Trả về `map[string][]string` chứa các tham số query. |
| `User()`     | Trả về thông tin người dùng nếu có (`user:pass@`).   |
| `String()`   | Trả lại URL hoàn chỉnh dạng string.                  |

---

## ✅ Ví dụ minh họa

```go
package main

import (
    "fmt"
    "net/url"
)

func main() {
    rawURL := "https://user:pass@example.com:8080/path?x=1&y=2#section"
    u, err := url.Parse(rawURL)
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }

    fmt.Println("Scheme:", u.Scheme)          // https
    fmt.Println("Host:", u.Host)              // example.com:8080
    fmt.Println("Hostname:", u.Hostname())    // example.com
    fmt.Println("Port:", u.Port())            // 8080
    fmt.Println("Path:", u.Path)              // /path
    fmt.Println("RawQuery:", u.RawQuery)      // x=1&y=2
    fmt.Println("Query x:", u.Query().Get("x")) // 1
    fmt.Println("Fragment:", u.Fragment)      // section
    fmt.Println("User:", u.User.Username())   // user
    fmt.Println("URL as string:", u.String())  // In lại toàn bộ URL
}
```

---

## 💡 Ghi chú thêm:

- Dùng `Query().Get("key")` để truy cập giá trị nhanh.
- Hàm `url.Parse(...)` rất quan trọng khi xử lý các input từ web, API.
- Có thể chỉnh sửa URL bằng cách gán lại các trường rồi dùng `u.String()` để tái tạo.

---

## 📘 `http.ResponseWriter` Methods in Go

| Method                  | Description                                                                                                                                  |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `Header()`              | Trả về `Header`, là alias của `map[string][]string`, cho phép thiết lập các header của phản hồi.                                             |
| `WriteHeader(code int)` | Thiết lập mã trạng thái HTTP cho phản hồi. Thường dùng với các hằng số trong gói `net/http` như `http.StatusOK`, `http.StatusNotFound`, v.v. |
| `Write(data []byte)`    | Ghi dữ liệu vào nội dung phản hồi. Hàm này thực thi interface `Writer`.                                                                      |

### 📌 Ví dụ sử dụng:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello"}`))
}
```

### 💡 Ghi chú:

- Gọi `WriteHeader` **trước** `Write`. Nếu không, status code mặc định là 200.
- `Header()` phải được dùng **trước khi gửi nội dung** qua `Write()` để chắc chắn header được gửi đúng.

## 📘 `http.ResponseWriter` Methods in Go

| Method                  | Description                                                                                                                                  |
| ----------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `Header()`              | Trả về `Header`, là alias của `map[string][]string`, cho phép thiết lập các header của phản hồi.                                             |
| `WriteHeader(code int)` | Thiết lập mã trạng thái HTTP cho phản hồi. Thường dùng với các hằng số trong gói `net/http` như `http.StatusOK`, `http.StatusNotFound`, v.v. |
| `Write(data []byte)`    | Ghi dữ liệu vào nội dung phản hồi. Hàm này thực thi interface `Writer`.                                                                      |

### 📌 Ví dụ sử dụng:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello"}`))
}
```

### 💡 Ghi chú:

- Gọi `WriteHeader` **trước** `Write`. Nếu không, status code mặc định là 200.
- `Header()` phải được dùng **trước khi gửi nội dung** qua `Write()` để chắc chắn header được gửi đúng.

---

## 🔧 Các hàm hỗ trợ phản hồi phổ biến trong `net/http`

| Function                         | Description                                                                                                                                                     |
| -------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `http.Error(w, message, code)`   | Thiết lập header với mã trạng thái `code`, thêm `Content-Type: text/plain`, ghi thông báo lỗi vào phản hồi. Cũng thêm header `X-Content-Type-Options: nosniff`. |
| `http.NotFound(w, r)`            | Gọi `Error()` với mã lỗi 404.                                                                                                                                   |
| `http.Redirect(w, r, url, code)` | Gửi phản hồi chuyển hướng (`3xx`) tới `url` với mã trạng thái `code`.                                                                                           |
| `http.ServeFile(w, r, filePath)` | Gửi nội dung của file đến client. Header `Content-Type` được thiết lập dựa trên phần mở rộng của file.                                                          |

### ✅ Ví dụ:

```go
func serveImage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./images/sample.png")
}

func redirectToHome(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/home", http.StatusFound)
}

func notFound(w http.ResponseWriter, r *http.Request) {
    http.NotFound(w, r)
}
```
## 🌐 HTTP Routing & Handler Functions in Go

### 📌 `http.Handle` vs `http.HandleFunc`

| Function                                                                           | Description                                                                                                       |
| ---------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `Handle(pattern string, handler http.Handler)`                                     | Đăng ký một `Handler` cụ thể cho URL khớp với `pattern`. Gọi phương thức `ServeHTTP` của handler khi request đến. |
| `HandleFunc(pattern string, handlerFunc func(http.ResponseWriter, *http.Request))` | Tiện lợi hơn khi handler được định nghĩa dưới dạng hàm, Go sẽ wrap nó lại thành `http.Handler`.                   |

### 🛠 `net/http` Functions for Creating Handlers

| Function                                                                   | Description                                                 |
| -------------------------------------------------------------------------- | ----------------------------------------------------------- |
| `FileServer(root http.FileSystem)`                                         | Trả về `Handler` phục vụ các file tịnh từ thư mục gốc.      |
| `NotFoundHandler()`                                                        | Trả về `Handler` sinh ra phản hồi 404.                      |
| `RedirectHandler(url string, code int)`                                    | Trả về `Handler` chuyển hướng HTTP đến URL mới.             |
| `StripPrefix(prefix string, handler http.Handler)`                         | Loại prefix trong URL rồi chuyển request cho handler khác.  |
| `TimeoutHandler(handler http.Handler, duration time.Duration, msg string)` | Bao bọc handler, sinh phản hồi lỗi nếu quá thời gian xử lý. |

### 📌 Ví dụ sử dụng `HandleFunc`:

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World")
})
```

### 📁 Ví dụ tạo static file server:

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static", fs))
```

### ⏱ Timeout Example

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(3 * time.Second)
    w.Write([]byte("done"))
})
http.Handle("/slow", http.TimeoutHandler(handler, 2*time.Second, "timeout!"))
```

---

Sử dụng các phương thức và handler trên giúp xây dựng hệ thống routing linh hoạt, hiệu quả và an toàn trong Go web server.
