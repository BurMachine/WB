package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

/*
	Shell утилита, содержит команды - cd/pwd/echo/kill/ps
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		cur, _ := user.Current()
		fmt.Print(cur.Name+"@", "-->  ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		err = Processing(command)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
	}
}

func Processing(str string) error {
	//var err error
	//str1 := strings.Trim(str, "\n")
	//arr := strings.Split(str1, ";")
	//for _, command := range arr {
	//	com := strings.Trim(command, " ")
	//	args := strings.Split(com, " ")
	//	switch args[0] {
	//	case "cd":
	//		if len(args) < 2 {
	//			fmt.Println("path error не написал")
	//			return nil
	//		}
	//
	//	case "exit":
	//		os.Exit(0)
	//	default:
	//		fmt.Println("Еще раз)")
	//	}
	//	cmd := exec.Command("/bin/sh", args[0], args[1])
	//	cmd.Stderr = os.Stderr
	//	cmd.Stdout = os.Stdout
	//	err = cmd.Run()
	//	if err != nil {
	//		return err
	//	}
	//}
	//return err
	var err error
	fmt.Println(str)
	non := strings.Trim(str, "\n")
	commands := strings.Split(non, ";")
	for _, command := range commands {
		com := strings.Trim(command, " ")
		args := strings.Split(com, " ")
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				return errors.New("path required")
			}
			continue
		case "exit":
			os.Exit(0)
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
	}

	return err
}
