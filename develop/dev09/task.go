package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func wget(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = file.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("URL required")
		return
	}

	url := os.Args[1]

	err := wget(url)
	if err != nil {
		fmt.Printf("Error downloading file: %s\n", err.Error())
		return
	}
	fmt.Println("File downloaded")
}
