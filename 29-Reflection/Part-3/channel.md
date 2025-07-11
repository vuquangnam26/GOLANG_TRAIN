# ✨ Đồng bộ trong Go: Race Condition, sync.Mutex, WaitGroup, Atomic

## ⚠️ Vấn đề: Race Condition

Khi nhiều goroutine truy cập và ghi cùng một biến, các thao tác như `counter++` có thể bị gây lỗi do **không phải là thao tác nguyên tử**. Dẫn đến kết quả sai hoặc khác nhau mỗi lần chạy.

**Ví dụ:**

```go
var counter int

for i := 0; i < 3; i++ {
    go func() {
        for j := 0; j < 5000; j++ {
            counter++ // Sai: không đồng bộ
        }
    }()
}
```

Dự kiến: 3 \* 5000 = 15000 ❤️ Thực tế: sai số, mỗi lần mỗi khác.

---

## 🔐 sync.Mutex

### ✅ Giải pháp: Dùng mutex để khoá truy cập

### Các phương thức:

| Phương thức | Mô tả                                                 |
| ----------- | ----------------------------------------------------- |
| `Lock()`    | Khoá mutex. Nếu đang bị khoá thì block goroutine chờ. |
| `Unlock()`  | Mở khoá. Bắt buộc phải gọi sau `Lock`.                |

### Ví dụ:

```go
var counter int
var mutex sync.Mutex

func doSum(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < n; i++ {
        mutex.Lock()
        counter++
        mutex.Unlock()
    }
}
```

---

## 🧠 sync.RWMutex

RWMutex cho phép nhiều goroutine **đọc song song** nhưng chỉ cho phép **một goroutine ghi tại một thời điểm**.

### Các phương thức:

| Phương thức | Mô tả                                                   |
| ----------- | ------------------------------------------------------- |
| `RLock()`   | Lấy read-lock. Block nếu đang có write-lock.            |
| `RUnlock()` | Giải phóng read-lock.                                   |
| `Lock()`    | Lấy write-lock. Block nếu đang có read/write-lock khác. |
| `Unlock()`  | Giải phóng write-lock.                                  |
| `RLocker()` | Trả về `Locker` interface dành cho read-lock.           |

### Khi nào dùng RWMutex?

- Khi có nhiều thao tác đọc và ít thao tác ghi.
- Cho phép tăng hiệu năng do nhiều goroutine được phép đọc đồng thời.

---

## 🤝 sync.WaitGroup

Dùng để chờ các goroutine hoàn thành trước khi main() thoát.

### Lưu ý:

- `Add(n)` trước khi start goroutine
- `Done()` đúng số lần
- `Wait()` để chờ hoàn thành

### Lỗi thường gặp:

| Lỗi                      | Nguyên nhân                                             |
| ------------------------ | ------------------------------------------------------- |
| `Done()` > `Add()`       | Panic: counter âm                                       |
| `Done()` < `Add()`       | Wait() chờ mãi                                          |
| Truyền giá trị WaitGroup | Mỗi goroutine dùng bản sao khác nhau, Wait() chờ vô tín |

### Ví dụ panic do `Done()` nhiều hơn `Add()`:

```go
var wg sync.WaitGroup
wg.Add(1)
wg.Done()
wg.Done() // ❌ panic: sync: negative WaitGroup counter
```

---

## ⚡ sync/atomic

Dùng cho các thao tác nguyên tử như tăng biến.

```go
import "sync/atomic"
var counter int64
atomic.AddInt64(&counter, 1)
```

---

## 🔍 Tổng kết

| Kỹ thuật       | Giải quyết gì?                        |
| -------------- | ------------------------------------- |
| `sync.Mutex`   | Truy cập biến chung đồng bộ           |
| `sync.RWMutex` | Cho phép đọc đồng thời, ghi độc quyền |
| `WaitGroup`    | Chờ goroutine hoàn thành              |
| `atomic`       | Tăng giá trị nguyên tử                |

---

> ✨ Luôn truyền `*sync.WaitGroup` thay vì `sync.WaitGroup` để các goroutine cùng thao tác trên một địa chỉ chung!
