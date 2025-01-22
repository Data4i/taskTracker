package main

func main() {
	tasks := Tasks{}
	storage := NewStorage[Tasks]("tasks.json")
	loadErr := storage.Load(&tasks)
	if loadErr != nil {
		panic(loadErr)
	}
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&tasks)
	saveErr := storage.Save(tasks)
	if saveErr != nil {
		panic(saveErr)
	}
}
