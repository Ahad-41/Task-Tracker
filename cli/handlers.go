package cli

import (
	"fmt"
	"task-cli/domain"
)

type TaskManager interface {
	Add(description string) (domain.Task, error)
	Update(id int, description string) error
	Delete(id int) error
	MarkStatus(id int, status string) error
	List(filterStatus string) ([]domain.Task, error)
}

type Handler struct {
	service TaskManager
}

func NewHandler(service TaskManager) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Execute(args []string) {
	if len(args) < 2 {
		h.printUsage()
		return
	}

	command := args[1]

	switch command {
	case "add":
		h.HandleAdd(args)
	case "update":
		h.HandleUpdate(args)
	case "delete":
		h.HandleDelete(args)
	case "mark-in-progress":
		h.HandleMark(args, "in-progress")
	case "mark-done":
		h.HandleMark(args, "done")
	case "list":
		h.HandleList(args)
	default:
		fmt.Printf("Error: Unknown command '%s'\n", command)
		h.printUsage()
	}
}

func (h *Handler) printUsage() {
	fmt.Println("Task Tracker CLI")
	fmt.Println("Usage:")
	fmt.Println("  task-cli add \"Buy groceries\"")
	fmt.Println("  task-cli update 1 \"Buy groceries and cook dinner\"")
	fmt.Println("  task-cli delete 1")
	fmt.Println("  task-cli mark-in-progress 1")
	fmt.Println("  task-cli mark-done 1")
	fmt.Println("  task-cli list")
	fmt.Println("  task-cli list done")
	fmt.Println("  task-cli list todo")
	fmt.Println("  task-cli list in-progress")
}
