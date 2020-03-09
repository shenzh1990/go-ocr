package settings

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var fconfig = flag.String("fc", "./pkg/settings/fileconfig.yaml", "config file")

type FileConfig struct {
	Seconds     uint64 `yaml:"Seconds"`
	FilePath    string `yaml:"FilePath"`
	TimeSetting int    `yaml:"TimeSetting"`
}

func (c *FileConfig) GetFileConf(filepath string) (*FileConfig, error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return c, nil
}

var FConfig FileConfig

func init() {
	flag.Parse()
	FConfig.GetFileConf(*fconfig)
}
