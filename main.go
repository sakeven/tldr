package main

import (
	"flag"
	"fmt"
	// "log"
	"os"

	"github.com/sakeven/tldr/client"
	"github.com/sakeven/tldr/local"
)

var tldrCli = client.New()

const (
	defaultPlatform = "default"
)

var (
	fupdate   = flag.Bool("u", false, "update a cmd or index")
	fplatform = flag.String("p", defaultPlatform, "set sepific platform")
	fhelp     = flag.Bool("h", false, "print usage")
	flist     = flag.Bool("l", false, "list all cmd")
)

func usage() {
	fmt.Println("Usage: tldr -[options] [cmd]\nOptions:")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	flag.Parse()
	args := flag.Args()

	if *fhelp {
		usage()
	}

	local.Init()

	cmd := ""
	if len(args) >= 1 {
		cmd = args[0]
	}

	if *fplatform == defaultPlatform {
		*fplatform = local.GetPlatform(cmd)
	}

	if *fupdate {
		if cmd != "" {
			local.UpdateCmd(*fplatform, cmd)
		} else {
			local.UpdateIndex()
		}
	}

	if *flist {
		printAllCmds()
		return
	}

	data, err := getTldr(*fplatform, cmd)
	if err != nil {
		// log.Printf("%s\n", err)
		fmt.Printf("Cmd %s no found. Please make sure you spell it right\n", cmd)
		os.Exit(1)
	}

	fmt.Printf(Render(data))
}

func printAllCmds() {
	for _, cmd := range local.GetAllCmds() {
		fmt.Println(cmd)
	}
	os.Exit(0)
}

func getTldr(platform, cmd string) (string, error) {
	data, err := local.GetTldr(platform, cmd)
	if err == nil {
		return data, nil
	}
	return tldrCli.GetTldr(platform, cmd)
}
