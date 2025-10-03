// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	DueDate     *time.Time
	Completed   bool
	CompletedAt *time.Time
	Priority    Priority
	Category    string
}

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type TaskManager struct {
	tasks  map[int]*Task
	nextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (tm *TaskManager) AddTask(title, description, category string, priority Priority, dueDate *time.Time) *Task {
	task := &Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		DueDate:     dueDate,
		Completed:   false,
		Priority:    priority,
		Category:    category,
	}

	tm.tasks[tm.nextID] = task
	fmt.Printf("Добавлена задача: '%s' (ID: %d, Приоритет: %s)\n", title, tm.nextID, tm.getPriorityText(priority))
	tm.nextID++
	return task
}

func (tm *TaskManager) DeleteTask(id int) error {
	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("задача не найдена")
	}

	delete(tm.tasks, id)
	fmt.Printf("Удалена задача: '%s' (ID: %d)\n", task.Title, id)
	return nil
}

// Задача есть и выполнена  ну проверка короч
func (tm *TaskManager) MarkDone(id int) error {
	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("задача не найдена")
	}

	if task.Completed {
		return errors.New("задача уже выполнена")
	}

	task.Completed = true
	now := time.Now()
	task.CompletedAt = &now
	fmt.Printf("Задача выполнена: '%s' (ID: %d)\n", task.Title, id)
	return nil
}

// задачи нет или не выполнена  тож проврека
func (tm *TaskManager) MarkUndone(id int) error {
	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("задача не найдена")
	}

	if !task.Completed {
		return errors.New("задача еще не выполнена")
	}

	task.Completed = false
	task.CompletedAt = nil
	fmt.Printf("Задача возвращена в работу: '%s' (ID: %d)\n", task.Title, id)
	return nil
}

// обновление стптуса задчи
func (tm *TaskManager) UpdateTask(id int, title, description, category string, priority Priority, dueDate *time.Time) error {
	task, exists := tm.tasks[id]
	if !exists {
		return errors.New("задача не найдена")
	}

	task.Title = title
	task.Description = description
	task.Category = category
	task.Priority = priority
	task.DueDate = dueDate

	fmt.Printf("Обновлена задача: '%s' (ID: %d)\n", title, id)
	return nil
}

func (tm *TaskManager) GetTasks() []*Task {
	tasks := make([]*Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// получение активных задач
func (tm *TaskManager) GetActiveTasks() []*Task {
	var active []*Task
	for _, task := range tm.tasks {
		if !task.Completed {
			active = append(active, task)
		}
	}
	return active
}

// полцучение выполненых задач
func (tm *TaskManager) GetCompletedTasks() []*Task {
	var completed []*Task
	for _, task := range tm.tasks {
		if task.Completed {
			completed = append(completed, task)
		}
	}
	return completed
}

// по категория
func (tm *TaskManager) GetTasksByCategory(category string) []*Task {
	var result []*Task
	for _, task := range tm.tasks {
		if task.Category == category {
			result = append(result, task)
		}
	}
	return result
}

// по приоритету жи ес
func (tm *TaskManager) GetTasksByPriority(priority Priority) []*Task {
	var result []*Task
	for _, task := range tm.tasks {
		if task.Priority == priority {
			result = append(result, task)
		}
	}
	return result
}

func (tm *TaskManager) GetOverdueTasks() []*Task {
	var overdue []*Task
	now := time.Now()
	for _, task := range tm.tasks {
		if !task.Completed && task.DueDate != nil && task.DueDate.Before(now) {
			overdue = append(overdue, task)
		}
	}
	return overdue
}

// поиск
func (tm *TaskManager) GetTask(id int) (*Task, error) {
	task, exists := tm.tasks[id]
	if !exists {
		return nil, errors.New("задача не найдена")
	}
	return task, nil
}

// приоритет
func (tm *TaskManager) getPriorityText(priority Priority) string {
	priorityMap := map[Priority]string{
		PriorityLow:    "Низкий",
		PriorityMedium: "Средний",
		PriorityHigh:   "Высокий",
	}
	return priorityMap[priority]
}

// получение статуса
func (tm *TaskManager) getStatusText(completed bool) string {
	if completed {
		return "✅ Выполнена"
	}
	return "⏳ В работе"
}

// вывод
func (tm *TaskManager) DisplayTasks(tasks []*Task, title string) {
	fmt.Printf("\n=== %s ===\n", title)

	if len(tasks) == 0 {
		fmt.Println("Задачи не найдены")
	} else {
		for _, task := range tasks {
			tm.displaySingleTask(task)
		}
	}
	fmt.Printf("Всего задач: %d\n", len(tasks))
	fmt.Printf("======================\n\n")
}

// вывод и немного красоты при самом выводе
func (tm *TaskManager) displaySingleTask(task *Task) {
	dueDateStr := "Не установлен"
	if task.DueDate != nil {
		dueDateStr = task.DueDate.Format("02.01.2006")
	}

	completedAtStr := ""
	if task.Completed && task.CompletedAt != nil {
		completedAtStr = fmt.Sprintf(" (завершена: %s)", task.CompletedAt.Format("02.01.2006 15:04"))
	}

	fmt.Printf("ID: %d | %s\n", task.ID, tm.getStatusText(task.Completed))
	fmt.Printf("  Заголовок: %s\n", task.Title)
	fmt.Printf("  Описание: %s\n", task.Description)
	fmt.Printf("  Категория: %s | Приоритет: %s\n", task.Category, tm.getPriorityText(task.Priority))
	fmt.Printf("  Срок выполнения: %s%s\n", dueDateStr, completedAtStr)
	fmt.Printf("  Создана: %s\n", task.CreatedAt.Format("02.01.2006"))
	fmt.Println("  ---")
}

func (tm *TaskManager) DisplayStats() {
	allTasks := tm.GetTasks()
	activeTasks := tm.GetActiveTasks()
	completedTasks := tm.GetCompletedTasks()
	overdueTasks := tm.GetOverdueTasks()

	fmt.Printf("\n=== Статистика задач ===\n")
	fmt.Printf("Всего задач: %d\n", len(allTasks))
	fmt.Printf("Активных: %d\n", len(activeTasks))
	fmt.Printf("Выполненных: %d\n", len(completedTasks))
	fmt.Printf("Просроченных: %d\n", len(overdueTasks))

	// Статистика по приоритетам
	highPriority := len(tm.GetTasksByPriority(PriorityHigh))
	mediumPriority := len(tm.GetTasksByPriority(PriorityMedium))
	lowPriority := len(tm.GetTasksByPriority(PriorityLow))
	fmt.Printf("Приоритеты: Высокий - %d, Средний - %d, Низкий - %d\n",
		highPriority, mediumPriority, lowPriority)
	fmt.Printf("========================\n\n")
}

func main() {
	taskManager := NewTaskManager()

	// Добавляем задачи
	fmt.Println("--- Добавление задач ---")
	due1 := time.Now().Add(24 * time.Hour)
	due2 := time.Now().Add(7 * 24 * time.Hour)
	due3 := time.Now().Add(-24 * time.Hour) // Просроченная задача

	taskManager.AddTask("Купить продукты", "Молоко, хлеб, яйца", "Покупки", PriorityHigh, &due1)
	taskManager.AddTask("Сделать домашнее задание", "Математика и физика", "Учеба", PriorityMedium, &due2)
	taskManager.AddTask("Записаться к врачу", "Терапевт на следующей неделе", "Здоровье", PriorityHigh, nil)
	taskManager.AddTask("Прочитать книгу", "'Гарри Поттер' до конца месяца", "Личное", PriorityLow, &due3)
	taskManager.AddTask("Убраться в комнате", "Пропылесосить и протереть пыль", "Дом", PriorityMedium, nil)

	taskManager.DisplayStats()

	// Показываем все задачи
	taskManager.DisplayTasks(taskManager.GetTasks(), "Все задачи")

	// Отмечаем выполнение задачи
	fmt.Println("--- Отметка выполнения ---")
	err := taskManager.MarkDone(1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	err = taskManager.MarkDone(3)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	taskManager.DisplayStats()

	// Показываем активные задачи
	taskManager.DisplayTasks(taskManager.GetActiveTasks(), "Активные задачи")

	// Показываем выполненные задачи
	taskManager.DisplayTasks(taskManager.GetCompletedTasks(), "Выполненные задачи")

	// Фильтрация по категории
	fmt.Println("--- Фильтрация по категории 'Учеба' ---")
	taskManager.DisplayTasks(taskManager.GetTasksByCategory("Учеба"), "Задачи по учебе")

	// Фильтрация по приоритету
	fmt.Println("--- Фильтрация по приоритету 'Высокий' ---")
	taskManager.DisplayTasks(taskManager.GetTasksByPriority(PriorityHigh), "Задачи с высоким приоритетом")

	// Просроченные задачи
	fmt.Println("--- Просроченные задачи ---")
	taskManager.DisplayTasks(taskManager.GetOverdueTasks(), "Просроченные задачи")

	// Обновление задачи
	fmt.Println("--- Обновление задачи ---")
	newDue := time.Now().Add(3 * 24 * time.Hour)
	err = taskManager.UpdateTask(2, "Сделать домашнее задание", "Математика, физика и программирование", "Учеба", PriorityHigh, &newDue)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Возврат задачи в работу
	fmt.Println("--- Возврат задачи в работу ---")
	err = taskManager.MarkUndone(1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Показываем обновленный список
	taskManager.DisplayTasks(taskManager.GetTasks(), "Обновленный список задач")

	// Удаление задачи
	fmt.Println("--- Удаление задачи ---")
	err = taskManager.DeleteTask(5)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	taskManager.DisplayStats()

	// Демонстрация обработки ошибок
	fmt.Println("--- Демонстрация обработки ошибок ---")
	err = taskManager.MarkDone(999)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	err = taskManager.DeleteTask(999)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	err = taskManager.MarkUndone(999)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
