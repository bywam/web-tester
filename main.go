package main

import (
	"fmt"
	"os"

	"web-tester/tasks"
)

func main() {
	mainTaskFiles := []string{"t0"}

	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "t0":
			tasks.RunTasksT0(cmd)
		}
	} else {
		tasks := fmt.Sprint(mainTaskFiles)
		fmt.Printf("Please provide the task name: %s", tasks)
	}
}
