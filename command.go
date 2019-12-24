package main

import (
	"flag"
	"fmt"
	"os"
)


var name string
var cmdLine = flag.NewFlagSet("my question", flag.ExitOnError)

func init(){
	//flag.Usage = usage


	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//flag.StringVar(&name, "name", "hello", "your name")
	cmdLine.StringVar(&name, "name", "hello", "your name")

}


func usage(){
	fmt.Fprintf(os.Stderr, "Usage of %s: \n", "question")
	flag.PrintDefaults()
}

func main() {
	//flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Println("cmd demo")
	fmt.Println("your name:", name)

}
