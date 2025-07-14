# Game_ALL_go_Bankend

## Giới thiệu
Đây là backend viết bằng Golang cho dự án Game_ALL. Thư mục này chứa mã nguồn, cấu hình, models, routes, services, migrations, v.v.

## Hướng dẫn sử dụng script generator.go
Script `scripts/generator.go` giúp bạn tạo nhanh file model cho một bảng mới.

### Cách sử dụng:
1. Mở terminal tại thư mục `Game_ALL_go_Bankend`.
2. Chạy lệnh sau để tạo model mới:

```bash
go run scripts/generator.go
```

3. Nhập tên bảng (ví dụ: `user_profile`).
4. File model sẽ được tạo trong thư mục `models/` với tên tương ứng (ví dụ: `models/user_profile.go`).

## Một số lệnh Go thường dùng

- **Chạy ứng dụng:**
  ```bash
  go run main.go
  ```
- **Cài đặt dependencies:**
  ```bash
  go mod tidy
  ```
- **Chạy migration (nếu có script):**
  ```bash
  go run migrations/ten_file_migration.go
  ```
- **Build project:**
  ```bash
  go build -o app main.go
  ```
- **Kiểm tra lỗi:**
  ```bash
  go vet ./...
  go fmt ./...
  ```

## Thư mục chính
- `models/`: Định nghĩa các struct model
- `controllers/`: Xử lý logic request
- `routes/`: Định nghĩa các route
- `services/`: Các service logic
- `config/`: Cấu hình (database, env, ...)
- `migrations/`: Script migration database
- `scripts/`: Script tiện ích (generator, ...)

---
Nếu có vấn đề hoặc cần thêm hướng dẫn, hãy liên hệ team backend! 