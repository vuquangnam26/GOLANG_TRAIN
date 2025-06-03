# 📘 Go Notes: Composition, Constructors, and Promoted Fields

## ✅ 1. Composition ( Thành phần)

* Go **không có kế thừa (inheritance)** như OOP truyền thống.
* Thay vào đó, Go sử dụng **composition** bằng cách **nhúng struct này vào struct khác**.
* Các trường (fields) từ struct được nhúng sẽ trở thành **promoted fields**.

```go
type Product struct {
    Name     string
    Category string
    price    float64
}

type Boat struct {
    Product           // embedded
    Capacity int
    Motorized bool
}
```

---

## 🔁 2. Promoted Fields

* Khi struct `Product` được nhúng trong `Boat`, bạn có thể truy cập các trường như `boat.Name` thay vì `boat.Product.Name` (chỉ sau khi đã tạo giá trị).
* **Lưu ý**: Cách truy cập `boat.Name` chỉ hoạt động **sau khi đã tạo giá trị**, không áp dụng cho **literal syntax**.
* Nếu có **trùng tên field hoặc method**, Go **không thể thực hiện promotion**, điều này gây ra lỗi hoặc kết quả không như mong đợi.

---

## 🚫 3. Literal Syntax & Lỗi thường gặp

```go
// ❌ Compiler error: unknown field 'Name'
boat := Boat{
    Name: "Kayak",     // lỗi
    Capacity: 1,
}
```

* Cách đúng:

```go
boat := Boat{
    Product: Product{
        Name: "Kayak",
        Category: "Watersports",
        price: 275,
    },
    Capacity: 1,
    Motorized: false,
}
```

---

## 💪 4. Constructors (Hàm tạo)

* Dùng constructor để đảm bảo các trường được khởi tạo đúng cách.
* Tránh việc quên gán giá trị (vd: `price = 0`).
* Tăng khả năng bảo trì và mở rộng.

```go
func NewProduct(name, category string, price float64) Product {
    return Product{Name: name, Category: category, price: price}
}

func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
    return &Boat{
        Product:   NewProduct(name, "Watersports", price),
        Capacity:  capacity,
        Motorized: motorized,
    }
}
```

---

## 🌟 5. SpecialDeal & Nested Composition

* Go cho phép **nhúng nhiều struct** trong một struct khác, ví dụ như `RentalBoat` chứa `*Boat` và `*Crew`:

```go
type Crew struct {
    Captain, FirstOfficer string
}

type RentalBoat struct {
    *Boat
    IncludeCrew bool
    *Crew
}
```

* Constructor kết hợp:

```go
func NewRentalBoat(name string, price float64, capacity int, motorized, crewed bool, captain, firstOfficer string) *RentalBoat {
    return &RentalBoat{
        Boat: NewBoat(name, price, capacity, motorized),
        IncludeCrew: crewed,
        Crew: &Crew{Captain: captain, FirstOfficer: firstOfficer},
    }
}
```

* Dùng promoted fields:

```go
rental := store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice")
fmt.Println(rental.Name)       // from Product (through Boat)
fmt.Println(rental.Captain)    // from Crew
```

* Nếu có **field hoặc method trùng tên**, Go sẽ không thể phân biệt và gây lỗi **ambiguity**.

---

## ✨ 6. Using Composition to Implement Interfaces

* Go sử dụng interface để mô tả tính đa hình.
* Khi một struct nhúng một struct khác, các method từ struct nhúng được **promote**, giúc struct cha tự động thỏ c hiện interface mà không cần duplicate method.

```go
type Describer interface {
    Describe() string
}

type SpecialDeal struct {
    *Product
    Offer string
}

func NewSpecialDeal(offer string, p *Product, discount float64) *SpecialDeal {
    p.price -= discount
    return &SpecialDeal{Product: p, Offer: offer}
}

func (sd *SpecialDeal) Describe() string {
    return fmt.Sprintf("Deal: %s (%s)", sd.Name, sd.Offer)
}
```

* `SpecialDeal` đang nhúng `*Product`, nên vẫn sử dụng được `Name`, `Category`, `Price()` như Product.
* Khi `SpecialDeal` được dùng với interface `Describer`, Go sẽ xem `Describe` và tất cả method promoted khi kiểm tra tính đồng nhất (conformance).

---

## 💬 7. Lời khuyên

* **Luôn dùng constructor nếu có**.
* Tránh literal syntax khi struct lồng nhau phức tạp.
* Đảm bảo các constructor gọi nhau khi cần khởi tạo nhiều lớp lòng.
* Tránh đặt tên trùng nhau trong các struct nhúng.
* Khi implement interface, đây là điểm mạnh của composition: không cần viết lại method.

---

## 🔹 8. Dùng Interface với Composition và Map

* Khi struct nhúng struct khác có method khớp interface, Go **tính cả promoted methods** khi xem xét interface.

```go
type ItemForSale interface {
    Price(taxRate float64) float64
}
```

* `Product` implement trực tiếp method `Price`:

```go
func (p *Product) Price(taxRate float64) float64 {
    return p.price * (1 + taxRate)
}
```

* `Boat` không có method `Price`, nhưng nhúng `Product`, nên thừa hưởng `Price()`.

```go
items := map[string]ItemForSale{
    "Kayak": store.NewProduct("Kayak", "Watersports", 279),
    "Yacht": store.NewBoat("Yacht", 50000, 5, true),
}

for name, item := range items {
    fmt.Println("Item:", name, "Price with Tax:", item.Price(0.2))
}
```

✅ **Output:**

```
Item: Kayak Price with Tax: 334.8
Item: Yacht Price with Tax: 60000
```

> ✉️ Dùng interface giúc gom nhiều kiểu khác nhau nhưng có chung hành vi (qua method) trong một map hoặc slice chung.

---

## 🌐 9. Lỗi khi Match Multiple Interface Types trong Type Switch

* Khi match nhiều kiểu trong `type switch`, Go **không thực hiện type assertion**.
* Biến trong switch (vd: `item`) sẽ mang kiểu giao diện (interface type), dẫn đến lỗi khi truy cập field cụ thể.

```go
switch item := val.(type) {
case *store.Product, *store.Boat:
    fmt.Println(item.Name) // ❌ Compiler error
}
```

* ❤️ **Giải pháp**: Tách switch theo kiểu riêng biệt:

```go
switch item := val.(type) {
case *store.Product:
    fmt.Println(item.Name)
case *store.Boat:
    fmt.Println(item.Name)
}
```

* Hoặc: Dùng `type assertion` sau khi match.

```go
if p, ok := val.(*store.Product); ok {
    fmt.Println(p.Name)
}
```
