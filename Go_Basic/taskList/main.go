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

type taskList struct {
	tasks []*task
}

func (tl *taskList) addTaskToList(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) removeTaskFromList(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl taskList) printList() {
	for _, task := range tl.tasks {
		fmt.Print("Name: ", task.name)
		fmt.Println("-- Description: ", task.description)
	}
}

func (tl taskList) printCompletedList() {
	for _, task := range tl.tasks {
		if task.completed {
			fmt.Print("Name: ", task.name)
			fmt.Println("-- Description: ", task.description)
		}
	}
}

func main() {
	t1 := &task{
		name:        "Complete my go course",
		description: "Complete my go course this week",
		completed:   false,
	}

	t2 := &task{
		name:        "Complete my Python course",
		description: "Complete my Python course this week",
		completed:   false,
	}

	t3 := &task{
		name:        "Complete my NodeJS course",
		description: "Complete my NodeJS course this week",
		completed:   false,
	}

	list := taskList{
		tasks: []*task{
			t1,
			t2,
		},
	}
	list.addTaskToList(t3)

	list.printList()

	list.tasks[0].markCompletedTask()
	list.printCompletedList()
}
