## 🌐 HTTP Client trong Go (net/http)

Go cung cấp sẵn các hàm để gửi HTTP request thông qua gói `net/http`. Dưới đây là các hàm phổ biến để gửi request:

---

### 📥 Gửi HTTP Request

#### `http.Get(url)`

- Gửi HTTP GET request tới URL.
- Trả về `*http.Response` và `error`.

```go
resp, err := http.Get("https://example.com")
```

#### `http.Head(url)`

- Gửi HEAD request tới URL.
- Chỉ lấy header, không lấy body.

```go
resp, err := http.Head("https://example.com")
```

#### `http.Post(url, contentType, body)`

- Gửi POST request với body tuỳ chỉnh.
- `contentType` ví dụ: `application/json`, `text/plain`
- `body` là `io.Reader` (thường là `strings.NewReader(...)`)

```go
resp, err := http.Post(
    "https://example.com",
    "application/json",
    strings.NewReader(`{"name": "Nam"}`),
)
```

#### `http.PostForm(url, data)`

- Gửi POST request dạng form `application/x-www-form-urlencoded`
- `data` là `url.Values` hoặc `map[string][]string`

```go
resp, err := http.PostForm("https://example.com", url.Values{
    "username": {"admin"},
    "password": {"123456"},
})
```

---

### 🔍 So sánh `Post` vs `PostForm`

| Tiêu chí         | `http.Post`                                   | `http.PostForm`                                    |
| ---------------- | --------------------------------------------- | -------------------------------------------------- |
| Loại dữ liệu gửi | Tuỳ chọn (`application/json`, `text/xml`,...) | Dạng form `application/x-www-form-urlencoded`      |
| Đầu vào          | `io.Reader` (ví dụ: `strings.NewReader`)      | `url.Values` hoặc `map[string][]string`            |
| Khi nào dùng     | Gửi JSON, XML, file, nội dung tuỳ chỉnh       | Gửi form login, contact, search đơn giản           |
| Content-Type     | Do lập trình viên tự set                      | Tự động set là `application/x-www-form-urlencoded` |
| Tính linh hoạt   | Linh hoạt hơn, tùy chỉnh cao                  | Đơn giản, nhanh gọn                                |

📌 **Tóm lại**:

- Dùng `Post` khi cần gửi định dạng nâng cao như JSON/XML.
- Dùng `PostForm` khi xử lý dữ liệu form HTML đơn giản.

---

### 📝 Ghi nhớ

- Nhớ `defer resp.Body.Close()` sau khi xử lý response.
- Kiểm tra `resp.StatusCode` để biết response thành công hay không.
- Dùng `io.ReadAll(resp.Body)` hoặc `json.NewDecoder(resp.Body)` để đọc dữ liệu trả về.

---

### 🔐 Bảo mật

- Sử dụng HTTPS thay vì HTTP để truyền dữ liệu an toàn.
- Với `POST`, nên dùng JSON (`application/json`) hoặc form đã encode (`application/x-www-form-urlencoded`).

---
