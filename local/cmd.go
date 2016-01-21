package local

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
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
	Common  = "common"
	OSX     = "osx"
	Linux   = "linux"
	SunOS   = "sunos"
	Default = "default"
)

// getOS gets user's operation system.
func getOS() string {
	switch runtime.GOOS {
	case "darwin":
		return OSX
	case "linux":
		return Linux
	case "sunos":
		return SunOS
	}
	return Common
}

// GetCmdPlatform returns cmd's Platform
func GetCmdPlatform(cmd string) string {
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

func GetTldr(platform, cmd string) (data string, etag string, err error) {
	cmdPath := fmt.Sprintf("%s/%s/%s", tldrPath, platform, cmd)
	b, err := ioutil.ReadFile(cmdPath)
	if err != nil {
		return "", "", err
	}

	meta := strings.SplitN(string(b), "\n", 2)
	if len(meta) != 2 {
		return "", "", fmt.Errorf("get local tldr error")
	}
	return meta[1], meta[0], err
}

func GetAllCmds(platform string) []string {
	var cs []string
	loadIndex()
	for _, cmd := range cmds.Cmds {
		for _, p := range cmd.Platform {
			if p == platform || platform == Default {
				cs = append(cs, cmd.Name)
				break
			}
		}
	}

	return cs
}

func UpdateCmd(platform string, cmd string) {
	_, etag, _ := GetTldr(platform, cmd)
	updateCmd(platform, cmd, etag)

}

func updateCmd(platform string, cmd string, etag string) {
	if ok, _ := tldrCli.IsExpired(platform, cmd, etag); !ok {
		return
	}

	data, etag, err := tldrCli.GetTldr(platform, cmd)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Create(tldrPath + "/" + platform + "/" + cmd)
	if err != nil {
		log.Printf("Can't create file %s/%s\n", tldrPath, cmd)
		return
	}

	f.WriteString(etag + "\n")
	f.WriteString(data)

	return
}
