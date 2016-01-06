package local

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "runtime"

    "github.com/sakeven/tldr/client"
)

type Command struct {
    Name     string   `json:"name"`
    Platform []string `json:"platform"`
}

type Commands struct {
    Cmds []*Command `json:"commands"`
}

var cmds = &Commands{}

var tldrPath = os.Getenv("HOME") + "/.tldr"
var needMakeIndex bool

func Init() {
    initLocalFile()
    initIndex()
    if needMakeIndex {
        makeIndex()
    }
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

    indexPath := fmt.Sprintf("%s/index.json", tldrPath)
    indexFile, err := os.Open(indexPath)
    if err != nil {
        indexFile, _ = os.Create(indexPath)
        tldrCli := client.New()
        index, err := tldrCli.GetIndex()
        if err != nil {
            log.Printf("Init index for tldr failed!")
            os.Exit(1)
        }
        indexFile.WriteString(index)
        needMakeIndex = true
    }
    indexFile.Close()
}

func initIndex() {
    indexPath := fmt.Sprintf("%s/index.json", tldrPath)
    indexFile, _ := os.Open(indexPath)

    err := json.NewDecoder(indexFile).Decode(cmds)
    if err != nil {
        log.Printf("Decode index.json error %s\n", err)
        os.Exit(1)
    }
}

const (
    Common = "common"
    OSX    = "osx"
    Linux  = "linux"
    SunOS  = "sunos"
)

func getOS() string {
    switch runtime.GOOS {
    case "darwin":
        return OSX
    case "linux":
        return Linux
    }
    return Common
}

func GetPlatform(cmd string) string {
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
