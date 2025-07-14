package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ControllerData struct {
	ModulePath    string
	ModelName     string
	TableName     string
	ControllerName string
}

func main() {
	// Parse command line flags
	modelName := flag.String("model", "", "Name of the model (required)")
	tableName := flag.String("table", "", "Name of the database table (optional, will use lowercase plural of model name if not provided)")
	modulePath := flag.String("module", "Game_ALL_go_bankend", "Go module path")

	flag.Parse()

	if *modelName == "" {
		log.Fatal("Error: Model name is required. Use -model flag to specify the model name.")
	}

	if *tableName == "" {
		*tableName = strings.ToLower(*modelName) + "s"
	}

	data := ControllerData{
		ModulePath:    *modulePath,
		ModelName:     *modelName,
		TableName:     *tableName,
		ControllerName: *modelName + "Controller",
	}

	// Create controllers directory if it doesn't exist
	if err := os.MkdirAll("../controllers", 0755); err != nil {
		log.Fatalf("Error creating controllers directory: %v", err)
	}

	// Create the controller file
	fileName := fmt.Sprintf("%s_controller.go", strings.ToLower(*modelName))
	filePath := filepath.Join("..", "controllers", fileName)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating controller file: %v", err)
	}
	defer file.Close()

	// Parse and execute the template
	tmpl := template.Must(template.New("controller").Parse(controllerTemplate))
	if err := tmpl.Execute(file, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Update auto_register.go to include the new controller
	updateAutoRegister(*modelName)

	fmt.Printf("✅ Successfully created controller: %s\n", filePath)
	fmt.Println("Don't forget to register your routes in routes/index.go")
}

func updateAutoRegister(modelName string) {
	// This function would update the auto_register.go file
	// Implementation depends on your project structure
	// For now, we'll just print a reminder
	fmt.Printf("ℹ️  Please add the following to pkg/migrate/auto_register.go:\n")
	fmt.Printf("\tRegisterModel(&models.%s{})\n\n", modelName)
}

const controllerTemplate = `package controllers

import (
	"net/http"
	"{{.ModulePath}}/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// {{.ControllerName}} handles HTTP requests for {{.ModelName}}
type {{.ControllerName}} struct {
	DB *gorm.DB
}

// New{{.ControllerName}} creates a new instance of {{.ControllerName}}
func New{{.ControllerName}}(db *gorm.DB) *{{.ControllerName}} {
	return &{{.ControllerName}}{DB: db}
}

// GetAll returns all {{.TableName}}
func (c *{{.ControllerName}}) GetAll(ctx *gin.Context) {
	var {{.TableName}} []models.{{.ModelName}}
	if err := c.DB.Find(&{{.TableName}}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": {{.TableName}}})
}

// GetByID returns a single {{.ModelName}} by ID
func (c *{{.ControllerName}}) GetByID(ctx *gin.Context) {
	var {{.ModelName | toLower}} models.{{.ModelName}}
	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&{{.ModelName | toLower}}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": {{.ModelName | toLower}}})
}

// Create creates a new {{.ModelName}}
func (c *{{.ControllerName}}) Create(ctx *gin.Context) {
	// Validate input
	var input models.{{.ModelName}}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create {{.ModelName}}
	if err := c.DB.Create(&input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating record"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": input})
}

// Update updates a {{.ModelName}}
func (c *{{.ControllerName}}) Update(ctx *gin.Context) {
	// Get model if exist
	var {{.ModelName | toLower}} models.{{.ModelName}}
	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&{{.ModelName | toLower}}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.{{.ModelName}}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&{{.ModelName | toLower}}).Updates(input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating record"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": {{.ModelName | toLower}}})
}

// Delete deletes a {{.ModelName}}
func (c *{{.ControllerName}}) Delete(ctx *gin.Context) {
	// Get model if exist
	var {{.ModelName | toLower}} models.{{.ModelName}}
	if err := c.DB.Where("id = ?", ctx.Param("id")).First(&{{.ModelName | toLower}}).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.DB.Delete(&{{.ModelName | toLower}}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting record"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
`
