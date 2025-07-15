package migrate

import (
	"Game_ALL_go_bankend/models"
	"log"
	"reflect"
)

var ModelRegistry []interface{}

func RegisterModel(model interface{}) {
	if model == nil { return }
	ModelRegistry = append(ModelRegistry, model)
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr { modelType = modelType.Elem() }
	log.Printf("✅ Đã đăng ký model: %s", modelType.Name())
}

func init() {
	RegisterModel(&models.User{})
	RegisterModel(&models.Product{})
	RegisterModel(&models.ProductDetail{})
	RegisterModel(&models.ProductImage{})
	RegisterModel(&models.Order{})
	RegisterModel(&models.Category{})
}
