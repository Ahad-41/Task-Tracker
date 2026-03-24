package cli

import (
	"fmt"
	"strconv"
	"task-cli/task"
)

func (h *Handler) HandleUpdate(args []string) {
	if len(args) < 4 {
		fmt.Println("Error: Missing task ID or new description.")
		fmt.Println("Usage: task-cli update <id> \"New Description\"")
		return
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	if err := h.service.Update(id, args[3]); err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
		} else {
			fmt.Printf("Error updating task: %s\n", err)
		}
		return
	}
	fmt.Printf("Task %d updated successfully\n", id)
}
