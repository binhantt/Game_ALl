# Game ALL - Go Backend

## 🚀 Hướng dẫn sử dụng Auto Migration

### 1. Tạo model mới

Tạo file mới trong thư mục `models/`, ví dụ `product.go`:

```go
package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name     string  `gorm:"size:100;not null" json:"name"`
    Price    float64 `gorm:"not null" json:"price"`
    Quantity int     `gorm:"default:0" json:"quantity"`
}
```

### 2. Đăng ký model
Mở file `pkg/migrate/auto_register.go` và thêm dòng đăng ký:

```go
func init() {
    RegisterModel(&models.User{})     // Model có sẵn
    RegisterModel(&models.Product{})  // Thêm dòng này
}
```

### 3. Chạy migration

**Cách 1: Dùng script (khuyến nghị)**
```bash
go run scripts/migrate.go
```

**Cách 2: Chạy server (tự động chạy migration)**
```bash
go run .
```

### 4. Kiểm tra kết quả

Sau khi chạy xong, kiểm tra trong MySQL:

```sql
USE game_all;
SHOW TABLES;
DESCRIBE products;  -- Xem cấu trúc bảng
```

## 📌 Lưu ý quan trọng

- Tên bảng sẽ được chuyển đổi tự động thành dạng số nhiều (vd: `Product` → `products`)
- Để tùy chỉnh tên bảng, thêm method `TableName()` vào model:
  ```go
  func (Product) TableName() string {
      return "san_pham"  // Tên bảng tùy chỉnh
  }
  ```
- Luôn backup dữ liệu trước khi chạy migration trong môi trường production

## 🔧 Các lệnh thường dùng

- **Khởi động server:** `go run .`
- **Chạy migration:** `go run scripts/migrate.go`
- **Cài đặt thư viện:** `go get <package>`
- **Build:** `go build -o app .`

## 📞 Hỗ trợ
Nếu cần hỗ trợ, vui lòng liên hệ quản trị viên hệ thống.
