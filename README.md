# Game_ALL

## Mục lục

- [Giới thiệu](#giới-thiệu)
- [Cấu trúc dự án](#cấu-trúc-dự-án)
- [Yêu cầu hệ thống](#yêu-cầu-hệ-thống)
- [Hướng dẫn cài đặt](#hướng-dẫn-cài-đặt)
  - [1. Cài đặt Backend (Go)](#1-cài-đặt-backend-go)
  - [2. Cài đặt Frontend Admin (React)](#2-cài-đặt-frontend-admin-react)
- [Chạy dự án](#chạy-dự-án)
- [Thông tin liên hệ](#thông-tin-liên-hệ)

---

## Giới thiệu

Game_ALL là hệ thống quản lý game gồm backend viết bằng Golang và frontend quản trị bằng React.

## Cấu trúc dự án

```
Game_ALL/
  ├── Game_ALL_go_Bankend/         # Backend Golang
  ├── web_react_Game_ALL_Admin/    # Frontend React Admin
  ├── web_Astro_user/              # (Tùy chọn) Frontend cho user
```

## Yêu cầu hệ thống

- Go >= 1.18
- Node.js >= 16.x
- pnpm hoặc npm hoặc yarn
- Git

## Hướng dẫn cài đặt

### 1. Cài đặt Backend (Go)

```bash
cd Game_ALL_go_Bankend
go mod tidy
```

- Cấu hình database trong file `config/database.go` nếu cần.
- Chạy migrate (nếu có):

```bash
go run scripts/migrate.go
```

- Chạy server:

```bash
go run main.go
```

### 2. Cài đặt Frontend Admin (React)

```bash
cd web_react_Game_ALL_Admin
pnpm install   # hoặc npm install hoặc yarn install
```

- Chạy ứng dụng:

```bash
pnpm dev       # hoặc npm run dev hoặc yarn dev
```

### (Tùy chọn) Cài đặt Frontend User (Astro)

Nếu có thư mục `web_Astro_user`, thực hiện tương tự như frontend admin.

## Chạy dự án

1. Khởi động backend trước.
2. Khởi động frontend admin.
3. Truy cập địa chỉ được in ra ở terminal (thường là http://localhost:5173 cho frontend, http://localhost:8080 cho backend).

## Thông tin liên hệ

- Email: [your-email@example.com]
- Zalo: [số điện thoại]
- Facebook: [link facebook] 
