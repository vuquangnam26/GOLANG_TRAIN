## 🍪 Cấu trúc Cookie trong Go (net/http)

Go định nghĩa cấu trúc `http.Cookie` để thiết lập và quản lý cookie trong ứng dụng web. Dưới đây là các thuộc tính quan trọng của `http.Cookie`:

---

### 🧱 Thuộc tính của http.Cookie

| Tên Field    | Kiểu dữ liệu        | Mô tả                                                          |
| ------------ | ------------------- | -------------------------------------------------------------- |
| **Name**     | `string`            | Tên của cookie.                                                |
| **Value**    | `string`            | Giá trị của cookie.                                            |
| **Path**     | `string (optional)` | Đường dẫn mà cookie hợp lệ (mặc định là toàn bộ site).         |
| **Domain**   | `string (optional)` | Tên miền mà cookie áp dụng (subdomain, domain).                |
| **Expires**  | `time.Time`         | Thời gian hết hạn (absolute).                                  |
| **MaxAge**   | `int`               | Số giây cookie còn sống (relative). `<=0` nghĩa là xóa cookie. |
| **Secure**   | `bool`              | Chỉ gửi cookie qua HTTPS nếu `true`.                           |
| **HttpOnly** | `bool`              | Ngăn không cho JavaScript truy cập cookie nếu `true`.          |
| **SameSite** | `http.SameSite`     | Chính sách gửi cookie cross-site (CSRF protection).            |

---

### 🔐 Các giá trị của SameSite:

- `http.SameSiteDefaultMode`
- `http.SameSiteLaxMode`
- `http.SameSiteStrictMode`
- `http.SameSiteNoneMode`

---

### 💡 Ví dụ sử dụng cookie trong Go:

```go
http.SetCookie(w, &http.Cookie{
    Name:     "session_id",
    Value:    "abc123",
    Path:     "/",
    HttpOnly: true,
    Secure:   true,
    MaxAge:   3600, // 1 giờ
    SameSite: http.SameSiteLaxMode,
})
```

---

### 📌 Ghi nhớ

- Dùng `Expires` nếu bạn muốn thiết lập thời gian cụ thể.
- `MaxAge` phù hợp khi muốn đặt thời gian sống động (ví dụ: 10 phút tính từ lúc set).
- Kết hợp `Secure + HttpOnly + SameSite` để bảo vệ chống lại XSS và CSRF.

---
## 🍪 Đọc Cookie trong Go (net/http)

Khi nhận request từ client, bạn có thể lấy cookie bằng các phương thức sau:

---

### 📥 Đọc một cookie theo tên

```go
cookie, err := r.Cookie("session_id")
if err != nil {
    if errors.Is(err, http.ErrNoCookie) {
        fmt.Println("Cookie không tồn tại")
    } else {
        fmt.Println("Lỗi đọc cookie:", err)
    }
} else {
    fmt.Println("Giá trị cookie:", cookie.Value)
}
```

* ✅ Trả về con trỏ `*http.Cookie`
* ❌ Trả lỗi nếu không tìm thấy

---

### 📥 Đọc toàn bộ cookie

```go
cookies := r.Cookies()
for _, c := range cookies {
    fmt.Printf("%s = %s\n", c.Name, c.Value)
}
```

* ✅ Trả về `[]*http.Cookie`
* Hữu ích khi cần log hoặc duyệt tất cả cookies trong request

---

### 📝 Ghi nhớ

* `r.Cookie(name)` chỉ lấy 1 cookie theo tên.
* `r.Cookies()` duyệt toàn bộ cookie.
* Cần kiểm tra lỗi `http.ErrNoCookie` khi dùng `r.Cookie(name)`.

---

### 🔐 Bảo mật

* Không nên dùng cookie để lưu dữ liệu nhạy cảm nếu không có mã hóa.
* Dùng `HttpOnly`, `Secure`, `SameSite` để tăng bảo vệ chống tấn công XSS và CSRF.

---
