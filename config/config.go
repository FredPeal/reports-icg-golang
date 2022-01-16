package config

import (
    "io/ioutil"
    "gopkg.in/yaml.v3"
    "path/filepath"
)


type Conf struct {
    Connection string `yaml:"connection"`
}

func GetConf() Conf {

    filename, _ := filepath.Abs("config.yml")
    yamlFile, err := ioutil.ReadFile(filename)
    var config Conf
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }
    return config
}