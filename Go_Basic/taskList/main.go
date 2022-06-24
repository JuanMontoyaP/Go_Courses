package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) markCompletedTask() {
	t.completed = true
}

func main() {
	t := task{
		name:        "Complete my go course",
		description: "Complete my go course this week",
		completed:   false,
	}

	fmt.Printf("%+v\n", t)

	t.markCompletedTask()
	t.updateName("Finish my go course")
	t.updateDescription("Finish my go course asap")

	fmt.Printf("%+v\n", t)
}
