package cmd

import (
	"os"
	"task-cli/cli"
	"task-cli/repo"
	"task-cli/task"
)

func Run() {
	// Storage repository (Adapter)
	taskRepo := repo.NewTaskJSONRepo("tasks.json")

	// Business Logic Service (Core)
	taskService := task.NewService(taskRepo)

	// User Interface Handlers (Adapter)
	cliHandler := cli.NewHandler(taskService)

	// Execute handling pipeline
	cliHandler.Execute(os.Args)
}
