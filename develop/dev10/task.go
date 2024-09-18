package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
программа должна также завершаться.

При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Int("timeout", 10, "sets timeout")

	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("two arguments needed: host and port")
	}

	args := flag.Args()
	host, port := args[0], args[1]

	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("port should be a number")
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	input := os.Stdin
	output := os.Stdout

	err = telnet(addr, *timeout, input, output)
	if err != nil {
		log.Fatal(err)
	}
}

func telnet(addr string, timeout int, input io.Reader, output io.Writer) (err error) {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	conn, err := net.DialTimeout("tcp", addr, time.Second*time.Duration(timeout))
	if err != nil {
		return err
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	errCh := make(chan error, 1)

	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Bytes()
			_, err = conn.Write(line)
			if err != nil {
				errCh <- err
				return
			}
		}
		if err := scanner.Err(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	go func() {
		defer wg.Done()
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				errCh <- err
				return
			}
			_, err = output.Write(buff[:n])
			if err != nil {
				errCh <- err
				return
			}
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			fmt.Println("Error:", err)
		}
	case <-sigCh:
		fmt.Println("Program closed")
		os.Exit(0)
	}

	wg.Wait()
	return nil
}
