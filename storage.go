package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func AddToStorage(s string) {
	Tasks = append(Tasks, Task{s, false})
}

func ShowInStorage() {
	var check string
	for i, task := range Tasks {
		if task.mark {
			check = "☑"
		} else {
			check = "☐"
		}
		fmt.Printf("☻ %d %s %s ☻\n", i+1, task.name, check)
	}
}

func MarkTask(index int) error {
	if index <= len(Tasks) && index > 0 {
		Tasks[index-1].mark = !Tasks[index-1].mark
	} else {
		return errors.New("incorrect input")
	}
	return nil
}

func DeleteTask(index int) error {
	if index <= len(Tasks) && index > 0 {
		Tasks = append(Tasks[:index-1], Tasks[index:]...)
	} else {
		return errors.New("invalid index")
	}
	return nil
}

func SaveFile() {
	file, err := os.Create("task.txt")
	if err != nil {
		fmt.Println("Unable to create file")
		os.Exit(1)
	}
	defer file.Close()
	for _, task := range Tasks {
		check := "☐"
		if task.mark {
			check = "☑"
		}
		line := fmt.Sprintf("%s %s\t%t\n", check, task.name, task.mark)
		file.WriteString(line)
	}
	fmt.Println("Tasks saves")
}

func LoadFile(filen string) ([]Task, error) {
	file, err := os.Open(filen)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	scanner := bufio.NewScanner(file)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		statusStr := parts[len(parts)-1]
		nameParts := parts[1 : len(parts)-1]
		name := strings.Join(nameParts, " ")

		var mark bool
		switch strings.ToLower(statusStr) {
		case "true":
			mark = true
		case "false":
			mark = false
		default:
		}
		tasks = append(tasks, Task{
			name: name,
			mark: mark,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

type Task struct {
	name string
	mark bool
}

func InitStorage() []Task {
	var Tasks []Task
	return Tasks
}
