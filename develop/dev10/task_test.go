package main

import (
	"bytes"
	"log"
	"net"
	"sync"
	"testing"
)

var (
	testMsg1 = "msg to server"
	testMsg2 = "msg from server"
)

func newServer(input *bytes.Buffer, output []byte) testServer {
	return testServer{
		input:  input,
		output: output,
	}
}

type testServer struct {
	input  *bytes.Buffer
	output []byte
}

func (ts *testServer) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go ts.readConn(conn, wg)
	go ts.writeConn(conn, wg)
	wg.Wait()
	return nil
}

func TestTelnet(t *testing.T) {
	addr := "localhost:8080"

	serverInput := &bytes.Buffer{}
	serverOutput := []byte(testMsg2)

	server := newServer(serverInput, serverOutput)
	go server.Run(addr)

	telnetInput := bytes.NewBufferString(testMsg1)
	telnetOutput := &bytes.Buffer{}

	telnet(addr, 10, telnetInput, telnetOutput)

	// Проверка, что до сервера доходят правильные сообщения
	if server.input.String() != testMsg1 {
		t.Errorf("wrong server input: got %q, wanted %q", server.input.String(), testMsg1)
	}

	if telnetOutput.String() != string(server.output) {
		t.Errorf("wrong telnet output: got %q, wanted %q", telnetOutput.String(), string(server.output))
	}
}

func (ts *testServer) readConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	// Read data from the client
	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)
	if err != nil {
		return
	}
	_, err = ts.input.Write(bytes[:n])
	if err != nil {
		log.Fatal("failed to write to server input", err)
	}
}

func (ts *testServer) writeConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := conn.Write([]byte(testMsg2))
	if err != nil {
		return
	}
}
