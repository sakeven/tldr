package local

import (
	"fmt"
	"log"
	"os"

	"github.com/sakeven/tldr/client"
)

func makeIndex() {
	fmt.Printf("Please wait a second for init local cache\n")

	tldrCli := client.New()

	for _, cmd := range cmds.Cmds {
		for _, p := range cmd.Platform {
			data, err := tldrCli.GetTldr(p, cmd.Name)
			if err != nil {
				log.Println(err)
			}
			f, err := os.Create(tldrPath + "/" + p + "/" + cmd.Name)
			if err != nil {
				log.Printf("Can't create file %s/%s\n", tldrPath, cmd.Name)
				return
			}

			f.WriteString(data)
		}
	}

}
