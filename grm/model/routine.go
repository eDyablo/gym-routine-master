package model

type Routine struct {
  Exercises []RoutineExcercise  `yaml:"exercises"`
  Name      string              `yaml:"name"`
}
