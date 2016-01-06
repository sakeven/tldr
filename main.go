package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sakeven/tldr/client"
	"github.com/sakeven/tldr/local"
)

var tldrCli = client.New()

func usage() {
	fmt.Println("Usage: tldr <cmd>")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}

	local.Init()

	cmd := os.Args[1]
	platform := local.GetPlatform(cmd)

	data, err := getTldr(platform, cmd)
	if err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf(Render(data))
}

func getTldr(platform, cmd string) (string, error) {
	data, err := local.GetTldr(platform, cmd)
	if err == nil {
		return data, nil
	}
	return tldrCli.GetTldr(platform, cmd)
}
