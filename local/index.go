package local

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// UpdateIndex updates index.json and downloads all cmd-tldr
func UpdateIndex() {
	indexFile, err := os.Create(indexPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	index, err := tldrCli.GetIndex()
	if err != nil {
		log.Printf("Init index for tldr failed!")
		os.Exit(1)
	}

	indexFile.WriteString(index)
	err = json.Unmarshal([]byte(index), &cmds)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buildIndex()

}

func buildIndex() {
	fmt.Printf("Please wait a second for init local cache\n")

	for _, cmd := range cmds.Cmds {
		for _, p := range cmd.Platform {
			log.Printf("update cmds %s %s\n", p, cmd.Name)
			UpdateCmd(p, cmd.Name)
		}
	}
}

func loadIndex() {
	indexPath := fmt.Sprintf("%s/index.json", tldrPath)
	indexFile, _ := os.Open(indexPath)

	err := json.NewDecoder(indexFile).Decode(cmds)
	if err != nil {
		log.Printf("Decode index.json error %s\n", err)
		os.Exit(1)
	}
}
