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
	List   bool
	Edit   string
	Update string
}

func NewCMDFlags() *CMDFlags {
	cf := CMDFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new command")
	flag.IntVar(&cf.Del, "del", 0, "Delete a command")
	flag.BoolVar(&cf.List, "list", false, "List all commands")
	flag.StringVar(&cf.Edit, "edit", "", "Update a command")
	flag.StringVar(&cf.Update, "update", "", "Update a command")

	flag.Parse()

	return &cf
}

func (cf *CMDFlags) Execute(tasks *Tasks) {
	switch {
	case cf.List:
		tasks.Print()
	case cf.Add != "":
		tasks.Add(cf.Add)
	case cf.Update != "":
		parts := strings.SplitN(cf.Update, " ", 2)
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
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, " ", 2)
		if len(parts) != 2 {
			fmt.Printf("Error, Invalid format for edit. Use \"id new_title\" \n")
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
	case cf.Del != -1:
		err := tasks.Delete(cf.Del)
		if err != nil {
			fmt.Printf("Error, %s\n", err)
		}
	default:
		fmt.Println("No such command")
	}
}
