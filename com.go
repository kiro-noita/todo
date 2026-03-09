package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func Add() {
	fmt.Print("Enter you task: ")
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	task := strings.TrimSpace(scanner.Text())
	if task == "" {
		fmt.Println("Task cant be empty")
		return
	}
	AddToStorage(task)
	fmt.Printf("Your task: %s was added\n", task)
}

func Show() {
	ShowInStorage()
}

func Mark() {
	fmt.Println("Enter nubmer of task")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		index, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("incorrect input")
			continue
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
		}
		if err := MarkTask(index); err != nil {
			fmt.Printf("Error\n")
			return
		}
		fmt.Printf("You task number: %d is marked\n", index)
		break
	}
}

func Del() {
	fmt.Println("Enter nubmer of task")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		index, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("incorrect input")
			continue
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error")
		}
		if err := DeleteTask(index); err != nil {
			fmt.Printf("Error\n")
			return
		}
		fmt.Printf("You task number: %d deleted\n", index)
		break
	}
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Save() {
	SaveFile()
}

func Load() {
	tasks, err := LoadFile("task.txt")
	if err != nil {
		fmt.Println("Error load")
		os.Exit(1)
	}
	Tasks = tasks
	fmt.Println("Tasks load:", len(tasks))

}
