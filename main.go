package main

func main() {
	tasks := Tasks{}

	storage := NewStorage[Tasks]("tasks.json")
	err := storage.Load(&tasks)
	if err != nil {
		panic(err)
	}

	cmdFlags := NewCMDFlags()
	cmdFlags.Execute(&tasks)
	_ = storage.Save(tasks)

}
