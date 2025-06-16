package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var commonDirs = []string{
	"config", "controllers", "errors", "middleware", "routes", "services", "utils",
}

var fullApiDirs = []string{
	"models/contracts", "models/transactions", "repositories",
}

func usage() {
	fmt.Println("Usage: goGinInitializer <module_name> [--full-api] [--git] [--gui]")
	fmt.Println("For more information, run with --help or -h")
	os.Exit(1)
}

func help() {
	fmt.Println("Usage: goGinInitializer <module_name> [--full-api] [--git]")
	fmt.Println("Options:")
	fmt.Println("  <module_name>   Name of the Go module to initialize. (Required)")
	fmt.Println("  --full-api      Initialize a full API structure with additional directories.")
	fmt.Println("  --git           Initialize a Git repository.")
	fmt.Println("  -g, --gui       Launch GUI for project initialization.")
	fmt.Println("  -h, --help      Show this help message.")
	os.Exit(0)
}

func createDirs(dirs []string) {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
	}
}

func writeFile(path string, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fmt.Printf("Error writing file %s: %v\n", path, err)
		os.Exit(1)
	}
}

func generateProject(moduleName string, fullApi bool, initGit bool) error {
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("go is not installed or not in PATH")
	}

	if err := exec.Command("go", "mod", "init", moduleName).Run(); err != nil {
		return fmt.Errorf("failed to initialize Go module: %v", err)
	}

	if err := exec.Command("go", "get", "github.com/gin-gonic/gin").Run(); err != nil {
		return fmt.Errorf("failed to get gin package: %v", err)
	}

	createDirs(commonDirs)
	if fullApi {
		createDirs(fullApiDirs)
	}

	if initGit {
		if err := exec.Command("git", "init").Run(); err != nil {
			return fmt.Errorf("failed to initialize git repository: %v", err)
		}
		fmt.Println("Initialized git repository.")
	}

	writeFile("main.go", fmt.Sprintf(`package main

import (
	"%s/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run()
}
`, moduleName))

	writeFile(filepath.Join("routes", "router.go"), strings.ReplaceAll(`
package routes

import (
	"{{MODULE}}/middleware"
	"{{MODULE}}/errors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.ErrorHandler())

	router.NoRoute(func(c *gin.Context) {
		err := errors.NewNotFoundError("The requested resource was not found")
		c.Error(err)
	})

	return router
}
`, "{{MODULE}}", moduleName))

	writeFile(filepath.Join("errors", "api_error.go"), strings.ReplaceAll(`
package errors

import "net/http"

type ApiError interface {
	error
	Status() int
	Message() string
	Type() string
}

type baseError struct {
	message    string
	statusCode int
	errType    string
}

func (e *baseError) Error() string   { return e.message }
func (e *baseError) Status() int     { return e.statusCode }
func (e *baseError) Message() string { return e.message }
func (e *baseError) Type() string    { return e.errType }

func defaultMessage(provided, fallback string) string {
	if provided != "" {
		return provided
	}
	return fallback
}

func NewBadRequestError(message string) ApiError {
	return &baseError{
		message:    defaultMessage(message, "Bad Request"),
		statusCode: http.StatusBadRequest,
		errType:    "bad_request",
	}
}

func NewNotFoundError(message string) ApiError {
	return &baseError{
		message:    defaultMessage(message, "Not Found"),
		statusCode: http.StatusNotFound,
		errType:    "not_found",
	}
}

func NewUnauthorizedError(message string) ApiError {
	return &baseError{
		message:    defaultMessage(message, "Unauthorized"),
		statusCode: http.StatusUnauthorized,
		errType:    "unauthorized",
	}
}

func NewInternalServerError(message string) ApiError {
	return &baseError{
		message:    defaultMessage(message, "Internal Server Error"),
		statusCode: http.StatusInternalServerError,
		errType:    "internal_server_error",
	}
}
`, "{{MODULE}}", moduleName))

	writeFile(filepath.Join("middleware", "error_handler.go"), strings.ReplaceAll(`
package middleware

import (
	"{{MODULE}}/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors

		if len(errs) == 0 {
			return
		}

		lastErr := errs.Last().Err

		if apiErr, ok := lastErr.(errors.ApiError); ok {
			c.JSON(apiErr.Status(), gin.H{
				"error": gin.H{
					"type":    apiErr.Type(),
					"message": apiErr.Message(),
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"type":    "INTERNAL_ERROR",
				"message": lastErr.Error(),
			},
		})
	}
}
`, "{{MODULE}}", moduleName))

	fmt.Println("✅ Project structure initialized.")
	return nil
}

func runCLI(args []string) {
	if len(os.Args) < 1 {
		usage()
	}

	moduleName := ""
	fullApi := false
	initGit := false

	for _, arg := range args {
		switch arg {
		case "--full-api":
			fullApi = true
		case "--git":
			initGit = true
		case "--help", "-h":
			help()
		default:
			if moduleName == "" {
				moduleName = arg
			} else {
				fmt.Printf("Unknown argument: %s\n", arg)
				usage()
			}
		}
	}

	if moduleName == "" {
		fmt.Println("Module name is required.")
		usage()
	}

	if err := generateProject(moduleName, fullApi, initGit); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
