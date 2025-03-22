package model

type TrainingSet struct {
  Drills    []TrainingDrill `yaml:"drills"`
  Exercise  string          `yaml:"exercise"`
}
