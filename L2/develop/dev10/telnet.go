package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Info struct {
	addr    string
	port    string
	timeout int
}

func (i *Info) Flags() {
	timeout := flag.Int("timeout", 10, "timeout time")
	flag.Parse()
	i.timeout = *timeout
	i.addr = os.Args[len(os.Args)-2]
	i.port = os.Args[len(os.Args)-1]
}

func main() {
	flags := new(Info)
	flags.Flags()
	fmt.Println(flags)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)
	// Подключаемся к сокету
	timeout := time.Duration(flags.timeout) * time.Second
	d := net.Dialer{Timeout: timeout}
	conn, err := d.Dial("tcp", flags.addr+":"+flags.port)
	if err != nil {
		log.Fatal("Connection error:\n", err)
	}
	defer conn.Close()
	go func(conn net.Conn) {
		for {
			// Чтение входных данных от stdin
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			// Отправляем в socket
			_, err = fmt.Fprintf(conn, text+"\n")
			if err != nil {
				log.Fatal("Send error:\n", err)
			}
			// Прослушиваем ответ
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: " + message)
		}
	}(conn)
	select {
	case <-quit:
		log.Println("Exit")
		os.Exit(0)
	}
}
