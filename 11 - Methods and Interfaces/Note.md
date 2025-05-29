# Go Slice - Lưu ý quan trọng

## Tổng quan
Slice là một trong những data structure quan trọng nhất trong Go. Tuy nhiên, việc hiểu sai về cách slice hoạt động có thể dẫn đến nhiều bug khó phát hiện. File này tổng hợp các lưu ý quan trọng khi làm việc với slice.

## 1. Slice là Reference Type

⚠️ **Lưu ý**: Slice không phải là value type mà là reference type. Khi pass slice vào function, bạn đang pass reference đến underlying array.

```go
func modifySlice(s []int) {
    s[0] = 999  // Thay đổi này sẽ ảnh hưởng đến slice gốc
}

func main() {
    nums := []int{1, 2, 3}
    modifySlice(nums)
    fmt.Println(nums) // Output: [999 2 3]
}
```

## 2. Length vs Capacity

Slice có hai thuộc tính quan trọng:
- **Length** (`len()`): số phần tử hiện tại trong slice
- **Capacity** (`cap()`): dung lượng tối đa của underlying array

```go
s := make([]int, 3, 5) // length=3, capacity=5
fmt.Println(len(s))    // 3
fmt.Println(cap(s))    // 5

// Có thể append thêm 2 phần tử nữa mà không cần reallocate
s = append(s, 4, 5)
fmt.Println(len(s))    // 5
fmt.Println(cap(s))    // 5
```

## 3. Append có thể tạo slice mới

⚠️ **Cảnh báo**: Khi capacity không đủ, `append()` sẽ tạo underlying array mới với capacity lớn hơn (thường gấp đôi).

```go
s1 := []int{1, 2, 3}
s2 := s1              // s2 cùng trỏ đến array với s1
s1 = append(s1, 4)    // Nếu capacity đủ: s1 và s2 vẫn share array
                      // Nếu capacity không đủ: s1 trỏ đến array mới

s2[0] = 999
fmt.Println(s1)       // Kết quả có thể khác nhau tùy vào capacity
fmt.Println(s2)
```

## 4. Zero Value và Nil Slice

```go
var s []int           // nil slice
s2 := []int{}         // empty slice (không phải nil)
s3 := make([]int, 0)  // empty slice (không phải nil)

fmt.Println(s == nil)   // true
fmt.Println(s2 == nil)  // false
fmt.Println(s3 == nil)  // false

// Cả 3 đều có thể append được
s = append(s, 1)
s2 = append(s2, 1)
s3 = append(s3, 1)
```

## 5. Slicing chia sẻ memory

⚠️ **Nguy hiểm**: Khi tạo sub-slice từ slice gốc, chúng sẽ chia sẻ cùng underlying array.

```go
original := []int{1, 2, 3, 4, 5}
sub := original[1:3]  // [2, 3]
sub[0] = 999
fmt.Println(original) // [1 999 3 4 5] - slice gốc bị thay đổi!
```

### Giải pháp: Sử dụng copy()

```go
original := []int{1, 2, 3, 4, 5}
sub := make([]int, 2)
copy(sub, original[1:3])  // Copy data thay vì share reference
sub[0] = 999
fmt.Println(original)     // [1 2 3 4 5] - không bị thay đổi
```

## 6. Memory Leak với large slice

⚠️ **Memory Leak**: Khi giữ reference đến một phần nhỏ của slice lớn.

```go
// BAD: Giữ reference đến toàn bộ large slice
func processData(data []byte) []byte {
    return data[10:20]  // Chỉ dùng 10 bytes nhưng giữ reference đến toàn bộ data
}

// GOOD: Copy data cần thiết
func processData(data []byte) []byte {
    result := make([]byte, 10)
    copy(result, data[10:20])
    return result  // data gốc có thể được garbage collected
}
```

## 7. Performance Best Practices

### Pre-allocate capacity
```go
// BAD: Capacity tăng dần, nhiều lần reallocate
var items []Item
for i := 0; i < 1000; i++ {
    items = append(items, Item{})
}

// GOOD: Pre-allocate capacity
items := make([]Item, 0, 1000)
for i := 0; i < 1000; i++ {
    items = append(items, Item{})
}
```

### Sử dụng copy() thay vì loop
```go
// BAD
for i, v := range source {
    dest[i] = v
}

// GOOD
copy(dest, source)
```

## 8. Common Mistakes

### Mistake 1: Modify slice trong loop
```go
// BAD: Có thể gây panic hoặc skip elements
for i, v := range slice {
    if condition {
        slice = append(slice[:i], slice[i+1:]...)  // Modify slice trong loop
    }
}

// GOOD: Loop ngược hoặc tạo slice mới
for i := len(slice) - 1; i >= 0; i-- {
    if condition {
        slice = append(slice[:i], slice[i+1:]...)
    }
}
```

### Mistake 2: Assume slice truyền vào function không thay đổi
```go
func process(data []int) {
    data[0] = 999  // Thay đổi slice gốc!
}

// Nếu muốn không thay đổi slice gốc
func process(data []int) {
    localData := make([]int, len(data))
    copy(localData, data)
    localData[0] = 999  // Safe
}
```

## 9. Testing Slice Code

```go
func TestSliceModification(t *testing.T) {
    original := []int{1, 2, 3}
    originalCopy := make([]int, len(original))
    copy(originalCopy, original)
    
    process(original)
    
    // Verify original slice wasn't modified unexpectedly
    if !reflect.DeepEqual(original, originalCopy) {
        t.Error("Original slice was modified")
    }
}
```

## 10. Debugging Tips

```go
func debugSlice(name string, s []int) {
    fmt.Printf("%s: len=%d, cap=%d, data=%v, ptr=%p\n", 
        name, len(s), cap(s), s, &s[0])
}

func main() {
    s1 := make([]int, 3, 5)
    debugSlice("s1", s1)
    
    s2 := append(s1, 4, 5)
    debugSlice("s2", s2)
    
    s3 := append(s2, 6)  // This will reallocate
    debugSlice("s3", s3)
}
## Các cách so sánh slice
1. Chỉ có thể so sánh với nil
govar s []int
fmt.Println(s == nil) // ✅ OK - true

s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
// fmt.Println(s1 == s2) // ❌ Compile error!
2. Sử dụng reflect.DeepEqual()
goimport "reflect"

s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
s3 := []int{1, 2, 4}

fmt.Println(reflect.DeepEqual(s1, s2)) // true
fmt.Println(reflect.DeepEqual(s1, s3)) // false
3. Viết function so sánh custom
gofunc equalSlices(a, b []int) bool {
if len(a) != len(b) {
return false
}
for i, v := range a {
if v != b[i] {
return false
}
}
return true
}

s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
fmt.Println(equalSlices(s1, s2)) // true
4. Sử dụng slices.Equal() (Go 1.21+)
goimport "slices"

s1 := []int{1, 2, 3}
s2 := []int{1, 2, 3}
fmt.Println(slices.Equal(s1, s2)) // true
Tại sao không thể so sánh slice?
Slice là reference type và có thể thay đổi được, nên Go không cho phép so sánh trực tiếp để tránh nhầm lẫn. Ví dụ:
gos1 := []int{1, 2, 3}
s2 := s1  // Cùng reference
s3 := []int{1, 2, 3}  // Khác reference nhưng cùng nội dung

// Nếu cho phép so sánh ==:
// s1 == s2 sẽ là true (cùng reference)
// s1 == s3 sẽ là gì? true hay false?
```

## Kết luận

Slice trong Go rất mạnh mẽ nhưng cần hiểu rõ cách hoạt động để tránh bug. Những điểm quan trọng nhất:

1. Slice là reference type
2. Hiểu rõ sự khác biệt giữa length và capacity
3. Cẩn thận với việc chia sẻ underlying array
4. Sử dụng `copy()` khi muốn tránh chia sẻ memory
5. Pre-allocate capacity khi có thể để tối ưu performance
6. Do là kiểu reference nên không thể so sánh với nhau , chỉ dc so sánh với nil
---

**Tác giả**: Vu Quang Nam 
**Version**: 1.0