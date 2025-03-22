package model

import (
  "gopkg.in/yaml.v3"
)

type IntInterval struct {
  Max int `yaml:"max"`
  Min int `yaml:"min"`
}

func (interval IntInterval) MarshalYAML() (interface{}, error) {
  if interval.Min == interval.Max {
    return interval.Min, nil
  }
  value := map[string]int{
    "max": interval.Max,
    "min": interval.Min,
  }
  return value, nil
}

func (interval *IntInterval) UnmarshalYAML(node *yaml.Node) error { 
  switch node.Tag {
  case "!!int":
    var value int
    if err := node.Decode(&value); err != nil {
      return err
    }
    interval.Min = value
    interval.Max = value
  case "!!map":
    var value map[string]int
    if err := node.Decode(&value); err != nil {
      return err
    }
    interval.Min = value["min"]
    interval.Max = value["max"]
  }
  return nil
}
