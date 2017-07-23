package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {

	var (
		help = flag.Bool("help", false, "Show all commands")
		psCommand = flag.NewFlagSet("ps", flag.ExitOnError)
		psListAll = psCommand.Bool("a", false, "List all docker containers")
		psHelp = psCommand.Bool("help", false, "Docker containers help")
		imagesCommand = flag.NewFlagSet("images", flag.ExitOnError)
		imagesListAll = imagesCommand.Bool("a", false, "List all docker images")
		imagesHelp = imagesCommand.Bool("help", false, "Docker images help")
	)

	if len(os.Args) > 1 {

		command := os.Args[2:]
		switch os.Args[1] {
		case "ps":
			psCommand.Parse(command)
			if *psListAll {
				fmt.Println("all conainers...")
				return ;
			}
			println("help=", *psHelp)
			psCommand.PrintDefaults()
			return ;
		case "images":
			imagesCommand.Parse(command)
			if *imagesListAll {
				fmt.Println("all images...")
				return;
			}
			println("help=", *imagesHelp)
			imagesCommand.PrintDefaults()
			return ;
		}

	}

	flag.Parse()
	flag.PrintDefaults()
	println("help=", *help)

}

