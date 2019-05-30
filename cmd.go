package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	args        []string
	class       string
	cpOption    string
	XjreOption  string
	helpFlag    bool
	versionFlag bool
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = pringUsage
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func pringUsage() {
	fmt.Print("Usage: %s [-options] class [args...]\n", os.Args[0])
}
