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
  if file, err := os.Create(filename); err == nil {
    defer file.Close()
    encoder := yaml.NewEncoder(file)
    encoder.SetIndent(2)
    return encoder.Encode(program)
  } else {
    return err
  }
}

func LoadProgram(filePath string) *Program {
  data, err := os.ReadFile(filePath)
  if err != nil {
    panic(err)
  }
  var program Program
  if err := yaml.Unmarshal([]byte(data), &program); err != nil {
    panic(err)
  }
  return &program
}
