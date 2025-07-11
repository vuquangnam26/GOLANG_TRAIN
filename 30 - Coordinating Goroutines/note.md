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

Dự kiến: 3 \* 5000 = 15000 ❤️  Thực tế: sai số, mỗi lần mỗi khác.

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

### Quy tắc hoạt động RWMutex:

* Nếu RWMutex đang **mở khoá**, thì cả `RLock()` và `Lock()` đều có thể giành được khoá.
* Nếu đã có **ít nhất một reader** (đã `RLock()`), thì các reader khác vẫn có thể `RLock()` tiếp mà **không bị block**.
* Nếu có bất kỳ reader nào đang giữ khoá, thì writer (`Lock()`) sẽ bị **block cho đến khi tất cả các reader `RUnlock()`**.
* Nếu RWMutex đang bị `Lock()` bởi writer, thì **mọi lời gọi `RLock()` và `Lock()` khác đều bị block** cho đến khi writer `Unlock()`.
* Nếu một writer đang chờ `Lock()` trong khi có các reader giữ khoá, thì **reader mới cũng sẽ bị block** để tránh writer bị chặn mãi mãi (starvation).

### Khi nào dùng RWMutex?

* Khi có nhiều thao tác đọc và ít thao tác ghi.
* Cho phép tăng hiệu năng do nhiều goroutine được phép đọc đồng thời.

### 🎯 So sánh thực tế:

| Loại khóa   | Ví dụ thực tế                           | Dùng khi nào                |
| ----------- | --------------------------------------- | --------------------------- |
| `Lock()`    | Người chỉnh sửa tài liệu trong thư viện | Khi cần ghi dữ liệu         |
| `RLock()`   | Nhiều người cùng đọc tài liệu           | Khi chỉ đọc dữ liệu         |
| `RLocker()` | Phiên bản `Locker` để dùng trong `Cond` | Khi kết hợp với `sync.Cond` |

---

## 🧩 sync.Cond – Điều phối goroutine theo điều kiện

`sync.Cond` giúp **đồng bộ goroutine dựa trên điều kiện** thay vì chỉ khóa thông thường.

### Cách tạo:

```go
cond := sync.NewCond(rwmutex.RLocker())
```

Bạn có thể dùng `sync.Mutex{}` hoặc `sync.RWMutex{}` tuỳ theo mục đích.

### Cấu trúc Cond:

| Thành phần    | Mô tả                                                                      |
| ------------- | -------------------------------------------------------------------------- |
| `L`           | Locker được truyền khi tạo `Cond`. Dùng `Lock/Unlock` để bảo vệ điều kiện. |
| `Wait()`      | Nhả khóa và dừng goroutine cho đến khi có `Signal`/`Broadcast`.            |
| `Signal()`    | Đánh thức 1 goroutine đang chờ.                                            |
| `Broadcast()` | Đánh thức tất cả goroutine đang chờ.                                       |

### Locker là gì?

* Là interface có 2 phương thức:

```go
type Locker interface {
    Lock()
    Unlock()
}
```

* `sync.Mutex` và `sync.RWMutex` đều implement interface này.
* Nếu dùng `RWMutex.RLocker()` thì `Cond` sẽ làm việc với **read lock**.

### Vì sao dùng `for` với `Wait()`?

```go
for len(squares) == 0 {
    cond.Wait()
}
```

* Vì có thể xảy ra **spurious wakeup** (goroutine bị đánh thức mà không có lý do).
* Dùng `for` giúp **kiểm tra lại điều kiện**, tránh lỗi truy cập dữ liệu chưa sẵn sàng.

### Ví dụ consumer:

```go
func readSquares(id, max, iterations int) {
    cond.L.Lock() // tương đương rwmutex.RLock()
    for len(squares) == 0 {
        cond.Wait()
    }
    for i := 0; i < iterations; i++ {
        key := rand.Intn(max)
        fmt.Printf("#%v Read value: %v = %v\n", id, key, squares[key])
        time.Sleep(time.Millisecond * 100)
    }
    cond.L.Unlock() // tương đương rwmutex.RUnlock()
    waitGroup.Done()
}
```

### Ví dụ producer:

```go
func generateSquares(max int) {
    rwmutex.Lock()
    fmt.Println("Đang sinh dữ liệu...")
    for val := 0; val < max; val++ {
        squares[val] = val * val
    }
    rwmutex.Unlock()
    fmt.Println("Phát tín hiệu đánh thức")
    cond.Broadcast()
    waitGroup.Done()
}
```

### ⚖️ So sánh cách dùng `Lock` trong `sync.Cond`

| Trường hợp          | Dùng gì?                           | Giải thích                       |
| ------------------- | ---------------------------------- | -------------------------------- |
| Đọc dữ liệu chờ sẵn | `readCond.L.Lock()`                | Lấy read-lock để kiểm tra và đợi |
| Ghi dữ liệu sinh ra | `rwmutex.Lock()`                   | Lấy write-lock để cập nhật map   |
| Không nên dùng      | `rwmutex.Lock()` trong hàm chỉ đọc | Gây block không cần thiết        |

---

## 🤝 sync.WaitGroup

Dùng để chờ các goroutine hoàn thành trước khi main() thoát.

### Lưu ý:

* `Add(n)` trước khi start goroutine
* `Done()` đúng số lần
* `Wait()` để chờ hoàn thành

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

## ⏱ context – Hủy và giới hạn thời gian xử lý goroutine

### Các method của `Context`:

| Phương thức  | Mô tả                                                                                             |
| ------------ | ------------------------------------------------------------------------------------------------- |
| `Value(key)` | Trả về giá trị gắn với key đã truyền vào Context.                                                 |
| `Done()`     | Trả về một channel dùng để nhận tín hiệu huỷ.                                                     |
| `Deadline()` | Trả về `time.Time` và `bool` nếu có deadline.                                                     |
| `Err()`      | Trả về lỗi tương ứng khi channel `Done` đóng: `context.Canceled` hoặc `context.DeadlineExceeded`. |

### Các hàm tạo Context:

| Hàm                                  | Mô tả                                          |
| ------------------------------------ | ---------------------------------------------- |
| `context.Background()`               | Context gốc mặc định.                          |
| `context.WithCancel(ctx)`            | Tạo context mới có thể hủy với hàm `cancel()`. |
| `context.WithDeadline(ctx, time)`    | Context với deadline cụ thể.                   |
| `context.WithTimeout(ctx, duration)` | Context tự hủy sau thời gian.                  |
| `context.WithValue(ctx, key, val)`   | Gắn thêm dữ liệu vào context.                  |

### ✅ Khi nào nên dùng `context.WithCancel()`

| Tình huống                                 | Mô tả                                                                       |
| ------------------------------------------ | --------------------------------------------------------------------------- |
| ❌ Client huỷ yêu cầu                       | HTTP/gRPC: user đóng trình duyệt, ngắt request -> backend biết để huỷ xử lý |
| ✅ Dừng tất cả worker khi có lỗi            | Một goroutine lỗi → cancel() → dừng toàn bộ các goroutine khác              |
| ✅ Huỷ chủ động theo logic                  | Bạn tự quyết định khi nào nên dừng tác vụ (logic riêng)                     |
| ✅ Gắn cùng một context cho nhiều goroutine | Gọi cancel() một lần → dừng tất cả goroutine dùng context đó                |
| ✅ Giải phóng tài nguyên đúng lúc           | Tránh truy cập DB, log, IO khi request đã bị huỷ                            |

### Ví dụ sử dụng `context.WithCancel()`

```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    wg := sync.WaitGroup{}
    wg.Add(1)

    go func() {
        defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                fmt.Println("⛔ Dừng xử lý: context huỷ")
                return
            default:
                fmt.Println("✅ Đang xử lý...")
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()

    time.Sleep(2 * time.Second)
    fmt.Println("🚨 Huỷ context")
    cancel()

    wg.Wait()
}
```

---

## 🔍 Tổng kết

| Kỹ thuật       | Giải quyết gì?                        |
| -------------- | ------------------------------------- |
| `sync.Mutex`   | Truy cập biến chung đồng bộ           |
| `sync.RWMutex` | Cho phép đọc đồng thời, ghi độc quyền |
| `sync.Cond`    | Chờ và đánh thức theo điều kiện       |
| `WaitGroup`    | Chờ goroutine hoàn thành              |
| `atomic`       | Tăng giá trị nguyên tử                |
| `context`      | Huỷ goroutine hoặc giới hạn timeout   |

---

> ✨ Luôn truyền `*sync.WaitGroup` thay vì `sync.WaitGroup` để các goroutine cùng thao tác trên một địa chỉ chung!
