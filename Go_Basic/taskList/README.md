# Task list

This project creates a taskList using structs and slices. The taskList can create new tasks and modify existing tasks. Each task has the following attributes.

- name: The title of the task.
- description: A brief description of the task.
- completed: True if the task has completed successfully.

## How to use it

Move to the project folder and run the project on the terminal with the command: `go run main.go` and it will print all the tasks and the completed tasks.

For adding the task you should add in the main.go file as follows:

    myTask := &task{
		name:        "Complete my NodeJS course",
		description: "Complete my NodeJS course this week",
		completed:   false,
	}

    list.addTaskToList(myTask)