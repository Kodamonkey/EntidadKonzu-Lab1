package utils

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
    Port       string `yaml:"port"`
    LogisticsAddress string `yaml:"logistics_address"`
}

func LoadConfig(file string) (*Config, error) {
    var config Config
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}
