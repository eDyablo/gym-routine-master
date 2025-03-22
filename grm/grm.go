package main

import (
  "fmt"
  "gopkg.in/yaml.v3"
  "os"
  "strings"
)

type Link struct {
  Kind string `yaml:"kind,omitempty"`
  Link string `yaml:"link"`
}

type Exercise struct {
  AlternativeNames  []string  `yaml:"alternativeNames,omitempty"`
  ConventionalName  string    `yaml:"conventionalName"`
  Links             []Link    `yaml:"links,omitempty"`
}

type RoutineExcerciseSetType int

const (
  UnknownSet RoutineExcerciseSetType = iota
  FailSet
  FeederSet
  WarmUpSet
  WorkingSet
)

func (setType RoutineExcerciseSetType) String() string {
  var setName = map[RoutineExcerciseSetType]string{
    FailSet:    "fail",
    FeederSet:  "feeder",
    WarmUpSet:  "warm-up",
    WorkingSet: "working",
  }
  return setName[setType]
}

func ToRoutineExcerciseSetType(setName string) RoutineExcerciseSetType {
  var setType = map[string]RoutineExcerciseSetType{
    "fail":     FailSet,
    "feeder":   FeederSet,
    "warm-up":  WarmUpSet,
    "working":  WorkingSet,
  }
  return setType[strings.ToLower(setName)]
}

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

type RoutineExcerciseSet struct {
  Count Count                   `yaml:"count"`
  Loads []int                   `yaml:"loads,omitempty"`
  Type  RoutineExcerciseSetType `yaml:"type,omitempty"`
}

func (value RoutineExcerciseSetType) MarshalYAML() (interface{}, error) {
  return value.String(), nil
}

func (value *RoutineExcerciseSetType) UnmarshalYAML(node *yaml.Node) error {
  var raw string
  if err := node.Decode(&raw); err != nil {
    return err
  }
  *value = ToRoutineExcerciseSetType(raw)
  return nil
}

type PerceivedExertionRate struct {
  Max int `yaml:"max"`
  Min int `yaml:"min"`
}

type RestPeriod struct {
  Max int `yaml:"max"`
  Min int `yaml:"min"`
}

type SubstitutionOption string

type RoutineExcercise struct {
  Name          string                `yaml:"name"`
  Notes         string                `yaml:"notes,omitempty"`
  Reps          Count                 `yaml:"reps"`
  Rest          RestPeriod            `yaml:"rest"`
  RPE           PerceivedExertionRate `yaml:"rpe"`
  Sets          []RoutineExcerciseSet `yaml:"sets"`
  Substitutions []SubstitutionOption  `yaml:"substitutions,omitempty"`
}

type Routine struct {
  Exercises []RoutineExcercise  `yaml:"exercises"`
  Name      string              `yaml:"name"`
}

type Program struct {
  Exercises map[string]Exercise `yaml:"exercises"`
  Name      string              `yaml:"name"`
  Routine   []Routine           `yaml:"routine"`
}

type TrainingDrill struct {
  Exercise  string  `yaml:"exercise,omitempty"`
  Load      int     `yaml:"load"`
  RepCount  int     `yaml:"reps"`
}

type TrainingSet struct {
  Drills    []TrainingDrill `yaml:"drills"`
  Exercise  string          `yaml:"exercise"`
}

type TrainingSession struct {
  Sets []TrainingSet `yaml:"sets"`
}

func (program *Program) save() error {
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

func loadProgram(name string) *Program {
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

func main() {
  program := loadProgram("Ultimate PPL")
  fmt.Println(program)
  program.save()
}
