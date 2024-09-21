package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор,
пока не будет введена команда выхода (например \quit).
*/

func main() {
	for {
		fmt.Print("input: ")
		command, args := readInput()
		switch command {
		case "pwd":
			curDir, err := pwd()
			if err != nil {
				fmt.Println("error: ", err.Error())
				break
			}
			fmt.Println(curDir)
		case "cd":
			if len(args) != 1 {
				fmt.Println("error amonut of argumens")
				break
			}
			dir := args[0]
			err := cd(dir)
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		case "exit":
			return
		case "echo":
			if len(args) != 1 {
				fmt.Println("error amonut of argumens")
				break
			}
			text := args[0]
			echo(text)
		case "ps":
			if len(args) != 0 {
				fmt.Println("error amonut of argumens")
				break
			}
			output, err := ps()
			if err != nil {
				fmt.Println("error: ", err.Error())
				break
			}
			fmt.Println(output)
		case "kill":
			if len(args) != 1 {
				fmt.Println("error amonut of argumens")
				break
			}
			pid := args[0]
			err := kill(pid)
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		case "exec":
			if len(args) != 1 {
				fmt.Println("error amonut of argumens")
				break
			}
			err := execCMD(args)
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		case "fork":
			if len(args) != 1 {
				fmt.Println("error amonut of argumens")
				break
			}
			pid := args[0]
			err := fork(pid)
			if err != nil {
				fmt.Println("error: ", err.Error())
			}
		}
	}
}

func readInput() (command string, args []string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	cleanString := strings.TrimSpace(input)
	inputArgs := strings.Split(cleanString, " ")
	command = inputArgs[0]
	args = inputArgs[1:]
	return
}

func pwd() (string, error) {
	return os.Getwd()
}

func cd(dir string) error {
	return os.Chdir(dir)
}

func echo(text string) {
	fmt.Println(text)
}

func ps() (string, error) {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func kill(pid string) error {
	cmd := exec.Command("kill", "-9", pid)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func fork(pid string) error {
	cmd := exec.Command("fork", pid)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func execCMD(args []string) error {
	cmd := exec.Command(args[0], args[:1]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
