package main

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	filename := "index.html"
	url := "https://www.youtube.com/" + filename

	err := wget(url)
	if err != nil {
		t.Errorf("Ошибка при скачивании файла: %s", err.Error())
	}

	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("Файл не был создан")
	}

	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Ошибка при удалении файла: %s", err.Error())
	}
}
