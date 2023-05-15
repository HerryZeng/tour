package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	var name string
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "name", "PHP语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	//flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	//flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	//flag.Parse()

	log.Printf("name: %s", name)
}
