// Корчагин Евгений 363
package main

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Info(msg string) {
	fmt.Printf("[INFO] %s %s\n", time.Now().Format("15:04:05"), msg)
}

func (c ConsoleLogger) Error(msg string) {
	fmt.Printf("[ERROR] %s %s\n", time.Now().Format("15:04:05"), msg)
}

func (c ConsoleLogger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s %s\n", time.Now().Format("15:04:05"), msg)
}

type FileLogger struct {
	filename string
}

func (f FileLogger) log(level string, msg string) {
	file, err := os.OpenFile(f.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("[%s] %s %s\n", level, time.Now().Format("2006-01-02 15:04:05"), msg)
	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Printf("Ошибка записи в файл: %v\n", err)
	}
}

func (f FileLogger) Info(msg string) {
	f.log("INFO", msg)
}

func (f FileLogger) Error(msg string) {
	f.log("ERROR", msg)
}

func (f FileLogger) Debug(msg string) {
	f.log("DEBUG", msg)
}

func main() {
	var logger Logger

	fmt.Println("=== Логи в консоли ===")
	logger = ConsoleLogger{}
	logger.Info("Приложение запущено")
	logger.Debug("Отладочная информация")
	logger.Error("Произошла ошибка")

	fmt.Println("\n===  Логи в файле ===")
	logger = FileLogger{filename: "app.log"}
	logger.Info("Приложение запущено")
	logger.Debug("Отладочная информация")
	logger.Error("Произошла ошибка")
	
	fmt.Println("Логи записаны в файл app.log")
}