package model

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
