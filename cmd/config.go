package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type myData struct {
	Conf struct {
		URI  string `yaml:"dburi"`
		User string `yaml:"user"`
		Pwd  string `yaml:"password"`
	}
}

func readConf(filename string) (*myData, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &myData{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}
