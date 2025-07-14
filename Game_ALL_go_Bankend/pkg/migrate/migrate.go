package migrate

import (
	"log"

	"gorm.io/gorm"
)

// AutoMigrate runs database migrations for all registered models
func AutoMigrate(db *gorm.DB) error {
	if len(ModelRegistry) == 0 {
		log.Println("⚠️  Không có model nào được đăng ký để migrate")
		return nil
	}

	log.Println("🔄 Đang chạy migration cho các model đã đăng ký...")
	
	// Tự động tạo/cập nhật bảng cho tất cả các model đã đăng ký
	if err := db.AutoMigrate(ModelRegistry...); err != nil {
		log.Printf("❌ Lỗi khi chạy migration: %v", err)
		return err
	}

	log.Printf("✅ Đã tạo/cập nhật %d bảng thành công!", len(ModelRegistry))
	return nil
}