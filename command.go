package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Update string
	Delete int
	Mark   string
	List   string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new entry")
	flag.StringVar(&cf.Update, "update", "", "Update a Description by -> go run ./ update id description")
	flag.IntVar(&cf.Delete, "delete", 0, "Delete a entry")
	flag.StringVar(&cf.Mark, "mark", "", "Mark a entry")
	flag.StringVar(&cf.List, "list", "", "List a list")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(tasks *Tasks) {
	switch {
	case cf.Add != "":
		tasks.Add(cf.Add)
	case cf.Delete != -1:
		err := tasks.delete(cf.Delete)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case cf.List == "":
		tasks.print(cf.List)
	case cf.List != "":
		tasks.print(cf.List)
	case cf.Update != "":
		parts := strings.SplitN(cf.Update, " ", 2)
		if len(parts) != 2 {
			fmt.Println("Error, Invalid format for edit. Use id:new_title")
			os.Exit(1)
		}
		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, Invalid format for edit. Use id new_description")
			os.Exit(1)
		}
		err = tasks.UpdateDescription(idx, parts[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command")
	}
}
