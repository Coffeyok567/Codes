// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Ingredient struct {
	Name   string
	Amount string
}

type Recipe struct {
	ID          int
	Title       string
	Description string
	Ingredients []Ingredient
	Steps       []string
	Category    string
	PrepTime    int
}

type RecipeBook struct {
	recipes map[int]*Recipe
	nextID  int
}

func NewRecipeBook() *RecipeBook {
	return &RecipeBook{
		recipes: make(map[int]*Recipe),
		nextID:  1,
	}
}

func (rb *RecipeBook) CreateRecipe(title, description, category string, prepTime int) *Recipe {
	recipe := &Recipe{
		ID:          rb.nextID,
		Title:       title,
		Description: description,
		Category:    category,
		PrepTime:    prepTime,
	}
	rb.recipes[rb.nextID] = recipe
	rb.nextID++
	fmt.Printf("Создан рецепт: '%s'\n", title)
	return recipe
}

func (r *Recipe) AddIngredient(name, amount string) {
	r.Ingredients = append(r.Ingredients, Ingredient{Name: name, Amount: amount})
}

func (r *Recipe) AddStep(step string) {
	r.Steps = append(r.Steps, step)
}

func (rb *RecipeBook) FilterByCategory(category string) []*Recipe {
	var result []*Recipe
	for _, recipe := range rb.recipes {
		if strings.EqualFold(recipe.Category, category) {
			result = append(result, recipe)
		}
	}
	return result
}

func (rb *RecipeBook) FilterByIngredient(ingredientName string) []*Recipe {
	var result []*Recipe
	for _, recipe := range rb.recipes {
		for _, ing := range recipe.Ingredients {
			if strings.Contains(strings.ToLower(ing.Name), strings.ToLower(ingredientName)) {
				result = append(result, recipe)
				break
			}
		}
	}
	return result
}

func (rb *RecipeBook) FindLongestRecipe() (*Recipe, error) {
	if len(rb.recipes) == 0 {
		return nil, errors.New("нет рецептов")
	}
	var longest *Recipe
	for _, recipe := range rb.recipes {
		if longest == nil || recipe.PrepTime > longest.PrepTime {
			longest = recipe
		}
	}
	return longest, nil
}

func (rb *RecipeBook) DisplayRecipe(recipe *Recipe) {
	fmt.Printf("\n=== %s ===\n", recipe.Title)
	fmt.Printf("Категория: %s | Время: %d мин\n", recipe.Category, recipe.PrepTime)
	fmt.Printf("Ингредиенты:\n")
	for _, ing := range recipe.Ingredients {
		fmt.Printf("  - %s: %s\n", ing.Name, ing.Amount)
	}
	fmt.Printf("Шаги:\n")
	for i, step := range recipe.Steps {
		fmt.Printf("  %d. %s\n", i+1, step)
	}
}

func main() {
	rb := NewRecipeBook()

	// Создаем рецепты
	caesar := rb.CreateRecipe("Салат Цезарь", "Классический салат", "Салаты", 30)
	caesar.AddIngredient("Курица", "300 г")
	caesar.AddIngredient("Салат", "1 пучок")
	caesar.AddIngredient("Помидоры черри", "150 г")
	caesar.AddStep("Приготовьте курицу")
	caesar.AddStep("Нарежьте овощи")
	caesar.AddStep("Смешайте ингредиенты")

	borscht := rb.CreateRecipe("Борщ", "Наваристый суп", "Супы", 120)
	borscht.AddIngredient("Говядина", "500 г")
	borscht.AddIngredient("Свекла", "2 шт.")
	borscht.AddIngredient("Капуста", "300 г")
	borscht.AddStep("Варите мясо 1.5 часа")
	borscht.AddStep("Добавьте овощи")
	borscht.AddStep("Варите до готовности")

	// Поиск и фильтрация
	fmt.Println("\n--- Салаты ---")
	salads := rb.FilterByCategory("Салаты")
	for _, r := range salads {
		fmt.Printf("- %s\n", r.Title)
	}

	fmt.Println("\n--- Рецепты с курицей ---")
	chicken := rb.FilterByIngredient("кури")
	for _, r := range chicken {
		fmt.Printf("- %s\n", r.Title)
	}

	fmt.Println("\n--- Самый долгий рецепт ---")
	longest, _ := rb.FindLongestRecipe()
	rb.DisplayRecipe(longest)
}
