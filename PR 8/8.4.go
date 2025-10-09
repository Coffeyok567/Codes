// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"strings"
)

// Validator интерфейс для валидации полей
type Validator interface {
	Validate(value string) error
}

// Field структура поля формы
type Field struct {
	Name       string
	Value      string
	Validators []Validator
}

// RequiredValidator проверяет обязательное поле
type RequiredValidator struct{}

func (v RequiredValidator) Validate(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("поле обязательно для заполнения")
	}
	return nil
}

// MinLengthValidator проверяет минимальную длину
type MinLengthValidator struct {
	MinLength int
}

func (v MinLengthValidator) Validate(value string) error {
	if len(value) < v.MinLength {
		return fmt.Errorf("минимальная длина %d символов", v.MinLength)
	}
	return nil
}

// MaxLengthValidator проверяет максимальную длину
type MaxLengthValidator struct {
	MaxLength int
}

func (v MaxLengthValidator) Validate(value string) error {
	if len(value) > v.MaxLength {
		return fmt.Errorf("максимальная длина %d символов", v.MaxLength)
	}
	return nil
}

// EmailValidator проверяет формат email
type EmailValidator struct{}

func (v EmailValidator) Validate(value string) error {
	if !strings.Contains(value, "@") || !strings.Contains(value, ".") {
		return errors.New("неверный формат email")
	}
	return nil
}

// Validate выполняет все валидаторы для поля
func (f *Field) Validate() []error {
	var errors []error
	for _, validator := range f.Validators {
		if err := validator.Validate(f.Value); err != nil {
			errors = append(errors, fmt.Errorf("%s: %v", f.Name, err))
		}
	}
	return errors
}

// Form структура формы
type Form struct {
	Fields []*Field
}

// Validate проверяет все поля формы
func (f *Form) Validate() []error {
	var allErrors []error
	for _, field := range f.Fields {
		if errors := field.Validate(); len(errors) > 0 {
			allErrors = append(allErrors, errors...)
		}
	}
	return allErrors
}

func main() {
	// Создаем форму регистрации
	form := Form{
		Fields: []*Field{
			{
				Name:  "Имя",
				Value: "",
				Validators: []Validator{
					RequiredValidator{},
					MinLengthValidator{MinLength: 2},
					MaxLengthValidator{MaxLength: 50},
				},
			},
			{
				Name:  "Email",
				Value: "invalid-email",
				Validators: []Validator{
					RequiredValidator{},
					EmailValidator{},
				},
			},
			{
				Name:  "Пароль",
				Value: "123",
				Validators: []Validator{
					RequiredValidator{},
					MinLengthValidator{MinLength: 6},
				},
			},
		},
	}

	// Валидируем форму
	fmt.Println("=== Валидация формы регистрации ===")
	errors := form.Validate()
	if len(errors) > 0 {
		fmt.Println("Найдены ошибки:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
	} else {
		fmt.Println("Форма валидна!")
	}

	// Форма с корректными данными
	fmt.Println("\n=== Форма с корректными данными ===")
	validForm := Form{
		Fields: []*Field{
			{
				Name:  "Имя",
				Value: "Иван",
				Validators: []Validator{
					RequiredValidator{},
					MinLengthValidator{MinLength: 2},
				},
			},
			{
				Name:  "Email",
				Value: "ivan@example.com",
				Validators: []Validator{
					RequiredValidator{},
					EmailValidator{},
				},
			},
			{
				Name:  "Пароль",
				Value: "securepassword",
				Validators: []Validator{
					RequiredValidator{},
					MinLengthValidator{MinLength: 6},
				},
			},
		},
	}

	errors = validForm.Validate()
	if len(errors) > 0 {
		fmt.Println("Найдены ошибки:")
		for _, err := range errors {
			fmt.Printf("  - %v\n", err)
		}
	} else {
		fmt.Println("Форма валидна!")
	}
}