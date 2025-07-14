@echo off
echo Generating controller...
go run generate_controller.go %*

echo.
echo Remember to:
echo 1. Add your model to pkg/migrate/auto_register.go
echo 2. Add routes in routes/index.go
echo 3. Run 'go mod tidy' to update dependencies
