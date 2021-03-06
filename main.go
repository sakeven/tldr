package main

import (
	"flag"
	"fmt"
	// "log"
	"math/rand"
	"os"
	"time"

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
	frandom   = flag.Bool("r", false, "show a cmd randomly")
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

	if *frandom {
		cmd = getRandomCmd(*fplatform)
	}

	if *fplatform == defaultPlatform && cmd != "" {
		*fplatform = local.GetCmdPlatform(cmd)
	}

	if *fupdate {
		if cmd != "" {
			local.UpdateCmd(*fplatform, cmd)
		} else {
			local.UpdateIndex()
		}
	}

	if *flist {
		printAllCmds(*fplatform)
		return
	}

	if cmd == "" {
		os.Exit(0)
	}

	data, err := getTldr(*fplatform, cmd)
	if err != nil {
		// log.Printf("%s\n", err)
		fmt.Printf("Cmd %s no found. Please make sure you spell it right.\n", cmd)
		os.Exit(1)
	}

	fmt.Printf(Render(data))
}

func printAllCmds(platform string) {
	for _, cmd := range local.GetAllCmds(platform) {
		fmt.Println(cmd)
	}
	os.Exit(0)
}

func getRandomCmd(platform string) string {
	cmds := local.GetAllCmds(platform)
	rand.Seed(time.Now().Unix())
	i := rand.Int() % (len(cmds))

	return cmds[i]
}

func getTldr(platform, cmd string) (string, error) {
	data, _, err := local.GetTldr(platform, cmd)
	if err == nil {
		return data, nil
	}
	fmt.Println(err)

	data, _, err = tldrCli.GetTldr(platform, cmd)

	return data, err
}
