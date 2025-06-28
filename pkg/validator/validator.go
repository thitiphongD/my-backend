package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Register JSON tag name
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// ValidateStruct validates a struct and returns validation errors
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// ParseAndValidate parses the request body and validates it
func ParseAndValidate(c *fiber.Ctx, s interface{}) error {
	if err := c.BodyParser(s); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON format")
	}

	if err := ValidateStruct(s); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, formatValidationError(err))
	}

	return nil
}

// formatValidationError formats validation errors into a readable string
func formatValidationError(err error) string {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errors = append(errors, err.Field()+" is required")
		case "email":
			errors = append(errors, err.Field()+" must be a valid email")
		case "min":
			errors = append(errors, err.Field()+" must be at least "+err.Param()+" characters")
		case "max":
			errors = append(errors, err.Field()+" must be at most "+err.Param()+" characters")
		default:
			errors = append(errors, err.Field()+" is invalid")
		}
	}

	return strings.Join(errors, ", ")
}
