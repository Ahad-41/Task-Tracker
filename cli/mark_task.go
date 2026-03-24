package cli

import (
	"fmt"
	"strconv"
	"task-cli/task"
)

func (h *Handler) HandleMark(args []string, status string) {
	if len(args) < 3 {
		fmt.Println("Error: Missing task ID.")
		commandName := "mark-" + status
		fmt.Printf("Usage: task-cli %s <id>\n", commandName)
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	if err := h.service.MarkStatus(id, status); err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
		} else {
			fmt.Printf("Error marking task: %s\n", err)
		}
		return
	}

	outputStatus := status
	switch status {
	case "in-progress":
		outputStatus = "as in-progress"
	case "done":
		outputStatus = "as done"
	}
	fmt.Printf("Task %d marked %s\n", id, outputStatus)
}
