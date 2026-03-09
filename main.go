package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Tasks []Task

func main() {
	Tasks = InitStorage()
	fmt.Println("\x1b[3m\x1b[31m► TODO LIST ◄\x1b[0m ")
	fmt.Println("\nList of commands:\n1.add \n2.show \n3.mark \n4.del \n5.save \n6.load \n7.exit \n8.help \n9.clr")

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		command := strings.TrimSpace(scanner.Text())
		if command == "" {
			continue
		}

		switch command {
		case "add":
			Add()
		case "show":
			Show()
		case "mark":
			Mark()
		case "del":
			Del()
		case "save":
			Save()
		case "load":
			Load()
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("\nList of commands:\n1.add \n2.show \n3.mark \n4.del \n5.save \n6.load \n7.exit \n8.help \n9.clr")
		case "clr":
			ClearScreen()
		default:
			fmt.Println("Invalid input")
		}
	}
}
