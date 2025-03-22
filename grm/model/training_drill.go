package model

type TrainingDrill struct {
  Exercise  string  `yaml:"exercise,omitempty"`
  Load      int     `yaml:"load"`
  RepCount  int     `yaml:"reps"`
}
