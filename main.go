package main


import (
	"fmt"
	"github.com/JamesClonk/go-todotxt"
	"os"
	"flag"
)

type todofile string

type TodoTasks struct {
	todotxt.TaskList
	todofile string
}

//TODO: breakout methods into another file for readability

func (todoList todofile) loadList() TodoTasks {
	todotxt.IgnoreComments = false

	tasklist, err := todotxt.LoadFromFilename(string(todoList))
	if err != nil {
		panic(err)
	}

	taskListLocal := TodoTasks{tasklist, string(todoList)}
	return taskListLocal

}

func (tl TodoTasks) addNewTask(task string) {

	parsedTask, err := todotxt.ParseTask(task)
	if err != nil {
		panic(err)
	}

	tl.AddTask(parsedTask)

	tl.WriteToFilename(tl.todofile)

	fmt.Println("New Task successfully added!")

}

func (tl TodoTasks) completeTask(taskID int) {

	task, _ := tl.GetTask(taskID)

	task.Completed = true

	tl.WriteToFilename(tl.todofile)
}

func (tl TodoTasks) rmTask(taskID int) {
	tl.RemoveTaskById(taskID)
	tl.WriteToFilename(tl.todofile)
}

func main() {

	//TODO: make this configurable
	var mainTodoList todofile = "/home/aric/todo.txt"
	
	checkTodoFile(mainTodoList)

	addPtr := flag.String("add", "", "Adds a new task")
	completePtr := flag.Int("complete", 0, "Completes task by ID")
	rmPtr := flag.Int("rm", 0, "Removes Task by ID")

	flag.Parse()
	
	

	taskList := mainTodoList.loadList()

	//Is this a best practice? 
	if *addPtr != "" {
		taskList.addNewTask(*addPtr)

	}

	if *completePtr != 0 {
		taskList.completeTask(*completePtr)
	}

	if *rmPtr != 0 {
		taskList.rmTask(*rmPtr)
	}

	fmt.Println("--Your ToDo List--")
	fmt.Println("------------------")

	//reload the list
	taskList = mainTodoList.loadList()

	for i, v := range taskList.TaskList {
		i = i+1
		fmt.Println(i, v)
	}
	fmt.Println("------------------")

}

func checkTodoFile(file todofile) todofile {
	
	todoListExists, err := os.Stat(string(file))
	if os.IsNotExist(err) {
		fmt.Println("Todo.txt file does not exist. Creating it now...", todoListExists)
		os.Create(string(file))
	} else {
		return file
	}

	return file
}