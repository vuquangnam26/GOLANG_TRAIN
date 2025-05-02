# ⚙️ File Cấu Hình Linter Cho Go (TOML)

Đây là file cấu hình để thiết lập các quy tắc kiểm tra chất lượng mã nguồn (linting) cho dự án Golang. File thường được dùng bởi các công cụ như `golangci-lint`, giúp phát hiện lỗi và giữ cho mã nguồn sạch sẽ, nhất quán.

---

## 🔧 Thiết Lập Chung

| Thuộc tính               | Giá trị     | Giải thích |
|--------------------------|-------------|------------|
| `ignoreGeneratedHeader`  | `false`     | Không bỏ qua các file được tạo tự động |
| `severity`               | `"warning"` | Mức độ cảnh báo mặc định là cảnh báo (warning) |
| `confidence`             | `0.8`       | Độ tin cậy tối thiểu để hiển thị cảnh báo |
| `errorCode`              | `0`         | Mã lỗi mặc định (không định nghĩa cụ thể) |
| `warningCode`            | `0`         | Mã cảnh báo mặc định (không định nghĩa cụ thể) |

---

## ✅ Các Quy Tắc Kiểm Tra Được Bật

| Quy Tắc                    | Mô Tả |
|----------------------------|-------|
| `blank-imports`            | Cảnh báo khi dùng `import _` mà không giải thích rõ ràng |
| `context-as-argument`      | Bắt buộc `context.Context` là đối số đầu tiên trong các hàm |
| `context-keys-type`        | Đảm bảo key trong context có kiểu riêng biệt để tránh lỗi |
| `dot-imports`              | Tránh dùng `import .` vì sẽ gây nhầm lẫn trong code |
| `error-return`             | Kiểm tra hàm có xử lý giá trị trả về là lỗi hay không |
| `error-strings`            | Lỗi nên được viết thường, không bắt đầu bằng chữ hoa hoặc chứa định dạng không cần thiết |
| `error-naming`             | Tên biến lỗi nên đặt là `err` để dễ hiểu và theo chuẩn |
| `if-return`                | Đơn giản hóa cấu trúc điều kiện `if` kết hợp với `return` khi có thể |
| `increment-decrement`      | Tránh viết biểu thức tăng/giảm biến một cách khó hiểu |
| `var-naming`               | Tên biến phải rõ ràng, dễ hiểu, tuân thủ quy tắc đặt tên |
| `var-declaration`          | Khuyến khích dùng `:=` khi khai báo và khởi tạo biến cùng lúc |
| `package-comments`         | Package cần có chú thích đầu mô tả chức năng rõ ràng |
| `range`                    | Phát hiện lỗi thường gặp khi sử dụng `for range` |
| `receiver-naming`          | Tên biến nhận trong method nên ngắn gọn, thường dùng chữ cái đầu |
| `time-naming`              | Biến liên quan đến thời gian nên có hậu tố như `Time`, `Duration` |
| `unexported-return`        | Không nên trả về kiểu không export từ các hàm export |
| `indent-error-flow`        | Phần xử lý lỗi nên được thụt lề hợp lý để dễ đọc |
| `errorf`                   | Khuyến khích dùng `fmt.Errorf()` với định dạng thông báo lỗi |

> 💡 Lưu ý: Rule `[rule.exported]` hiện đang bị **tắt** (bằng cách comment) và sẽ không kiểm tra tên hàm/struct export.

---

## 📂 Mục Đích

Giúp đảm bảo mã nguồn Go:
- Dễ đọc
- Dễ bảo trì
- Tuân thủ chuẩn cộng đồng Go

Bạn có thể thay đổi file `.toml` này tùy theo nhu cầu kiểm tra cụ thể trong dự án.

---
