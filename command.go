package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMDFlags struct {
	Add    string
	Del    int
	List   string
	Mark   string
	Update string
}

func NewCMDFlags() *CMDFlags {
	cf := CMDFlags{}

	flag.StringVar(&cf.Add, "add", "", `Add a new task -> "-add description"`)
	flag.IntVar(&cf.Del, "del", 0, `Delete a task -> "-del description_id"`)
	flag.StringVar(&cf.List, "list", "", `List tasks -> 
"-list all" for all tasks
"-list todo" for tasks with todo status
"-list in-progress" for tasks with in-progress status
"-list done" for tasks with done status
`)
	flag.StringVar(&cf.Mark, "mark", "", `Update status of a task ->
"-mark id new_status" for a new status
Available Status -> 
	. done 
	. todo
	. in-progress
`)
	flag.StringVar(&cf.Update, "update", "", `Update description of a task -> "-update id new_description"`)

	flag.Parse()

	return &cf
}

func (cf *CMDFlags) Execute(tasks *Tasks) {
	switch {
	case cf.List != "":
		tasks.Print(cf.List)
	case cf.Add != "":
		tasks.Add(cf.Add)
	case cf.Update != "":
		parts := strings.SplitN(cf.Update, " ", 2)
		if len(parts) != 2 {
			fmt.Printf("Error, Invalid format for edit. Use \"id new_description\" \n")
			os.Exit(1)
		}
		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, Invalid index for edit")
			os.Exit(1)
		}
		err = tasks.UpdateDescription(idx, parts[1])
		if err != nil {
			fmt.Printf("Error, %s\n", err)
		}
	case cf.Mark != "":
		parts := strings.SplitN(cf.Mark, " ", 2)
		if len(parts) != 2 {
			fmt.Printf("Error, Invalid format for edit. Use \"id new_status\" \n")
			os.Exit(1)
		}
		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, Invalid index for edit")
			os.Exit(1)
		}
		err = tasks.UpdateStatus(idx, parts[1])
		if err != nil {
			fmt.Printf("Error, %s\n", err)
		}
	case cf.Del != -1:
		err := tasks.Delete(cf.Del)
		if err != nil {
			fmt.Printf("Error, %s\n", err)
		}
	default:
		fmt.Println("No such command")
	}
}
