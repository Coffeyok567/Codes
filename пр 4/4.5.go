//Корчагин Евгений 363

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type Result struct {
	FilePath string
	Hash     string
	Err      error
}

func worker(id int, wg *sync.WaitGroup, files <-chan string, results chan<- Result) {
	defer wg.Done()

	for file := range files {
		fmt.Printf("Воркер %d обрабатывает файл: %s\n", id, file)

		hash, err := calculateMD5(file)

		results <- Result{
			FilePath: file,
			Hash:     hash,
			Err:      err,
		}
	}

	fmt.Printf("Воркер %d завершил работу\n", id)
}

func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func createTestFile(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf("Это тестовое содержимое файла %s", filePath)
	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
		"file4.txt",
		"file5.txt",
		"file6.txt",
		"file7.txt",
		"file8.txt",
	}

	fileChan := make(chan string, len(files))
	resultChan := make(chan Result, len(files))

	var wg sync.WaitGroup

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, &wg, fileChan, resultChan)
	}

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("Файл %s не существует, создаем...\n", file)
			if err := createTestFile(file); err != nil {
				fmt.Printf("Ошибка создания файла %s: %v\n", file, err)
				continue
			}
		}
		fileChan <- file
	}

	close(fileChan)
	wg.Wait()
	close(resultChan)

	fmt.Println("\nРезультаты вычисления хеш-сумм:")
	for result := range resultChan {
		if result.Err != nil {
			fmt.Printf("Ошибка при обработке файла %s: %v\n", result.FilePath, result.Err)
		} else {
			fmt.Printf("%s: %s\n", result.FilePath, result.Hash)
		}
	}
}
