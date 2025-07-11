# 📦 Các Hàm HỖ Trợ Kiểm Thử Trong Go

## 🥪 Nhóm các hàm `t.*` phổ biến trong testing

Trong gói `testing`, struct `*testing.T` cung cấp nhiều phương thức để log, đánh dấu lỗi và dừng test.

| Hàm                         | Chức năng                                                               |
| --------------------------- | ----------------------------------------------------------------------- |
| `t.Log(...vals)`            | Ghi log thông tin (không đánh dấu lỗi)                                  |
| `t.Logf(template, ...vals)` | Ghi log có định dạng                                                    |
| `t.Error(...errs)`          | Ghi lỗi và đánh dấu test thất bại, **vẫn tiếp tục thực thi**            |
| `t.Errorf(template, ...)`   | Ghi lỗi có định dạng, đánh dấu test thất bại, **vẫn tiếp tục thực thi** |
| `t.Fail()`                  | Đánh dấu test thất bại, **không dừng test**                             |
| `t.FailNow()`               | Đánh dấu test thất bại và **dừng test ngay lập tức**                    |
| `t.Fatal(...vals)`          | Ghi lỗi + `FailNow()`                                                   |
| `t.Fatalf(template, ...)`   | Ghi lỗi có định dạng + `FailNow()`                                      |
| `t.Failed()`                | Trả về `true` nếu test đã fail                                          |

## 🔍 Ví dụ minh họa các hàm

```go
func TestExample(t *testing.T) {
    t.Log("Bắt đầu test...") // t.Log

    t.Logf("Giá trị ban đầu: %d", 10) // t.Logf

    if 2+2 != 4 {
        t.Error("Phép tính sai") // t.Error
    }

    if 2*2 != 5 {
        t.Errorf("Kỳ vọng 5, nhưng nhận %d", 2*2) // t.Errorf
    }

    if false {
        t.Fail() // Đánh dấu fail nhưng không log
    }

    if true {
        t.Log("Dừng test ngay lập tức")
        t.FailNow() // Dừng test ngay
        t.Log("Không bao giờ đến được đây")
    }
}
```

## 💡 Gợi ý sử dụng

- Dùng `t.Error`/`t.Errorf` nếu muốn **ghi nhận lỗi nhưng không dừng test**.
- Dùng `t.Fatal`/`t.Fatalf` nếu **lỗi nghiêm trọng cần dừng ngay lập tức** (ví dụ: không mở được file test).

---

Bạn có thể dùng thêm `go test -v` để xem đầy đủ log các test đang chạy.

> 📚 Gói `testing` là cơ bản nhưng rất mạnh mẽ. Go hỗ trợ benchmark, test song song và mock đơn giản chỉ với `testing.T` và các thư viện mở rộng như `testify`.
