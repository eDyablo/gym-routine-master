package model

import (
  "gopkg.in/yaml.v3"
  "os"
)

type Program struct {
  Exercises map[string]Exercise `yaml:"exercises"`
  Name      string              `yaml:"name"`
  Routine   []Routine           `yaml:"routine"`
}

func (program *Program) Save() error {
  filename := program.Name + ".yaml"
  file, err := os.Create(filename)
  if err == nil {
    defer file.Close()
    encoder := yaml.NewEncoder(file)
    encoder.SetIndent(2)
    err = encoder.Encode(program)
  }
  return err
}

func LoadProgram(name string) *Program {
  filename := name + ".yaml"
  data, err := os.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  var program Program
  if err := yaml.Unmarshal([]byte(data), &program); err != nil {
    panic(err)
  }
  return &program
}
