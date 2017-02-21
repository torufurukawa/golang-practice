package main

import (
	"flag"
	"fmt"
	"os"
)

// go run flag-practice.go show --foo=chinko

func main() {
	foo := ""
	showCommand := flag.NewFlagSet("show", flag.ExitOnError)
	showCommand.StringVar(&foo, "foo", "DEFAULT", "foo param")

	if len(os.Args) == 1 {
		fmt.Println("go run flag-practice.go <COMMAND> [--foo=<FOO>]")
		return
	}

	switch os.Args[1] {
	case "show":
		showCommand.Parse(os.Args[2:])
		fmt.Println(foo)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
