package main

import (
	"flag"
	"fmt"
)

type CMDFlags struct {
	Add  string
	Del  int
	List bool
}

func NewCMDFlags() *CMDFlags {
	cf := CMDFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new command")
	flag.IntVar(&cf.Del, "del", 0, "Delete a command")
	flag.BoolVar(&cf.List, "list", false, "List all commands")

	flag.Parse()

	return &cf
}

func (cf *CMDFlags) Execute(tasks *Tasks) {
	switch {
	case cf.List:
		tasks.Print()
	case cf.Add != "":
		tasks.Add(cf.Add)
	case cf.Del != -1:
		err := tasks.Delete(cf.Del)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("No such command")
	}
}
