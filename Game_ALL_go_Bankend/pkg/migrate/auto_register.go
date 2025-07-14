package migrate

import (
	"Game_ALL_go_bankend/models"
	"log"
	"reflect"
)

// ModelRegistry chứa tất cả các model cần đăng ký
var ModelRegistry []interface{}

// RegisterModel đăng ký một model vào danh sách tự động migrate
func RegisterModel(model interface{}) {
	if model == nil {
		return
	}
	ModelRegistry = append(ModelRegistry, model)

	// Lấy tên model để log
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	log.Printf("✅ Đã đăng ký model: %s", modelType.Name())
}

// AutoRegisterModels tự động đăng ký tất cả các model
type ModelInterface interface {
	// Định nghĩa interface rỗng để đảm bảo chỉ các model mới được đăng ký
}

// Hàm init sẽ tự động chạy khi package được import
func init() {
	// Đăng ký tất cả các model ở đây
	// Khi thêm model mới, chỉ cần import package models và đăng ký ở đây
	RegisterModel(&models.User{})
	RegisterModel(&models.Product{})
	RegisterModel(&models.ProductDetail{})
	RegisterModel(&models.Order{})
	RegisterModel(&models.ProductImage{})
	RegisterModel(&models.Category{})
	// Thêm các model khác ở đây
	// RegisterModel(&models.Product{})
}
