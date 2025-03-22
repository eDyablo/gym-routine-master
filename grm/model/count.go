package model

import (
  "gopkg.in/yaml.v3"
)

type Count struct {
  Max int `yaml:"max"`
  Min int `yaml:"min"`
}

func (count Count) MarshalYAML() (interface{}, error) {
  if count.Min == count.Max {
    return count.Min, nil
  }
  return map[string]int{"min": count.Min, "max": count.Max,}, nil
}

func (count *Count) UnmarshalYAML(node *yaml.Node) error { 
  switch node.Tag {
  case "!!int":
    var value int
    if err := node.Decode(&value); err != nil {
      return err
    }
    count.Min = value
    count.Max = value
  case "!!map":
    var value map[string]int
    if err := node.Decode(&value); err != nil {
      return err
    }
    count.Min = value["min"]
    count.Max = value["max"]
  }
  return nil
}