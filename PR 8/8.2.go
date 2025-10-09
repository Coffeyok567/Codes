// Корчагин Евгений 363
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string `json:"database_url"`  //  тут данные
	Port        string `json:"port"`          // наш  порт
	Debug       bool   `json:"debug"`        
}

type ConfigProvider interface {
	Load() (*Config, error)
}

type FileConfigProvider struct {
	filename string  // где спрятаны настройки
}

func (f FileConfigProvider) Load() (*Config, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, fmt.Errorf("файл конфига убежал: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("JSON сломался: %v", err)
	}

	return &config, nil
}

type EnvConfigProvider struct{}  

func (e EnvConfigProvider) Load() (*Config, error) {
	config := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		Debug:       os.Getenv("DEBUG") == "true",  // магия превращения true в true 
	}
	return config, nil
}

func main() {
	// Готовим тестовые данные для демонстрации
	configData := `{
		"database_url": "postgres://localhost:5432/mydb",
		"port": "8080",
		"debug": true
	}`
	
	filename := "config.json"
	os.WriteFile(filename, []byte(configData), 0644)
	fmt.Printf("Файл %s создан! Можно посмотреть его содержимое\n", filename)

	os.Setenv("DATABASE_URL", "postgres://prod:5432/proddb")
	os.Setenv("PORT", "9090")
	os.Setenv("DEBUG", "false")

	var provider ConfigProvider

	fmt.Println("=== Загрузка из файла config.json ===")
	provider = FileConfigProvider{filename: filename}
	fileConfig, err := provider.Load()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("DatabaseURL: %s\n", fileConfig.DatabaseURL)
		fmt.Printf("Port: %s\n", fileConfig.Port)
		fmt.Printf("Debug: %t\n", fileConfig.Debug)
	}

	fmt.Println("\n=== Загрузка из переменных окружения ===")
	provider = EnvConfigProvider{}
	envConfig, err := provider.Load()
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("DatabaseURL: %s\n", envConfig.DatabaseURL)
		fmt.Printf("Port: %s\n", envConfig.Port)
		fmt.Printf("Debug: %t\n", envConfig.Debug)
	}

	// Не удаляем файл, чтобы можно было посмотреть
	fmt.Printf("\nФайл %s остался на диске. Можно открыть его и посмотреть!\n", filename)
	fmt.Println("Содержимое файла:")
	fmt.Println(configData)
}