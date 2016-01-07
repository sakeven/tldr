package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	ShortDesc   string `yaml:short_desc`
	ExampleDesc string `yaml:"example_desc`
	ExampleCode string `yaml:"example_code"`
}

const (
	TldrDir   = os.Getenv("HOME") + "/.tldr"
	YamlFile  = tldrDir + "/conf.yaml"
	IndexFile = tldrDir + "/index.json"
)

var conf = new(Conf)

func loadYamlConf() {
	bs, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		// fmt.Printf("error: miss daocloud.yml\n")
		os.Exit(1)
	}

	if err = yaml.Unmarshal(bs, conf); err != nil {
		// fmt.Printf("error: can not decode daocloud.yml, %s\n", err.Error())
		os.Exit(1)
	}
}
