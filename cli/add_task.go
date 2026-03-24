package cli

import (
	"fmt"
)

func (h *Handler) HandleAdd(args []string) {
	if len(args) < 3 {
		fmt.Println("Error: Missing description for the task.")
		fmt.Println("Usage: task-cli add \"Task description\"")
		return
	}
	t, err := h.service.Add(args[2])
	if err != nil {
		fmt.Printf("Error adding task: %s\n", err)
		return
	}
	fmt.Printf("Task added successfully (ID: %d)\n", t.ID)
}
