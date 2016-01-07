package local

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

type Command struct {
	Name     string   `json:"name"`
	Platform []string `json:"platform"`
}

type Commands struct {
	Cmds []*Command `json:"commands"`
}

var cmds = &Commands{}

const (
	Common = "common"
	OSX    = "osx"
	Linux  = "linux"
	SunOS  = "sunos"
)

// getOS gets user's operation system.
func getOS() string {
	switch runtime.GOOS {
	case "darwin":
		return OSX
	case "linux":
		return Linux
	}
	return Common
}

// GetPlatform returns cmd's platform
func GetPlatform(cmd string) string {
	loadIndex()
	for _, c := range cmds.Cmds {
		if c.Name == cmd {
			for _, p := range c.Platform {
				if p == getOS() {
					return p
				}
			}
			return c.Platform[0]
		}
	}
	return ""
}

func GetTldr(platform, cmd string) (string, error) {
	cmdPath := fmt.Sprintf("%s/%s/%s", tldrPath, platform, cmd)
	data, err := ioutil.ReadFile(cmdPath)
	return string(data), err
}

func GetAllCmds() []string {
	var cs []string
	loadIndex()
	for _, cmd := range cmds.Cmds {
		cs = append(cs, cmd.Name)
	}

	return cs
}

func UpdateCmd(platform string, cmd string) {
	data, err := tldrCli.GetTldr(platform, cmd)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Create(tldrPath + "/" + platform + "/" + cmd)
	if err != nil {
		log.Printf("Can't create file %s/%s\n", tldrPath, cmd)
		return
	}

	f.WriteString(data)
}
