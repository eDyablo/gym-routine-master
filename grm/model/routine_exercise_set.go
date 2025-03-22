package model

import (
  "gopkg.in/yaml.v3"
  "strings"
)

type RoutineExerciseSetType int

const (
  UnknownSet RoutineExerciseSetType = iota
  FailSet
  FeederSet
  WarmUpSet
  WorkingSet
)

func (setType RoutineExerciseSetType) String() string {
  var setName = map[RoutineExerciseSetType]string{
    FailSet:    "fail",
    FeederSet:  "feeder",
    WarmUpSet:  "warm-up",
    WorkingSet: "working",
  }
  return setName[setType]
}

func ToRoutineExerciseSetType(setName string) RoutineExerciseSetType {
  var setType = map[string]RoutineExerciseSetType{
    "fail":     FailSet,
    "feeder":   FeederSet,
    "warm-up":  WarmUpSet,
    "working":  WorkingSet,
  }
  return setType[strings.ToLower(setName)]
}

type RoutineExcerciseSet struct {
  Count Count                   `yaml:"count"`
  Loads []int                   `yaml:"loads,omitempty"`
  Type  RoutineExerciseSetType  `yaml:"type,omitempty"`
}

func (value RoutineExerciseSetType) MarshalYAML() (interface{}, error) {
  return value.String(), nil
}

func (value *RoutineExerciseSetType) UnmarshalYAML(node *yaml.Node) error {
  var raw string
  if err := node.Decode(&raw); err != nil {
    return err
  }
  *value = ToRoutineExerciseSetType(raw)
  return nil
}
