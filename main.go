package main

func main() {
	tasks := Tasks{}
	storage := NewStorage[Tasks]("tasks.json")
	err := storage.Load(&tasks)
	if err != nil {
		panic(err)
	}

	tasks.Add("Build Task Tracker")
	tasks.Add("Eat Banana")
	tasks.Add("Put it on Github")
	tasks.Add("Build a documentation")
	tasks.Add("Upload it on roadmap.io")

	_ = tasks.UpdateDescription(2, "Get it out on Github")
	_ = tasks.delete(1)
	_ = tasks.UpdateStatus(0, "In Progress")
	_ = tasks.UpdateStatus(1, "Done")

	tasks.print("")

	err = storage.Save(tasks)
	if err != nil {
		panic(err)
	}
}
