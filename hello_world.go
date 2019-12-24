package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// flag.StringVar(&name, "name", "everyone", "The greeting object.")
	cmdLine.StringVar(&name, "name", "everyone", "The geeting object.")
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
}

func main() {
	// flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("hello %s!\n", name)
}
