package local

import (
	"fmt"
	// "log"
	"os"

	"github.com/sakeven/tldr/client"
	// "github.com/sakeven/tldr/conf"
)

var tldrPath = os.Getenv("HOME") + "/.tldr"
var indexPath = fmt.Sprintf("%s/index.json", tldrPath)
var tldrCli = client.New()

func Init() {
	initLocalFile()
}

func initLocalFile() {
	finfo, err := os.Stat(tldrPath)
	if os.IsNotExist(err) {
		os.Mkdir(tldrPath, 0777)
		os.Mkdir(tldrPath+"/"+Common, 0777)
		os.Mkdir(tldrPath+"/"+OSX, 0777)
		os.Mkdir(tldrPath+"/"+Linux, 0777)
		os.Mkdir(tldrPath+"/"+SunOS, 0777)
	} else if finfo.IsDir() == false {
		fmt.Printf("Can't create dir %s", tldrPath)
		os.Exit(1)
	}

	indexFile, err := os.Open(indexPath)
	if err != nil {
		UpdateIndex()
	}
	indexFile.Close()
}
