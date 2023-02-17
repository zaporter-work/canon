package main

import (
	"fmt"
	"os"
	"flag"
)

var defaultArgs = []string{"bash", "-l"}

func main() {
	parseConfigs()
	args := flag.Args()
	if len(args) <= 0 {
		checkErr(shell(defaultArgs))
	} else {
		switch args[0] {
		case "shell":
			checkErr(shell(defaultArgs))
		case "config":
			showConfig(activeProfile)
		case "update":
			all := false
			if len(args) >= 2 {
				if args[1] == "-a" || args[1] == "-all" || args[1] == "--all" {
					all = true
				}
			}
			runUpdate(all)
		case "--":
			fallthrough
		case "run":
			checkErr(shell(args[1:]))
		default:
			checkErr(shell(args))
		}
	}
}

func checkErr(err error) {
	if err == nil {
		return
	}
	_, err2 := fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	if err2 != nil {
		fmt.Printf("Error encountered printing to stderr: %s\nOriginal Error: %s", err2, err)
	}
}
