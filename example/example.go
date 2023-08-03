package main

import (
	"fmt"
	configparser "gconf"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type ExampleConfig struct {
	StringV   string             `yaml:"string_v"`
	IntV      int                `yaml:"int_v"`
	BoolV     bool               `yaml:"bool_v"`
	FloatV    float64            `yaml:"float_v"`
	SubConfig []ExampleSubConfig `yaml:"sub_config"`
}

type ExampleSubConfig struct {
	StringList []string  `yaml:"string_list"`
	IntList    []int     `yaml:"int_list"`
	BoolList   []bool    `yaml:"bool_list"`
	FloatList  []float64 `yaml:"float_list"`
}

func main() {
	runPath, _ := os.Getwd()
	confPath := path.Join(runPath, "conf/example.yaml")

	env := []string{"dev", "test", "pre", "pro"}
	for _, v := range env {
		var conf ExampleConfig
		curConfig, err := configparser.LoadConfigFromFile(confPath, v)
		if err != nil {
			fmt.Printf("load config file failed, err: %v", err)
		}
		err = yaml.Unmarshal([]byte(curConfig), &conf)
		if err != nil {
			fmt.Printf("unmarshal config file failed, err: %v", err)
		}
		fmt.Printf("%s\nconfig: %v\n", v, conf)
	}

}
