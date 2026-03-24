package cli

import (
	"fmt"
)

func (h *Handler) HandleList(args []string) {
	status := ""
	if len(args) >= 3 {
		status = args[2]
		if status != "done" && status != "todo" && status != "in-progress" {
			fmt.Println("Error: Invalid status. Use 'done', 'todo', or 'in-progress'.")
			return
		}
	}
	tasks, err := h.service.List(status)
	if err != nil {
		fmt.Printf("Error listing tasks: %s\n", err)
		return
	}

	if len(tasks) == 0 {
		if status != "" {
			fmt.Printf("No tasks found with status: %s\n", status)
		} else {
			fmt.Println("No tasks found.")
		}
		return
	}

	fmt.Printf("%-5s | %-12s | %-30s | %s\n", "ID", "Status", "Description", "Created At")
	fmt.Println("--------------------------------------------------------------------------------")
	for _, t := range tasks {
		desc := t.Description
		if len(desc) > 30 {
			desc = desc[:27] + "..."
		}
		fmt.Printf("%-5d | %-12s | %-30s | %s\n", t.ID, t.Status, desc, t.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
