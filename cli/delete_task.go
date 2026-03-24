package cli

import (
	"fmt"
	"strconv"
	"task-cli/task"
)

func (h *Handler) HandleDelete(args []string) {
	if len(args) < 3 {
		fmt.Println("Error: Missing task ID.")
		fmt.Println("Usage: task-cli delete <id>")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	if err := h.service.Delete(id); err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
		} else {
			fmt.Printf("Error deleting task: %s\n", err)
		}
		return
	}
	fmt.Printf("Task %d deleted successfully\n", id)
}
