# README: Log Package in Go

Gói `log` trong Go được sử dụng để ghi log trong quá trình chạy chương trình.

## 1. Cấu hình log

### SetPrefix(prefix string)

- Gán prefix (tiền tố) cho mỗi dòng log.

```go
log.SetPrefix("[INFO] ")
```

### SetFlags(flags int)

- Cấu hình định dạng log bằng cách dùng các cờ như:

  - `log.Ldate`: ngày
  - `log.Ltime`: giờ
  - `log.Lshortfile`: tên file + dòng

```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
```

### SetOutput(io.Writer)

- Ghi log ra output khác, như file thay vì stdout.

```go
file, _ := os.Create("app.log")
log.SetOutput(file)
```

---

## 2. Ghi log

### Print / Println / Printf

```go
log.Print("Đây là log thường")
log.Println("Log có xuống dòng")
log.Printf("Số lượng: %d", 123)
```

### Fatal / Fatalf

- Ghi log và kết thúc chương trình (exit code 1).

```go
log.Fatal("Lỗi nghiêm trọng!")
log.Fatalf("Lỗi %s tại dòng %d", "ghi file", 42)
```

### Panic / Panicf

- Ghi log và panic (văng lỗi, in stack trace).

```go
log.Panic("Lỗi không thể tiếp tục!")
log.Panicf("Lỗi ở %s", "module A")
```

### Output(depth, msg)

- Ghi log thủ công với `depth` là độ sâu stack trace.

```go
log.Output(2, "Thông báo ghi log thủ công")
```

---

## 3. Flags thường dùng:

| Flag             | Ý nghĩa                    |                    |
| ---------------- | -------------------------- | ------------------ |
| `log.Ldate`      | Ghi ngày                   |                    |
| `log.Ltime`      | Ghi giờ                    |                    |
| `log.Lshortfile` | Ghi tên file + dòng (ngắn) |                    |
| `log.Llongfile`  | Ghi full path + dòng (dài) |                    |
| `log.LstdFlags`  | Kết hợp \`Ldate            | Ltime\` (mặc định) |

---

## 4. Thực hành cơ bản

```go
log.SetPrefix("[DEBUG] ")
log.SetFlags(log.LstdFlags | log.Lshortfile)
log.Println("Chạy ứng dụng...")
```

---

## 5. Ghi log nâng cao: Ghi vào file JSON

```go
package main

import (
    "encoding/json"
    "log"
    "os"
    "time"
)

type LogEntry struct {
    Level   string    `json:"level"`
    Message string    `json:"message"`
    Time    time.Time `json:"time"`
}

func main() {
    file, _ := os.Create("log.json")
    defer file.Close()

    logger := log.New(file, "", 0) // Không dùng prefix/flag

    entry := LogEntry{
        Level:   "INFO",
        Message: "Ứng dụng đã khởi chạy",
        Time:    time.Now(),
    }
    jsonBytes, _ := json.Marshal(entry)
    logger.Println(string(jsonBytes))
}
```

---

## Tài liệu tham khảo

- [https://pkg.go.dev/log](https://pkg.go.dev/log)

> ✨ Ghi log hiệu quả giúp bảo trì, debug và giám sát hệ thống dễ dàng hơn!
